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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PersistentVolumeScalerSpec defines the desired state of PersistentVolumeScaler
type PersistentVolumeScalerSpec struct {
	PVCRefName      []string        `json:"pvcRefName"`
	ScaleParameters ScaleParameters `json:"scaleParameters"`
	RestartPods     *bool           `json:"restartPods,omitempty"`
}

// ScaleParameters defines the scaling parameters for PersistentVolume
type ScaleParameters struct {
	ThresholdValue int `json:"thresholdValuePercentage"`
	ScaleValue     int `json:"scaleValuePercentage"`
}

// PersistentVolumeScalerStatus defines the observed state of PersistentVolumeScaler
type PersistentVolumeScalerStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PersistentVolumeScaler is the Schema for the persistentvolumescalers API
type PersistentVolumeScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PersistentVolumeScalerSpec   `json:"spec,omitempty"`
	Status PersistentVolumeScalerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PersistentVolumeScalerList contains a list of PersistentVolumeScaler
type PersistentVolumeScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PersistentVolumeScaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PersistentVolumeScaler{}, &PersistentVolumeScalerList{})
}
