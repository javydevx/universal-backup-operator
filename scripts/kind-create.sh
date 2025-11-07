#!/usr/bin/env bash
set -e

CLUSTER_NAME="universal-backup"
CONFIG_FILE="deployments/local/kind-config.yaml"

echo "🚀 Creating kind cluster: $CLUSTER_NAME..."
kind create cluster --name "$CLUSTER_NAME" --config "$CONFIG_FILE"

echo "✅ Kind cluster '$CLUSTER_NAME' is ready!"
kubectl cluster-info --context kind-$CLUSTER_NAME