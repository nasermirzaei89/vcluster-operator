/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VClusterSpec defines the desired state of VCluster
type VClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Chart VClusterSpecChart `json:"chart,omitempty"`

	// If true the virtual cluster will not sync any ingresses
	DisableIngressSync bool `json:"disableIngressSync,omitempty"`

	// Kubernetes distro to use for the virtual cluster. Allowed distros: k3s, k0s, k8s, eks (default "k3s")
	Distro string `json:"distro,omitempty"`

	// If true will create a load balancer service to expose the vcluster endpoint
	Expose bool `json:"expose,omitempty"`

	// If true vcluster and its workloads will run in an isolated environment
	Isolate bool `json:"isolate,omitempty"`

	// The kubernetes version to use (e.g. v1.20). Patch versions are not supported
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
}

type VClusterSpecChart struct {
	// The virtual cluster chart name to use (default "vcluster")
	// +optional
	Name string `json:"name,omitempty"`

	// The virtual cluster chart repo to use (default "https://charts.loft.sh")
	// +optional
	Repo string `json:"repo,omitempty"`

	// The virtual cluster chart version to use (e.g. v0.4.0) (default "0.7.1")
	// +optional
	Version string `json:"version,omitempty"`
}

// VClusterStatus defines the observed state of VCluster
type VClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VCluster is the Schema for the vclusters API
type VCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VClusterSpec   `json:"spec,omitempty"`
	Status VClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VClusterList contains a list of VCluster
type VClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VCluster{}, &VClusterList{})
}
