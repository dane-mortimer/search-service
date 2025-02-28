#!/usr/bin/env bash

export MY_INDEX='my-index'

sleep 2

curl -X PUT "http://localhost:9200/$MY_INDEX" -H "Content-Type: application/json" -d '
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

sleep 2 


# Read the contents of data.txt into an array, splitting by newlines
readarray -t docs < ./scripts/data.txt

echo -e "\n"
echo -e "\tUploading docs..."
echo -e ""

# Loop through each item in the array
for i in "${!docs[@]}"; do
  # Extract the title and content using '|' as the delimiter
  title="${docs[$i]%%|*}"
  content="${docs[$i]#*|}"

  # Print the progress (number of docs added so far)
  echo -ne "\r\tDocs Added: $((i+1))"

  # Send a POST request to the Elasticsearch index
  curl -s -o /dev/null -X POST "http://localhost:9200/$MY_INDEX/_doc/$((i+1))" \
    -H "Content-Type: application/json" \
    -d "{
      \"title\": \"$title\",
      \"content\": \"$content\"
    }"
done

echo -e "\n\n"

# Optional: Fetch and display the contents of the Elasticsearch index
# curl -X GET "http://localhost:9200/$MY_INDEX/_search?pretty"