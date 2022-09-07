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

package k8sgo

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	describe "k8s.io/kubectl/pkg/describe"
)

// ListAssociatedPodsWithPVC is a method to list associated pods
func ListAssociatedPodsWithPVC(pvcName, namespace string) ([]string, error) {
	logger := logGenerator(pvcName, namespace, pvcName)
	describer := describe.PersistentVolumeClaimDescriber{Interface: generateK8sClient()}
	peristentVolumeDetails, err := describer.Describe(namespace, pvcName, describe.DescriberSettings{})
	if err != nil {
		return []string{}, err
	}
	logger.Info("Updated Size of persistent volume", "PVC Info", peristentVolumeDetails)
	return []string{}, nil
}

// ResizePersistentVolume is a method to resize peristent volume
func ResizePersistentVolume(pvcName, namespace string, size int) error {
	logger := logGenerator(pvcName, namespace, pvcName)
	pvcInfo, err := generateK8sClient().CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvcName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	newSize, err := resource.ParseQuantity(fmt.Sprintf("%vGi", size))
	if err != nil {
		return nil
	}
	pvcInfo.Spec.Resources.Requests[corev1.ResourceStorage] = newSize

	_, err = generateK8sClient().CoreV1().PersistentVolumeClaims(namespace).Update(context.TODO(), pvcInfo, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	logger.Info("Updated Size of persistent volume", "Size", fmt.Sprintf("%vGi", size))
	return nil
}
