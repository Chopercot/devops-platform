#!/bin/bash

DATE=$(date +%F)

echo "Creating backup..."

kubectl get all -n devops-platform -o yaml > backup-$DATE.yaml
