package route

import (
	"encoding/json"

	routev1 "github.com/openshift/api/route/v1"
	"k8s.io/klog"
)

func IsAdmitted(route *routev1.Route) bool {
	admittedSet := false
	admittedValue := true
	buffer, err := json.Marshal(route)
	if err != nil {
		klog.Errorf("Error in serializing route information as JSON")
	} else {
		klog.Infof("route information: %v", string(buffer))
	}
	for _, ingress := range route.Status.Ingress {
		for _, condition := range ingress.Conditions {
			if condition.Type == "Admitted" {
				admittedSet = true
				if condition.Status != "True" {
					admittedValue = false
				}
			}
		}
	}
	return admittedSet && admittedValue
}
