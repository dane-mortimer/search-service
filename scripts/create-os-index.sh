#!/usr/bin/env bash

if [ $# != 1 ]; then
  echo -e "$0 expects 1 argument: index_name"
  exit 1
fi 

INDEX_NAME=$1

curl -X PUT "http://localhost:9200/$INDEX_NAME" -H "Content-Type: application/json" -d '
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