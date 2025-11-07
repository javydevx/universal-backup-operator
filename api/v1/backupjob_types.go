/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupTarget defines what resource or path to back up.
type BackupTarget struct {
	Kind      string `json:"kind"`      // e.g. StatefulSet, Deployment, PVC
	Name      string `json:"name"`      // resource name
	Namespace string `json:"namespace"` // optional, defaults to current
	Path      string `json:"path"`      // path inside container or volume
}

// BackupDestination defines where to store the backup.
type BackupDestination struct {
	Type      string `json:"type"`      // s3, gcs, azure, git, nfs, local, custom
	URI       string `json:"uri"`       // e.g. s3://bucket/path or git@repo
	SecretRef string `json:"secretRef"` // name of Secret for credentials
}

// BackupStrategy defines how the backup is performed.
type BackupStrategy struct {
	Type    string `json:"type"`              // dump, snapshot, copy, custom
	Command string `json:"command,omitempty"` // optional shell command override
}

// BackupJobSpec defines the desired state of BackupJob.
type BackupJobSpec struct {
	Target      BackupTarget      `json:"target"`
	Destination BackupDestination `json:"destination"`
	Strategy    BackupStrategy    `json:"strategy"`
	Schedule    string            `json:"schedule,omitempty"` // optional Cron schedule
}

// BackupJobStatus defines the observed state of BackupJob.
type BackupJobStatus struct {
	Phase        string      `json:"phase,omitempty"` // Pending, Running, Succeeded, Failed
	LastRunTime  metav1.Time `json:"lastRunTime,omitempty"`
	Message      string      `json:"message,omitempty"`
	LastJobName  string      `json:"lastJobName,omitempty"`
	LastExitCode int32       `json:"lastExitCode,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BackupJob is the Schema for the backupjobs API.
type BackupJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupJobSpec   `json:"spec,omitempty"`
	Status BackupJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupJobList contains a list of BackupJob.
type BackupJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupJob{}, &BackupJobList{})
}
