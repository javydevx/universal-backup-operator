#!/usr/bin/env bash
set -euo pipefail

# Use same cluster name variable as in Makefile / setup
KIND_CLUSTER_NAME=${KIND_CLUSTER_NAME:-universal-backup}

echo "🧹 Deleting kind cluster: $KIND_CLUSTER_NAME..."
kind delete cluster --name "$KIND_CLUSTER_NAME"
echo "✅ Cluster deleted!"
