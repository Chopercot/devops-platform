#!/bin/bash

IMAGE=devops-platform

echo "Building docker image..."

docker build -t $IMAGE -f docker/app/Dockerfile .
