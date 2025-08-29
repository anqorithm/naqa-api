#!/bin/bash
set -e

PROJECT_ID="ai-credit-product"

echo "🚀 Building with Cloud Build..."
gcloud builds submit --config cloudbuild.yaml --project $PROJECT_ID

echo "🚀 Deploying with Terraform..."
cd terraform
terraform init
terraform apply -auto-approve

echo "✅ Deployment completed!"