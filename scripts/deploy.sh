#!/bin/bash

set -e

echo "🔨 Building Docker image..."
docker build -t devops-platform:latest -f docker/app/Dockerfile .

echo "📦 Loading image into Kind..."
kind load docker-image devops-platform:latest --name devops-cluster

echo "🚀 Applying Kubernetes manifests..."
kubectl apply -k k8s/base/

echo "♻️ Restarting deployment..."
kubectl rollout restart deployment devops-app -n devops-platform

echo "⏳ Waiting for pods..."
kubectl rollout status deployment devops-app -n devops-platform

echo "✅ Done!"
echo "👉 Open: http://devops.local"
