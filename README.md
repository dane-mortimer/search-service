# Search Service

A simple search and text-suggestion service built on OpenSearch with a NextJS frontend, prometheus and grafana for observability. 

The plan is to extend this service, add OpenSearch cluster on AWS as data store and ingest data from DynamoDB into OpenSearch via DynamoDB streams.

The compute will be deployed on ECS with the react application deployed on S3. 

# Architecture

![Architecture](./assets/SearchServiceArchitecture.png)

# Project Structure


```
├── clients               # External API clients (e.g. Opensearch)
├── docker-compose.yml    # Docker Compose, deploy the app locally
├── go.mod                # Depedency Management
├── grafana.json          # Grafana configuration
├── handlers              # API handlers
├── main.go               # entry point for search service
├── middleware            # Middleware - security, logging, cors etc
├── models                # Application comdels
├── prometheus.yml        # Prometheus configuration
├── react-search-app      # Basic react app to use application
├── scripts               # Upload services
├── services              # API Services
└── utils                 # Utility Files
```

# Usage 

## Requirements

* Go installed - v1.24
* Docker installed - v20.10.10
* Docker compose installed - v1.29.2

``` bash
pip install aws-local
```

## Build Lambda

``` bash
cd lambda
python3 -m venv env
source env/bin/activate
mkdir package
pip install -r requirements.txt --target package 
```

## Launch Service

``` bash
docker-compose up -d --build
``` 

## Load mock data into service

Once docker-compose has built and deployed, run this. 

This script creates the search index and uploads about 10,000 lines of mock data. 

``` bash
./scripts/upload-opensearch-local.sh
``` 

## Build Grafana Dashboard

Metrics included 

```
P99 Latency
Total Requests
Requests per interval
Error rate
```

1. Navigate to `http:localhost:3000`
2. Login using username: `admin`, password: `admin` 
3. Add Prometheus as a data source: `http:localhost:9090`
4. Import dashboard from JSON
5. Copy contents of grafana.json

## Hit the service

Navigate to `localhost:3001` and use the service

Or 

This script will run and hit the service 2000 times, you can see the metrics being imported into grafana.

``` bash
./scripts/test-service.sh
```

## Clean up

``` bash 
docker-compose down
```
