package test

import "k8s.io/client-go/kubernetes"

func ServiceJob(clientset kubernetes.Clientset) (success bool, reasons []string, err error) {
	return success, reasons, err
}

func ServiceTestCase(clientset kubernetes.Clientset) {
	ServiceJob(clientset)
}
