#!/usr/bin/env python

import json
import requests
import time

headers = {
    "Content-Type": "application/json",
    "Accept": "application/json",
    "X-CSRF-Token": "MTc0MDkwNjYyM3xJbWgyWm5OVVUza3JOblZhWm5SS2NWQjFPSGsyYTBkcFpVWnBXWGxGVjFJeWRUZFJWVGxKUlVSM1ptYzlJZ289fJE1ELnVNfOqN4a3WhbeFVp0c1RKsJ6p5A9E5oo6HI1X",  # Add the CSRF token to the headers
}

# URL to which the data will be posted
url = "http://localhost:8080/api/v1/course/"  # Replace with your actual URL

# Load JSON data from a file
with open("data.json", "r") as file:
    json_data = json.load(file)

success_docs_count = 0 
failed_docs_count = 0 

# Iterate through each item in the "data" array
for index, item in enumerate(json_data["data"]):
    try:
        # Send a POST request with the item as the JSON body
        response = requests.post(url, json=item, headers=headers)
        # Check if the request was successful
        if response.status_code == 201:
            success_docs_count += 1
        else:
            failed_docs_count += 1

        print(f"\rSuccess docs count: {success_docs_count}, Failed docs count: {failed_docs_count}", end="")
        
        # Sleep to take load off the lambda function
        time.sleep(2)
    except requests.exceptions.RequestException as e:
        print(f"Error posting {item['title']}: {e}")