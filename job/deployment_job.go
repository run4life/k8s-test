package job

import (
	"fmt"
	"k8s-test/common"

	apps_v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

func DeploymentJobRun(param []interface{}) {
	var clientset kubernetes.Clientset
	var deployment apps_v1.Deployment
	var paramMap map[string]interface{}
	var resultChan chan interface{}
	for _, val := range param {
		switch v := val.(type) {
		case map[string]interface{}:
			paramMap = v
		case chan interface{}:
			resultChan = v
		}
	}
	clientset = paramMap["clientset"].(kubernetes.Clientset)
	deployment = paramMap["deployment"].(apps_v1.Deployment)
	fmt.Println("start a deployment job")
	c := DeploymentJob(clientset, deployment)
	if resultChan != nil {
		resultChan <- c
	} else {
		fmt.Println(c)
	}

	//time.Sleep(time.Millisecond*10);
}
func DeploymentJob(clientset kubernetes.Clientset, deployment apps_v1.Deployment) int {
	fmt.Println("Create One deployment")
	//deployment = common.ReadDeploymentYaml("./conf/nginx.yaml")
	common.ApplyDeployment(clientset, deployment)
	return 1
}
