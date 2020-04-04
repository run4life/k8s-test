package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

// 读取statefulset yaml文件
func ReadPVCYaml(filename string) core_v1.PersistentVolumeClaim {
	// 读取YAML
	pvcYaml, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return core_v1.PersistentVolumeClaim{}
	}

	// YAML转JSON
	deployJson, err := yaml.ToJSON(pvcYaml)
	if err != nil {
		fmt.Println(err)
		return core_v1.PersistentVolumeClaim{}
	}

	// JSON转struct
	var pvc core_v1.PersistentVolumeClaim
	err = json.Unmarshal(deployJson, &pvc)
	if err != nil {
		fmt.Println(err)
		return core_v1.PersistentVolumeClaim{}
	}
	return pvc
}

func ApplyPVC(clientset kubernetes.Clientset, new_pvc core_v1.PersistentVolumeClaim) {

	// 查询k8s是否有该pvc
	_, err := clientset.CoreV1().PersistentVolumeClaims("default").Get(context.TODO(), new_pvc.Name, meta_v1.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			fmt.Println(err)
			return
		}
		// 不存在则创建
		_, err = clientset.CoreV1().PersistentVolumeClaims("default").Create(context.TODO(), &new_pvc, meta_v1.CreateOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
	} else { // 已存在则更新
		_, err = clientset.CoreV1().PersistentVolumeClaims("default").Update(context.TODO(), &new_pvc, meta_v1.UpdateOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("apply pvc %s success!\n", new_pvc.Name)
}