#!/usr/bin/env bash

# if [ $# != '4' ]; then
#     echo -e "$0 expects 4 arguments: COURSE_INDEX, COURSE_TABLE, opensearch_endpoint, localstack_endpoint"
#     echo -e "Example: $0 course-index course-table http://localhost:9200"
#     echo -e "got: $#"
#     exit 1
# fi 

# COURSE_INDEX=$1
# COURSE_TABLE=$2
# OPENSEARCH_ENDPOINT=$3
# LOCAL_STACK_ENDPOINT=$4


echo -e "Course Index: $COURSE_INDEX"
echo -e "Course Table: $COURSE_TABLE"
echo -e "OpenSearch Endpoint: $OPENSEARCH_ENDPOINT"
echo -e "Local Stack Endpoint: $LOCAL_STACK_ENDPOINT"

export FUNCTION_NAME="${COURSE_TABLE}-to-${COURSE_INDEX}"

export LAMBDA_BASE_DIR="ingestion-service"

# Build lambda package
(
    cd ${LAMBDA_BASE_DIR} \
    && python -m venv env \
    && source env/bin/activate \
    && mkdir -p package \
    && pip install -r requirements.txt --target package  \
    && deactivate
)
(cd ${LAMBDA_BASE_DIR}/package && zip -r ../lambda_function.zip . >/dev/null)
(cd ${LAMBDA_BASE_DIR} && zip lambda_function.zip handler.py >/dev/null)

echo -e "\nCreating Course Index" 

curl -X PUT "${OPENSEARCH_ENDPOINT}/$COURSE_INDEX" -H "Content-Type: application/json" -d '
{
  "mappings": {
    "properties": {
      "title": {
        "type": "search_as_you_type"
      },
      "content": {
        "type": "search_as_you_type"
      }
    }
  }
}'


# Pro local stack feature - cannot implement yet
# USER_POOL_ID=$(
#     awslocal cognito-idp create-user-pool \
#     --endpoint ${LOCAL_STACK_ENDPOINT} \
#     --pool-name "user-pool" \
#     --schema Name=role,AttributeDataType=string,DeveloperOnlyAttribute=false,Mutable=true,Required=true \
#     --pool-name MyUserPool \
#     --query 'UserPool.Id' \
#     --output text
# )

# if [ -z "$USER_POOL_ID" ]; then
#     echo "Failed to create Cognito User Pool"
# else
#     echo "Storing UserPoolId in SSM Parameter Store..."
#     awslocal ssm put-parameter \
#         --endpoint ${LOCAL_STACK_ENDPOINT}
#         --name "/cognito/user-pool-id" \
#         --value "$USER_POOL_ID" \
#         --type "String" \
#         --overwrite
# fi

echo -e "\nCreating DynamoDB Table" 
TABLE_STREAM=$(
    awslocal dynamodb create-table \
        --endpoint ${LOCAL_STACK_ENDPOINT} \
        --table-name ${COURSE_TABLE} \
        --attribute-definitions AttributeName=id,AttributeType=S \
        --key-schema AttributeName=id,KeyType=HASH \
        --billing-mode PAY_PER_REQUEST \
        --stream-specification StreamEnabled=true,StreamViewType=NEW_AND_OLD_IMAGES \
        --output text \
        --query TableDescription.LatestStreamArn 
)

echo -e "\nCreating Lambda Function" 
awslocal lambda create-function \
    --endpoint ${LOCAL_STACK_ENDPOINT} \
    --function-name ${FUNCTION_NAME} \
    --runtime python3.11 \
    --handler handler.handler \
    --environment "Variables={OPENSEARCH_INDEX=${COURSE_INDEX},OPENSEARCH_ENDPOINT=${OPENSEARCH_ENDPOINT}}" \
    --role arn:aws:iam::000000000000:role/MyLambdaRole \
    --zip-file fileb://${LAMBDA_BASE_DIR}/lambda_function.zip \
    --output text >/dev/null

echo -e "\nCreating Lambda / DynamoDB Stream Integration" 

awslocal lambda create-event-source-mapping \
    --endpoint ${LOCAL_STACK_ENDPOINT} \
    --function-name ${FUNCTION_NAME} \
    --event-source-arn $TABLE_STREAM \
    --starting-position LATEST \
    --output text >/dev/null

echo -e "\nWaiting for lambda function to become active" 
awslocal lambda wait function-active-v2 --endpoint ${LOCAL_STACK_ENDPOINT} --function-name ${FUNCTION_NAME}

echo -e "\nAll Resources provisioned" 

# Mark container as ready
touch /tmp/ready

echo "Entrypoint script finished."
sleep 60


# Examples
# echo -e "\nAdding Test Item to DynamoDB" 
# awslocal dynamodb put-item \
#     --table-name ${COURSE_TABLE} \
#     --item '{"id": {"S": "1"}, "title": {"S": "Test Item1"}, "content": { "S": "Test Content1" }, "owner": {"S": "Test Owner" } }'
#
# echo -e "\nChecking OpenSearch Index is updated" 
# curl -X GET "http://localhost:9200/${COURSE_INDEX}/_search?pretty" -H 'Content-Type: application/json'
#
# PAYLOAD=$(echo -e '{"Records": [{"eventID": "1", "eventName": "INSERT", "dynamodb": {"NewImage": {"id": {"S": "1"}, "title": {"S": "Test Item1"}, "content": { "S": "Test Content" }, "owner": {"S": "Test Owner" } }}}]}' | base64)
#
# awslocal lambda invoke \
#     --function-name ${FUNCTION_NAME} \
#     --payload  ${PAYLOAD} \
#     output.txt