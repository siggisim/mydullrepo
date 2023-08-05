external/gcloud/gcloud services enable run.googleapis.com --project=$GCP_PROJECT

external/gcloud/gcloud run deploy test-$RANDOM --image=gcr.io/flame-public/buildbuddy-app-onprem:latest --port=8080 --region=us-west1 --project=$GCP_PROJECT --allow-unauthenticated
