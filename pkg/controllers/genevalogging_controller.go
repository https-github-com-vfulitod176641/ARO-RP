package controllers

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	securityclient "github.com/openshift/client-go/security/clientset/versioned"
	"github.com/sirupsen/logrus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	aro "github.com/Azure/ARO-RP/operator/apis/aro.openshift.io/v1alpha1"
	"github.com/Azure/ARO-RP/pkg/genevalogging"
	aroclient "github.com/Azure/ARO-RP/pkg/util/aro-operator-client/clientset/versioned/typed/aro.openshift.io/v1alpha1"
)

// GenevaloggingReconciler reconciles a Cluster object
type GenevaloggingReconciler struct {
	Kubernetescli kubernetes.Interface
	Securitycli   securityclient.Interface
	AROCli        aroclient.AroV1alpha1Interface
	Log           *logrus.Entry
	Scheme        *runtime.Scheme
}

// +kubebuilder:rbac:groups=aro.openshift.io,resources=clusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.v1,resources=daemonsets,verbs=get;update;patch;create
// +kubebuilder:rbac:groups="",resources=namespaces;serviceaccounts;configmaps,verbs=get;create;update

func (r *GenevaloggingReconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	if request.Name != aro.SingletonClusterName || request.Namespace != OperatorNamespace {
		return reconcile.Result{}, nil
	}
	r.Log.Info("Reconsiling genevalogging deployment")

	ctx := context.TODO()
	instance, err := r.AROCli.Clusters(request.Namespace).Get(request.Name, v1.GetOptions{})
	if err != nil {
		// Error reading the object or not found - requeue the request.
		return reconcile.Result{}, err
	}

	newCert, err := r.Kubernetescli.CoreV1().Secrets(instance.Spec.GenevaLogging.Namespace).Get("certificates", v1.GetOptions{})
	if err != nil && apierrors.IsNotFound(err) {
		// copy the certificates from our namespace into the genevalogging one.
		certs, err := r.Kubernetescli.CoreV1().Secrets(OperatorNamespace).Get("certificates", v1.GetOptions{})
		if err != nil {
			r.Log.Errorf("Error reading the certificates secret: %v", err)
			return reconcile.Result{}, err
		}
		newCert = certs.DeepCopy()
		newCert.Namespace = instance.Spec.GenevaLogging.Namespace
		newCert.ResourceVersion = ""
	}

	gl := genevalogging.NewForOperator(r.Log, &instance.Spec, r.Kubernetescli, r.Securitycli, newCert)
	err = gl.CreateOrUpdate(ctx)
	if err != nil {
		r.Log.Error(err)
		return reconcile.Result{}, err
	}

	r.Log.Info("done, requeueing")
	return ReconcileResultRequeue, nil
}

func (r *GenevaloggingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&aro.Cluster{}).
		Complete(r)
}
