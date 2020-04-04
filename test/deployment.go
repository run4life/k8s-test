package test

import (
	"fmt"
	"k8s-test/common"
	"time"

	apps_v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

func checkdeploymentstatus(clientset kubernetes.Clientset, deployment apps_v1.Deployment) {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 60 {
				fmt.Println("check deployment status timeout!")
				goto Loop
			}
		}
		fmt.Println("Check deployment status")
		success, _, _ := common.GetDeploymentStatus(clientset, deployment)
		if success {
			fmt.Println("Deployment status is OK!")
			goto Loop
		}
	}
Loop:
	fmt.Println("Check finish")
}

func deploymentJob(clientset kubernetes.Clientset) {
	fmt.Println("Create One deployment")
	deployment := common.ReadDeploymentYaml("./conf/nginx.yaml")
	common.ApplyDeployment(clientset, deployment)
	checkdeploymentstatus(clientset, deployment)
}

func DeploymentTestCase(clientset kubernetes.Clientset) {
	deploymentJob(clientset)

}
