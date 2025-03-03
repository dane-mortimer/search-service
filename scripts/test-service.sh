#!/usr/bin/env bash

COURSE_API_ENDPOINT=http://localhost:8080/api/v1/course

for i in {1..2000}; do

  echo -e "\n🔎 Searching\n"

  curl "${COURSE_API_ENDPOINT}/search?q=advanced&page=1&size=10"


  echo -e "\n💬 Suggestion\n" 
  curl "${COURSE_API_ENDPOINT}/suggest?q=advanced"
  sleep 1
done


curl "${COURSE_API_ENDPOINT}/01JNCMFV2RQPVM6MJRS94XS3JD"