package test

import "k8s.io/client-go/kubernetes"

func PodJob(clientset kubernetes.Clientset) (success bool, reasons []string, err error) {
	return success, reasons, err
}

func PodTestCase(clientset kubernetes.Clientset) {
	PodJob(clientset)
}
