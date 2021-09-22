/*
Copyright 2021.

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

package network

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	networkv1alpha1 "github.com/fennec-project/network-profile-operator/apis/network/v1alpha1"
)

// NetworkProfileReconciler reconciles a NetworkProfile object
type NetworkProfileReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=network.fennecproject.io,resources=networkprofiles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=network.fennecproject.io,resources=networkprofiles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=network.fennecproject.io,resources=networkprofiles/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NetworkProfile object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *NetworkProfileReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Get NetworkProfiles

	// Log Info networkprofile name in use
	// Log Debug whole network profile values

	// check finalizers and deletion timestamps

	// Log Debug finalizers and deletion timestamps if any

	// List pods with annotations AND for each pod:

	// Check if name is in status if not write it and proceed else skip
	// Log info if pod is or not already configured

	// check on networkConfiguration - is configuration present?
	// what is the status on the system not CR.
	// set networkConfiguration for pod
	// Log Info networkConfiguration name for pod name
	// Log Debug networkConfiguration full parameters

	// Check on policy - is it configured?
	// set network policy
	// Log Info networkPolicy for pod all fields

	// Check on QoS marking is it done for the Pod's interface?
	// set network QoS marking
	// Log Info networkProfileClass for pod name
	// Log Debug all values from networkProfileClass

	//  write to status: Conditions and Pods listed for that networkProfile

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NetworkProfileReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkv1alpha1.NetworkProfile{}).
		Complete(r)
}
