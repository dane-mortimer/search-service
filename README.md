# Search Service

A simple search and text-suggestion service built on OpenSearch with a NextJS frontend, prometheus and grafana for observability. 

The plan is to extend this service, add OpenSearch cluster on AWS as data store and ingest data from DynamoDB into OpenSearch via DynamoDB streams.

The compute will be deployed on ECS with the react application deployed on S3. 

# Architecture

![Architecture](./assets/SearchServiceArchitecture.png)

# Project Structure

```
.
├── assets                                    # Architecture Diagrams
│   ├── SearchServiceArchitecture.drawio
│   └── SearchServiceArchitecture.png         
├── docker-compose.yaml                       # Local development docker compose
├── frontend
│   └── search-app                            # NextJS frontned app
├── grafana                                   # Grafana Configuration
├── ingestion-service                   
│   ├── handler.py                            # Ingestion service lambda
│   └── requirements.txt                      # Ingestion service dependencies
├── prometheus                                # Prometheus Configuration
├── scripts                                   # Local development setup scripts             
└── search-service
    ├── Dockerfile                            
    ├── clients                               # External Clients, Opensearch DynamoDB
    ├── controllers                           # Application controllers
    ├── dao                                   # Database access layer
    ├── go.mod                                # Dependency management
    ├── go.sum
    ├── handlers                              # Route handlers
    ├── main.go                               # Application entry point
    ├── middleware                            # Middleware - cors, security, loggers etc 
    ├── models                                # Models
    └── utils                                 # Utility functions
```

# Usage 

## Requirements

* Go installed - v1.24
* Docker installed - >v20.10.10
* Docker compose installed - >v2.1.2
* Python3.11 installed
    * pip installed

``` bash
pip install aws-local
```

## Setup Environment Variables

``` bash
ENV=local
PREFIX="course"
COURSE_INDEX=${PREFIX}-index
COURSE_TABLE=${PREFIX}-table
OPENSEARCH_ENDPOINT=http://opensearch:9200 # Opensearch docker network endpoint
```

## Launch Service

``` bash
# Launch services
ENV=$ENV COURSE_INDEX=$COURSE_INDEX COURSE_TABLE=$COURSE_TABLE OPENSEARCH_ENDPOINT=$OPENSEARCH_ENDPOINT docker-compose up -d --build

# In seperate terminals you can get the logs 
# for the different containers

# AWS Local Stack Logs 
docker logs localstack -f 

# Search service logs 
docker logs search-service -f
```

## Create Opensearch Index and AWS resources

Once docker-compose has built and deployed, run the following scripts. 

``` bash
# Create OpenSearch Index
./scripts/create-course-index.sh ${COURSE_INDEX}

# Build AWS Resources
./scripts/provision-stack.sh ${COURSE_INDEX} ${COURSE_TABLE} ${OPENSEARCH_ENDPOINT}
``` 

## Load mock data

Now that the index is created an local stack is provisioned, lets load some data. 

This script creates a new item every 2 seconds to prevent overloading system resources and allowing the lambda to index documents in OpenSearch.

We can execute in a new terminal and leave it running in the background. 

``` bash
./scripts/create-mock-data.sh
```

## Build Grafana Dashboard

Metrics included 

> P99 Latency \
> Total Requests \
> Requests per interval \
> Error rate

1. Navigate to `http:localhost:3000`
2. Login using username: `admin`, password: `admin`, then change the admin password.
3. The dashboard is automatically loaded into Grafana and Prometheus is connected as a datasource

## Hit the service

Navigate to `localhost:3001` and use the web app. 

Alternatively, this script will run and hit the service 2000 times, you can see the metrics being imported into grafana.

``` bash
./scripts/test-service.sh
```

## Clean up

``` bash 
docker-compose down
```
