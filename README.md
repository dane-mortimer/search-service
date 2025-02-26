# Search Service

A simple search service built on OpenSearch with prometheus and grafana for observability. 

The plan is to extend this service, add OpenSearch cluster on AWS as data store and ingest data from DynamoDB into OpenSearch via DynamoDB streams.

# Usage 

## Requirements

* Go installed
* Docker installed
* Docker compose installed

## Launch Service

``` bash
docker-compose up -d --build
``` 

## Load mock data into service

``` bash
./scripts/upload-opensearch-local.sh
``` 

## Build Grafana Dashboard

1. Navigate to `http:localhost:3000`
2. Login using username: `admin`, password: `admin` 
3. Add Prometheus as a data source: `http:localhost:9090`
4. Import dashboard from JSON
5. Copy contents of grafana.json

## Hit the service

This script will run and hit the service 200 times, you can see the metrics being imported into grafana.

``` bash
./scripts/test-service.sh
```

## Clean up

``` bash 
docker-compose down
```
