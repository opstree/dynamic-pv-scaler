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

package controllers

import (
	"context"
	"time"

	"dynamic-pv-scaler/k8sgo"
	"dynamic-pv-scaler/promutils"
	"dynamic-pv-scaler/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	pvcv1 "dynamic-pv-scaler/api/v1"
)

// PersistentVolumeScalerReconciler reconciles a PersistentVolumeScaler object
type PersistentVolumeScalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=pvc.scaler.opstreelabs.in,resources=persistentvolumescalers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pvc.scaler.opstreelabs.in,resources=persistentvolumescalers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=pvc.scaler.opstreelabs.in,resources=persistentvolumescalers/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=pods;persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *PersistentVolumeScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	instance := &pvcv1.PersistentVolumeScaler{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{RequeueAfter: time.Second * 10}, nil
		}
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err := controllerutil.SetControllerReference(instance, instance, r.Scheme); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}

	for _, pvcInfo := range instance.Spec.PVCRefName {
		params := promutils.QueryInputs{
			Namespace:        instance.Namespace,
			PersistentVolume: pvcInfo,
		}
		currentPercentage, err := promutils.GetPersistentVolumeUsagePercentage(params)
		if err != nil {
			return ctrl.Result{RequeueAfter: time.Second * 10}, err
		}
		if (100 - currentPercentage.Value) <= instance.Spec.ScaleParameters.ThresholdValue {
			totalCapacity, err := promutils.GetPersistentVolumeTotalCapacity(params)
			if err != nil {
				return ctrl.Result{RequeueAfter: time.Second * 10}, err
			}
			_ = utils.CalculateUpdatedSize(totalCapacity.Value, instance.Spec.ScaleParameters.ScaleValue)
			_, err = k8sgo.ListAssociatedPodsWithPVC(pvcInfo, instance.Namespace)
			if err != nil {
				return ctrl.Result{RequeueAfter: time.Second * 10}, err
			}
		}
	}
	return ctrl.Result{RequeueAfter: time.Second * 10}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PersistentVolumeScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pvcv1.PersistentVolumeScaler{}).
		Complete(r)
}
