#!/bin/bash
set -e

PROJECT_ID="ai-credit-product"
SERVICE_NAME="naqa-api"
REGION="us-central1"
IMAGE_NAME="gcr.io/$PROJECT_ID/$SERVICE_NAME"

echo "🚀 Building and deploying $SERVICE_NAME..."

docker build -t $IMAGE_NAME:latest .
docker push $IMAGE_NAME:latest

cd terraform
terraform init
terraform apply -auto-approve

echo "✅ Deployment completed!"