package deployment

import (
	"fmt"
	"k8s-test/common"

	"k8s.io/client-go/kubernetes"
)

func TestDeployment(clientset kubernetes.Clientset) {
	deployment := common.ReadDeploymentYaml("./conf/nginx.yaml")
	sum := 0
	for {
		//deployment := common.ReadDeploymentYaml("./conf/nginx.yaml")
		sum++
		if sum > 100 {
			break
		}
		fmt.Println(sum)
		//name := deployment.ObjectMeta.Name
		//deployment.ObjectMeta.Name = name + strconv.Itoa(sum)
		//fmt.Println(deployment.ObjectMeta.Name)
		//common.ApplyDeployment(clientset, deployment)
	}
	//common.ApplyDeployment(clientset, deployment)
	success, reasons, err := common.GetDeploymentStatus(clientset, deployment)
	fmt.Println(success)
	fmt.Println(reasons)
	fmt.Println(err)
	common.PrintDeploymentStatus(clientset, deployment)
}
