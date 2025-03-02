#!/usr/bin/env bash


MOCK_DATA_DIR=./scripts/mock-data

(
    cd $MOCK_DATA_DIR \
    && python3 -m venv env \
    && source env/bin/activate \
    && pip install -r requirements.txt \
    && python ingest-mock-data.py
)