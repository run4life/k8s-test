package main

import (
	"fmt"
	"k8s-test/common"
	"k8s-test/deployment"
	"k8s-test/job"
	"k8s-test/worker"
)

var poolOne worker.WorkPool

func init() {
	worker.StartPool(6)
	poolOne.InitPool()
}

func main() {
	var resultChan = make(chan interface{}, 200)
	for i := 0; i < 5; i++ {
		var paramMap = make(map[string]interface{})
		paramMap["id"] = i
		paramMap["a"] = 1 + i
		paramMap["result"] = ""
		poolOne.Run(job.Run, paramMap, resultChan)
		fmt.Println(<-resultChan)
	}
	poolOne.Stop()

	clientset, err := common.InitClient("./conf/admin.conf")
	if err != nil {
		fmt.Println(err)
		return
	}
	deployment.TestDeployment(clientset)
}
