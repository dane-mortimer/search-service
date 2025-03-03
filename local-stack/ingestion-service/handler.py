import os
import json
from opensearchpy import OpenSearch, RequestsHttpConnection
from requests_aws4auth import AWS4Auth
import boto3

# Environment variables
OPENSEARCH_ENDPOINT = os.getenv("OPENSEARCH_ENDPOINT")
OPENSEARCH_INDEX = os.getenv("OPENSEARCH_INDEX")

OPENSEARCH_DOMAIN = OPENSEARCH_ENDPOINT.replace('http://', '').replace('https://', '').split(":")
OPENSEARCH_HOST   = OPENSEARCH_DOMAIN[0]
OPENSEARCH_PORT   = OPENSEARCH_DOMAIN[1]

# Initialize OpenSearch client
def get_opensearch_client():
    if os.getenv("AWS_REGION"):  # If running in AWS, use AWS SigV4 authentication
        credentials = boto3.Session().get_credentials()
        aws_auth = AWS4Auth(
            credentials.access_key,
            credentials.secret_key,
            os.getenv("AWS_REGION"),
            'es',
            session_token=credentials.token
        )
        return OpenSearch(
            hosts=[{'host': OPENSEARCH_HOST, 'port': OPENSEARCH_PORT}],
            http_auth=aws_auth,
            use_ssl=True if OPENSEARCH_ENDPOINT.startswith('https://') else False,
            verify_certs=True,
            connection_class=RequestsHttpConnection
        )
    else:  # For local development or non-AWS environments
        return OpenSearch(
            hosts=[{'host': OPENSEARCH_HOST, 'port': OPENSEARCH_PORT}],
            http_auth=('username', 'password'),  
            use_ssl=True if OPENSEARCH_ENDPOINT.startswith('https://') else False,
            verify_certs=True,
            connection_class=RequestsHttpConnection
        )

def handler(event, context):
    # Initialize OpenSearch client
    client = get_opensearch_client()

    # Parse the event to get the item details and action
    for record in event['Records']:
        action = record['eventName']  # e.g., INSERT, MODIFY, REMOVE
        item = record['dynamodb']['NewImage'] if action in ['INSERT', 'MODIFY'] else record['dynamodb']['OldImage']
        
        item_id = item['id']['S']
        title = item['title']['S']
        content = item['content']['S']

        # Prepare the OpenSearch document
        document = {
            "title": title,
            "content": content
        }

        try:
            if action == 'INSERT' or action == 'MODIFY':
                # Index or update the document in OpenSearch
                response = client.index(
                    index=OPENSEARCH_INDEX,
                    id=item_id,
                    body=document,
                    refresh=True  # Optional: Refresh the index to make the changes visible immediately
                )
                print(f"Document indexed: {response}")
            elif action == 'REMOVE':
                # Delete the document from OpenSearch
                response = client.delete(
                    index=OPENSEARCH_INDEX,
                    id=item_id
                )
                print(f"Document deleted: {response}")
        except Exception as e:
            print(f"Error processing record: {e}")
            return {
                'statusCode': 500,
                'body': json.dumps(f"Error processing record: {e}")
            }

    return {
        'statusCode': 200,
        'body': json.dumps('OpenSearch index updated successfully')
    }