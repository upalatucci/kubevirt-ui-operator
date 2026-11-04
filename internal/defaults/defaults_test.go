package defaults

import (
	"testing"

	consoleuiv1alpha1 "github.com/kubevirt-ui/kubevirt-ui-operator/api/v1alpha1"
	"k8s.io/utils/ptr"
)

func TestApply_defaultPriorityClassCreate(t *testing.T) {
	spec := &consoleuiv1alpha1.KubeVirtConsoleUISpec{}
	Apply(spec)

	if !ptr.Deref(spec.PriorityClass.Create, false) {
		t.Fatal("expected priorityClass.create true by default")
	}
}

func TestApply_explicitFalsePriorityClassCreate(t *testing.T) {
	spec := &consoleuiv1alpha1.KubeVirtConsoleUISpec{
		PriorityClass: consoleuiv1alpha1.PriorityClassSpec{
			Create: new(false),
		},
	}
	Apply(spec)
	if ptr.Deref(spec.PriorityClass.Create, true) {
		t.Fatal("expected priorityClass.create to stay false")
	}
}
