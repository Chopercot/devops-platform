#!/bin/bash

echo "🔁 Switching to GREEN..."

kubectl patch service devops-service \
  -n devops-platform \
  -p '{"spec":{"selector":{"app":"devops-app","version":"green"}}}'

echo "✅ Traffic switched to GREEN"
