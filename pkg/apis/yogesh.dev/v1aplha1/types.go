package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FilePath",type=string,JSONPath=`.status.FilePath`
// +kubebuilder:printcolumn:name="InSync",type=bool,JSONPath=`.status.InSync`
type Permsync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FileSpec   `json:"spec,omitempty"`
	Status FileStatus `json:"status,omitempty"`
}

type FileStatus struct {
	FilePath string `json:"filePath"`
	InSync   bool   `json:"inSync"`
}

type FileSpec struct {
	Name    string `json:"name,omitempty"`
	Path    string `json:"path,omitempty"`
	Version string `json:"version,omitempty"`

	Permission PermissionBlock `json:"permissionBlock,omitempty"`
}

type PermissionBlock struct {
	Read    string `json:"read,omitempty"`
	Write   string `json:"write,omitempty"`
	Execute string `json:"execute,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PermsyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Permsync `json:"items,omitempty"`
}
