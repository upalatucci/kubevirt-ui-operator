package defaults

import consoleuiv1alpha1 "github.com/kubevirt-ui/kubevirt-ui-operator/api/v1alpha1"

// Apply fills unset spec fields (pointer booleans).
func Apply(spec *consoleuiv1alpha1.KubeVirtConsoleUISpec) {
	if spec.PriorityClass.Create == nil {
		spec.PriorityClass.Create = new(true)
	}
}
