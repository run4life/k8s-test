package test

import (
	"k8s-test/common"
	"strconv"
	"sync"

	"k8s.io/client-go/kubernetes"
)

var wg sync.WaitGroup

func stressParaDeploymentJob(clientset kubernetes.Clientset) {
	sum := 0
	deployment := common.ReadDeploymentYaml("./conf/nginx.yaml")
	n := deployment.ObjectMeta.Name
	for {
		wg.Add(1)
		name := n + strconv.Itoa(sum)
		go func(name string) {
			defer wg.Done()
			deployment.ObjectMeta.Name = name
			common.ApplyDeployment(clientset, deployment)
		}(name)
		if sum >= 5 {
			break
		}
		sum++
	}
	wg.Wait()
	//success, reasons, err := common.GetDeploymentStatus(clientset, deployment)
}

func stressParaStatefulSetJob(clientset kubernetes.Clientset) {
	sum := 0
	statefulset := common.ReadStateFulSetYaml("./conf/statefulset/mysql-autopvc-statefulset.yaml")
	n := statefulset.ObjectMeta.Name
	for {
		wg.Add(1)
		name := n + strconv.Itoa(sum)
		go func(name string) {
			defer wg.Done()
			statefulset.ObjectMeta.Name = name
			common.ApplyStatefulSet(clientset, statefulset)
		}(name)
		if sum >= 5 {
			break
		}
		sum++
	}
	wg.Wait()
	//success, reasons, err := common.GetDeploymentStatus(clientset, deployment)
}

func stressJob(clientset kubernetes.Clientset) {
	sum := 0
	deployment := common.ReadDeploymentYaml("./conf/nginx.yaml")
	n := deployment.ObjectMeta.Name
	for {
		name := n + strconv.Itoa(sum)
		deployment.ObjectMeta.Name = name
		common.ApplyDeployment(clientset, deployment)

		if sum >= 5 {
			break
		}
		sum++
	}
	//success, reasons, err := common.GetDeploymentStatus(clientset, deployment)
}

func StressTestCase(clientset kubernetes.Clientset) {
	//stressParaDeploymentJob(clientset)
	stressParaStatefulSetJob(clientset)

}
