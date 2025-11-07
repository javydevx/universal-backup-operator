#!/usr/bin/env bash
set -e

CLUSTER_NAME="universal-backup"

echo "🧹 Deleting kind cluster: $CLUSTER_NAME..."
kind delete cluster --name "$CLUSTER_NAME"
echo "✅ Cluster deleted."
