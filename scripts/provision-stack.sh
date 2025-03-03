#!/usr/bin/env bash

set -e

if [ $# != 3 ]; then
    echo -e "$0 expects 3 arguments: index_name, table_name, opensearch_endpoint"
    echo -e "Example: $0 course-index course-table http://localhost:9200"
    exit 1
fi 

INDEX_NAME=$1
TABLE_NAME=$2
OPENSEARCH_ENDPOINT=$3

export FUNCTION_NAME="${TABLE_NAME}-to-${INDEX_NAME}"

export LAMBDA_BASE_DIR="ingestion-service"


# Build lambda package
(
    cd ${LAMBDA_BASE_DIR} \
    && python3 -m venv env \
    && source env/bin/activate \
    && mkdir -p package \
    && pip install -r requirements.txt --target package  \
    && deactivate
)
(cd ${LAMBDA_BASE_DIR}/package && zip -r ../lambda_function.zip . >/dev/null)
(cd ${LAMBDA_BASE_DIR} && zip lambda_function.zip handler.py >/dev/null)

echo -e "\nCreating DynamoDB Table" 
TABLE_STREAM=$(
    awslocal dynamodb create-table \
        --table-name ${TABLE_NAME} \
        --attribute-definitions AttributeName=id,AttributeType=S \
        --key-schema AttributeName=id,KeyType=HASH \
        --billing-mode PAY_PER_REQUEST \
        --stream-specification StreamEnabled=true,StreamViewType=NEW_AND_OLD_IMAGES \
        --output text \
        --query TableDescription.LatestStreamArn 
)

echo -e "\nCreating Lambda Function" 
awslocal lambda create-function \
    --function-name ${FUNCTION_NAME} \
    --runtime python3.11 \
    --handler handler.handler \
    --environment "Variables={OPENSEARCH_INDEX=${INDEX_NAME},OPENSEARCH_ENDPOINT=${OPENSEARCH_ENDPOINT}}" \
    --role arn:aws:iam::000000000000:role/MyLambdaRole \
    --zip-file fileb://${LAMBDA_BASE_DIR}/lambda_function.zip \
    --output text >/dev/null

echo -e "\nCreating Lambda / DynamoDB Stream Integration" 

awslocal lambda create-event-source-mapping \
    --function-name ${FUNCTION_NAME} \
    --event-source-arn $TABLE_STREAM \
    --starting-position LATEST \
    --output text >/dev/null

echo -e "\nWaiting for lambda function to become active" 
awslocal lambda wait function-active-v2 --function-name ${FUNCTION_NAME}

echo -e "\nAll Resources provisioned" 

# Examples
# echo -e "\nAdding Test Item to DynamoDB" 
# awslocal dynamodb put-item \
#     --table-name ${TABLE_NAME} \
#     --item '{"id": {"S": "1"}, "title": {"S": "Test Item1"}, "content": { "S": "Test Content1" }, "owner": {"S": "Test Owner" } }'
#
# echo -e "\nChecking OpenSearch Index is updated" 
# curl -X GET "http://localhost:9200/${INDEX_NAME}/_search?pretty" -H 'Content-Type: application/json'
#
# PAYLOAD=$(echo -e '{"Records": [{"eventID": "1", "eventName": "INSERT", "dynamodb": {"NewImage": {"id": {"S": "1"}, "title": {"S": "Test Item1"}, "content": { "S": "Test Content" }, "owner": {"S": "Test Owner" } }}}]}' | base64)
#
# awslocal lambda invoke \
#     --function-name ${FUNCTION_NAME} \
#     --payload  ${PAYLOAD} \
#     output.txt