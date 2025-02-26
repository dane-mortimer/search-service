#!/usr/bin/env bash

for i in {1..2000}; do

  echo -e "\n🔎 Searching\n"

  curl "http://localhost:8080/search?q=OpenSearch&page=1&size=10"


  echo -e "\n💬 Suggestion\n" 
  curl "http://localhost:8080/suggest?q=OpenSearch"
  sleep 1
done