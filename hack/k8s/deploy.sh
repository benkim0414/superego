#!/bin/bash
kubectl set image deployment/${GCP_PROJECT_NAME} ${IMAGE_NAME}=${GCR_HOSTNAME}/${GCP_PROJECT_ID}/${IMAGE_NAME}:latest
