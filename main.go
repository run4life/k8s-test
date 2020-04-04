package main

import (
	"fmt"
	"k8s-test/common"
	"k8s-test/test"
)

//var poolOne worker.WorkPool

//func init() {
//worker.StartPool(6)
//	poolOne.InitPool()
//}

func main() {
	//初始化k8s连接
	clientset, err := common.InitClient("./conf/admin.conf")
	if err != nil {
		fmt.Println(err)
		return
	}

	//var resultChan = make(chan interface{}, 200)
	//开启deployment job测试
	//var paramMap = make(map[string]interface{})
	//deployment := common.ReadDeploymentYaml("./conf/nginx.yaml")
	//paramMap["clientset"] = clientset
	//paramMap["deployment"] = deployment
	//poolOne.Run(job.DeploymentJobRun, paramMap, resultChan)
	//for ret := range resultChan {
	//	fmt.Print(ret)
	//}
	//poolOne.Stop()

	//开启deployment测试
	//test.DeploymentTestCase(clientset)
	//开启service测试
	//test.ServiceTestCase(clientset)
	//开启pod测试
	//test.PodTestCase(clientset)
	//开启statefulset测试
	test.StateFulSetTestCase(clientset)
	//压力测试
	//test.StressTestCase(clientset)
}
