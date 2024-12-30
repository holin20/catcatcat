#!/bin/bash

go run cmd/crawler/main.go 2>&1 | jq .