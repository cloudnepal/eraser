package util

import (
	eraserv1alpha1 "github.com/Azure/eraser/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func NeverOnCreate(_ event.CreateEvent) bool {
	return false
}

func NeverOnDelete(_ event.DeleteEvent) bool {
	return false
}

func NeverOnGeneric(_ event.GenericEvent) bool {
	return false
}

func NeverOnUpdate(_ event.UpdateEvent) bool {
	return false
}

func AlwaysOnCreate(_ event.CreateEvent) bool {
	return true
}

func AlwaysOnDelete(_ event.DeleteEvent) bool {
	return true
}

func AlwaysOnGeneric(_ event.GenericEvent) bool {
	return true
}

func AlwaysOnUpdate(_ event.UpdateEvent) bool {
	return true
}

func IsCompletedOrFailed(p eraserv1alpha1.JobPhase) bool {
	return (p == eraserv1alpha1.PhaseCompleted || p == eraserv1alpha1.PhaseFailed)
}

func FilterJobListByOwner(jobs []eraserv1alpha1.ImageJob, owner *metav1.OwnerReference) []eraserv1alpha1.ImageJob {
	ret := []eraserv1alpha1.ImageJob{}

	for i := range jobs {
		job := jobs[i]

		for j := range job.OwnerReferences {
			or := job.OwnerReferences[j]

			if or.UID == owner.UID {
				ret = append(ret, job)
				break // inner
			}
		}
	}

	return ret
}