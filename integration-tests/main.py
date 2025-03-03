import pytest
import requests
import os
import time

# Base URL for the API
SEARCH_SERVICE_DOMAIN = os.getenv("SEARCH_SERVICE_DOMAIN")

COURSE_URI = "api/v1/course"

BASE_URL = f"{SEARCH_SERVICE_DOMAIN}/{COURSE_URI}"

def test_create_course_should_pass():
    # Test creating a new course, happy path
    payload = {
        "title": "Introduction to Go",
        "content": "Learn the basics of Go programming.",
        "owner": "user-123"
    }
    response = requests.post(f"{BASE_URL}", json=payload)
    assert response.status_code == 201
    response_dict = response.json()
    data = response_dict['data']
    assert "id" in data
    assert data["title"] == payload["title"]
    assert data["content"] == payload["content"]
    assert data["owner"] == payload["owner"]

# Need to fix this test case
# def test_create_course_should_fail():
#     # Test create course should fail, invalid key
#     payload = {
#         "title": "Introduction to Go",
#         "content": "Learn the basics of Go programming.",
#         "owner": 1234
#     }
#     response = requests.post(f"{BASE_URL}", json=payload)
#     assert response.status_code == 400
#     response_dict = response.json()
#     assert data not in response_dict
#     assert response["status"] == "error"

def test_get_course_should_pass():
    # First, create a course to retrieve
    payload = {
        "title": "Advanced Go",
        "content": "Learn advanced Go programming.",
        "owner": "user-456"
    }
    create_response = requests.post(f"{BASE_URL}", json=payload)
    course_id = create_response.json()["data"]["id"]

    # Test retrieving the course
    response = requests.get(f"{BASE_URL}/{course_id}")
    assert response.status_code == 200
    response_dict = response.json()
    data = response_dict['data']
    assert data["id"] == course_id
    assert data["title"] == payload["title"]
    assert data["content"] == payload["content"]
    assert data["owner"] == payload["owner"]


def test_get_course_should_fail():
    # First, create a course to retrieve
    payload = {
        "title": "Advanced Go",
        "content": "Learn advanced Go programming.",
        "owner": "user-456"
    }
    create_response = requests.post(f"{BASE_URL}", json=payload)
    course_id = create_response.json()["data"]["id"]

    # Test retrieving the course
    response = requests.get(f"{BASE_URL}/{course_id}")
    assert response.status_code == 200
    response_dict = response.json()
    data = response_dict['data']
    assert data["id"] == course_id
    assert data["title"] == payload["title"]
    assert data["content"] == payload["content"]
    assert data["owner"] == payload["owner"]
    
def test_search_course_should_pass():
    # Sleep to wait for Opensearch docs to be indexed
    time.sleep(2)
    # Test searching for courses
    response = requests.get(f"{BASE_URL}/search", params={"q": "Go"})
    assert response.status_code == 200
    response_dict = response.json()
    data = response_dict['data']
    assert isinstance(data, list)

def test_search_course_should_be_none():
    # Sleep to wait for Opensearch docs to be indexed
    time.sleep(2)
    # Test searching for courses
    response = requests.get(f"{BASE_URL}/search", params={"q": "2fds8"})
    assert response.status_code == 200
    response_dict = response.json()
    data = response_dict['data']
    assert data is None
    
def test_suggest_course_should_pass():
    # Sleep to wait for Opensearch docs to be indexed
    time.sleep(2)
    # Test suggesting courses
    response = requests.get(f"{BASE_URL}/suggest", params={"q": "Go"})
    assert response.status_code == 200
    response_dict = response.json()
    data = response_dict['data']
    assert isinstance(data, list)

def test_suggest_course_should_be_none():
    # Test suggesting courses
    response = requests.get(f"{BASE_URL}/suggest", params={"q": "2fds8"})
    assert response.status_code == 200
    response_dict = response.json()
    data = response_dict['data']
    assert data is None

def test_options_request():
    # Test OPTIONS request for CORS
    response = requests.options(f"{BASE_URL}")
    assert response.status_code == 200
    assert response.headers["Access-Control-Allow-Origin"] == "*"
    assert response.headers["Access-Control-Allow-Methods"] == "GET, POST, OPTIONS, PUT, DELETE"
    assert response.headers["Access-Control-Allow-Headers"] == "Content-Type, Authorization"