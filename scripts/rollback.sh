#!/bin/bash

echo "↩️ Rolling back to BLUE..."

kubectl patch service devops-service \
  -n devops-platform \
  -p '{"spec":{"selector":{"app":"devops-app","version":"blue"}}}'

echo "✅ Traffic switched to BLUE"
