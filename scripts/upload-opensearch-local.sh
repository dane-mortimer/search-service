#!/usr/bin/env bash

if [ $# != 1 ]; then
  echo -e "$0 expects 1 argument: index_name"
  exit 1
fi 

INDEX_NAME=$1

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
  curl -s -o /dev/null -X POST "http://localhost:9200/$INDEX_NAME/_doc/$((i+1))" \
    -H "Content-Type: application/json" \
    -d "{
      \"title\": \"$title\",
      \"content\": \"$content\"
    }"
done

echo -e "\n\n"

# Optional: Fetch and display the contents of the Elasticsearch index
# curl -X GET "http://localhost:9200/$MY_INDEX/_search?pretty"