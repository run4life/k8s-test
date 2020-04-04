package test

import (
	"fmt"
	"k8s-test/common"

	"k8s.io/client-go/kubernetes"
)

func StateFulSetJob(clientset kubernetes.Clientset) {
	//首先创建pvc
	fmt.Println("Create One pvc")
	pvc := common.ReadPVCYaml("./conf/statefulset/mysql-pvc.yaml")
	common.ApplyPVC(clientset, pvc)
	//创建statefulset
	fmt.Println("Create One statefulset")
	statefulset := common.ReadStateFulSetYaml("./conf/statefulset/mysql-statefulset.yaml")
	common.ApplyStatefulSet(clientset, statefulset)
	//创建service
	fmt.Println("Create One service")
	service := common.ReadServiceYaml("./conf/statefulset/mysql-service.yaml")
	common.ApplyService(clientset, service)

}

func StateFulSetAutoPVCJob(clientset kubernetes.Clientset) {
	//创建statefulset
	fmt.Println("Create One autopvc statefulset")
	statefulset := common.ReadStateFulSetYaml("./conf/statefulset/mysql-autopvc-statefulset.yaml")
	common.ApplyStatefulSet(clientset, statefulset)
	//创建service
	fmt.Println("Create One service")
	service := common.ReadServiceYaml("./conf/statefulset/mysql-autopvc-service.yaml")
	common.ApplyService(clientset, service)

}

func StateFulSetTestCase(clientset kubernetes.Clientset) {
	//StateFulSetJob(clientset)
	StateFulSetAutoPVCJob(clientset)
}
