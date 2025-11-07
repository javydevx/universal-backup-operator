# 🧩 universal-backup-operator

A Kubernetes Operator written in Go that provides a **declarative way** to define and run backups to **any storage destination** — S3, GCS, Azure, Git repositories, NFS, or local PVCs.

---

## 🚀 Overview

The `universal-backup-operator` introduces a custom resource called `BackupJob` that allows cluster users to declare:

- **what** to back up (`target`)
- **how** to back it up (`strategy`)
- **where** to store it (`destination`)

The operator automatically schedules and executes Kubernetes `Jobs` or `CronJobs` to perform backups based on this definition.

---

## 📦 Example CRD Usage

```yaml
apiVersion: backup.example.com/v1
kind: BackupJob
metadata:
  name: postgres-daily-backup
spec:
  target:
    kind: StatefulSet
    name: postgres
    namespace: default
    path: /var/lib/postgresql/data
  destination:
    type: s3
    uri: s3://my-bucket/backups/postgres
    secretRef: s3-credentials
  strategy:
    type: dump
    command: "pg_dumpall > /backup/dump.sql"
  schedule: "0 2 * * *"
```

---

## 🌐 Supported Destinations (planned)

| Type | Description | Example |
|------|--------------|----------|
| `s3` | AWS S3, MinIO, Wasabi | `s3://mybucket/backups` |
| `gcs` | Google Cloud Storage | `gs://mybucket/backups` |
| `azure` | Azure Blob Storage | `az://container/path` |
| `git` | Git-based backup repo | `git@github.com:user/backups.git` |
| `nfs` | On-prem or local NFS mount | `nfs://server/path` |
| `local` | PersistentVolumeClaim | `pvc://my-backups` |
| `custom` | Run arbitrary backup command | - |

---

## 🧱 Architecture

- **Custom Resource** → `BackupJob` CRD defines backup specifications.  
- **Controller** → watches for CR changes, creates `Job`/`CronJob`.  
- **Backup Runner** → small container images implementing backup logic for each destination type.  
- **Status updates** → operator tracks `BackupJob.status` (phase, lastRun, message).

---

## 🧰 Prerequisites

Before development or deployment:

- Go **1.22.5+**
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Docker](https://www.docker.com/)
- [kind](https://kind.sigs.k8s.io/) or another Kubernetes cluster
- [kubebuilder](https://book.kubebuilder.io/quick-start.html)

---

## ⚙️ Development Setup

### 1️⃣ Clone or Create the Repo

```bash
mkdir -p ~/code/universal-backup-operator
cd ~/code/universal-backup-operator
git init
git branch -m main
```

---

### 2️⃣ Initialize Go Module

```bash
go mod init github.com/YOUR_GITHUB_USER/universal-backup-operator
```

---

### 3️⃣ Scaffold Project Using Kubebuilder

Kubebuilder is the **canonical** way to scaffold a Go Operator.

```bash
kubebuilder init --domain javy.dev --repo github.com/javydevx/universal-backup-operator
kubebuilder create api --group backup --version v1 --kind BackupJob --resource --controller
```

This creates the following structure:

```
api/v1/backupjob_types.go           # CRD schema
controllers/backupjob_controller.go # Reconcile logic
config/                             # CRDs, RBAC, manager manifests
main.go                             # entrypoint
```

---

### 4️⃣ Generate Manifests

```bash
make manifests
```

---

### 5️⃣ Run the Operator Locally

```bash
make run
```

---

### 6️⃣ Apply CRD and Create a Sample BackupJob

```bash
kubectl apply -f config/crd/bases/backup.example.com_backupjobs.yaml
kubectl apply -f config/samples/backup_v1_backupjob.yaml
```

---

## 🧩 Roadmap

| Milestone | Status | Description |
|------------|---------|-------------|
| Scaffold Operator with Kubebuilder | ✅ | Basic setup |
| Implement S3 backups | 🚧 | MVP target |
| Add Git backup support | ⏳ | Push to Git repos |
| Add GCS & Azure destinations | ⏳ | Cloud expansion |
| Support PVC/NFS local backups | ⏳ | On-prem support |
| Add Cron scheduling | ⏳ | Recurring backups |
| Add status tracking | ⏳ | Phase, timestamps |
| Expose `/restore` endpoint via Aggregated API Server | ⏳ | API extension feature |

---

## 🧑‍💻 Contributing

Contributions are welcome!  
This project is **open-source and educational**, intended to help developers learn **Kubernetes Operator patterns** and **API extension techniques**.

---

## 📘 References

| Resource | Link | Notes |
|-----------|------|-------|
| 🧱 Kubebuilder Book | [https://book.kubebuilder.io/](https://book.kubebuilder.io/) | Official guide for building operators |
| ⚙️ Operator SDK | [https://sdk.operatorframework.io/docs/](https://sdk.operatorframework.io/docs/) | Alternative framework |
| 📘 controller-runtime | [https://github.com/kubernetes-sigs/controller-runtime](https://github.com/kubernetes-sigs/controller-runtime) | Core library |
| 🧩 CronJob tutorial | [https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial.html](https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial.html) | Good starting example |
| 🧰 Kubernetes Code Generator | [https://github.com/kubernetes/code-generator](https://github.com/kubernetes/code-generator) | Underlying codegen tool |

---

## 📄 License

This project is licensed under the **MIT License**.

---

### ✨ Author’s Note

The goal of this project is to demonstrate:

- Building a Kubernetes Operator in Go  
- Designing flexible CRDs  
- Extending the Kubernetes API in two ways:
  1. **CustomResourceDefinitions (CRDs)**
  2. **Aggregated API Servers**

By the end, you’ll have a working backup operator that can evolve into a production-grade system.

---

💡 **Next Step:**  
Run `kubebuilder init` and scaffold the project, then we’ll define the CRD (`BackupJob`) fields and reconcile logic next!
