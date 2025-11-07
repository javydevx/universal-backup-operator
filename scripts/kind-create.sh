#!/usr/bin/env bash
set -euo pipefail

KIND_CLUSTER_NAME=${KIND_CLUSTER_NAME:-universal-backup}
KIND_CONFIG_PATH=${KIND_CONFIG_PATH:-./config/local/kind-config.yaml}

echo "🚀 Creating kind cluster '${KIND_CLUSTER_NAME}'..."
kind create cluster --name "${KIND_CLUSTER_NAME}" --config "${KIND_CONFIG_PATH}"

kubectl cluster-info --context kind-${KIND_CLUSTER_NAME}
echo "✅ Kind cluster ready!"
