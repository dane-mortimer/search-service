# Search Service

A simple search and text-suggestion service built on OpenSearch with a NextJS frontend, prometheus and grafana for observability. 

The plan is to extend this service, add OpenSearch cluster on AWS as data store and ingest data from DynamoDB into OpenSearch via DynamoDB streams.

The compute will be deployed on ECS with the react application deployed on S3. 

# Architecture

![Architecture](./assets/SearchServiceArchitecture.png)

# Project Structure

```
.
├── README.md               
├── assets                  # Architecture diagrams 
├── docker-compose.yaml     # Docker-compose
├── frontend                # NextJS frontend app 
├── grafana                 # Grafana dashboards configuration
├── ingestion-service       # DynamoDB to Opensearch ingestion service
├── integration-tests       # Integration tests
├── Dockerfile.localstack   # Dockerfile to setup local stack resources
├── prometheus              # Prometheus configuration
├── scripts                 # Helper scripts
└── search-service          # Backend Search service
```

# Usage 

## Requirements

* Go installed - v1.24
* Docker installed - >v20.10.10
* Docker compose installed - >v2.1.2
* Python3.11 installed
    * pip installed

``` bash
# Helper for debugging AWS local 
pip install aws-local

# e.g.
# awslocal dynamodb list-tables
```

## Launch Service

``` bash
PREFIX="course"

# Launch services
ENV=local \
COURSE_INDEX=${PREFIX}-index \
COURSE_TABLE=${PREFIX}-table \
OPENSEARCH_ENDPOINT=http://opensearch:9200 \
docker-compose up -d --build

# In seperate terminals you can get the logs 
# for the different containers

# AWS Local Stack Logs 
docker logs localstack -f 

# Search service logs 
docker logs search-service -f
```

## Load mock data

Now that the index is created an local stack is provisioned, lets load some data. 

This script creates a new item every 2 seconds to prevent overloading system resources and allowing the lambda to index documents in OpenSearch.

We can execute in a new terminal and leave it running in the background. 

``` bash
./scripts/create-mock-data.sh
```

## View metrics

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
