package controllers

import (
	"context"

	consolev1 "github.com/openshift/api/console/v1"
	corev1 "k8s.io/api/core/v1"
	operatorv1 "github.com/openshift/api/operator/v1"
	schedulingv1 "k8s.io/api/scheduling/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	consoleuiv1alpha1 "github.com/kubevirt-ui/kubevirt-ui-operator/api/v1alpha1"
	"github.com/kubevirt-ui/kubevirt-ui-operator/internal/hcoversion"
	"github.com/kubevirt-ui/kubevirt-ui-operator/internal/names"
)

func (r *KubeVirtConsoleUIReconciler) mapConsolePluginToRequests(ctx context.Context, obj client.Object) []reconcile.Request {
	if cp, ok := obj.(*consolev1.ConsolePlugin); ok && cp.GetName() == names.ConsolePluginName {
		return r.enqueueOperandCRs(ctx)
	}
	return nil
}

func (r *KubeVirtConsoleUIReconciler) mapPriorityClassToRequests(ctx context.Context, obj client.Object) []reconcile.Request {
	if pc, ok := obj.(*schedulingv1.PriorityClass); ok && pc.GetName() == names.PriorityClass {
		return r.enqueueOperandCRs(ctx)
	}
	return nil
}

func (r *KubeVirtConsoleUIReconciler) mapConsoleToRequests(ctx context.Context, obj client.Object) []reconcile.Request {
	if console, ok := obj.(*operatorv1.Console); ok && console.GetName() == "cluster" {
		return r.enqueueOperandCRs(ctx)
	}
	return nil
}

func (r *KubeVirtConsoleUIReconciler) mapFeaturesConfigMapToRequests(ctx context.Context, obj client.Object) []reconcile.Request {
	cm, ok := obj.(*corev1.ConfigMap)
	if !ok || cm.GetNamespace() != r.Config.OperandNamespace || cm.GetName() != names.FeaturesConfigMap {
		return nil
	}
	return r.enqueueOperandCRs(ctx)
}

func (r *KubeVirtConsoleUIReconciler) mapHyperConvergedToRequests(ctx context.Context, obj client.Object) []reconcile.Request {
	if _, ok := obj.(*hcoversion.HyperConverged); ok {
		return r.enqueueOperandCRs(ctx)
	}
	return nil
}

func (r *KubeVirtConsoleUIReconciler) enqueueOperandCRs(ctx context.Context) []reconcile.Request {
	list, err := r.listConsoleUIsInOperandNamespace(ctx)
	if err != nil {
		return nil
	}
	out := make([]reconcile.Request, 0, len(list.Items))
	for i := range list.Items {
		cr := &list.Items[i]
		out = append(out, reconcile.Request{NamespacedName: client.ObjectKeyFromObject(cr)})
	}
	return out
}

func (r *KubeVirtConsoleUIReconciler) listConsoleUIsInOperandNamespace(ctx context.Context) (consoleuiv1alpha1.KubeVirtConsoleUIList, error) {
	var list consoleuiv1alpha1.KubeVirtConsoleUIList
	if err := r.List(ctx, &list, client.InNamespace(r.Config.OperandNamespace)); err != nil {
		log.FromContext(ctx).Error(err, "list KubeVirtConsoleUI in operand namespace", "namespace", r.Config.OperandNamespace)
		return list, err
	}
	return list, nil
}
