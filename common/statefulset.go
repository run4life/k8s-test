package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	apps_v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

// 读取statefulset yaml文件
func ReadStateFulSetYaml(filename string) apps_v1.StatefulSet {
	// 读取YAML
	statefulsetYaml, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return apps_v1.StatefulSet{}
	}

	// YAML转JSON
	deployJson, err := yaml.ToJSON(statefulsetYaml)
	if err != nil {
		fmt.Println(err)
		return apps_v1.StatefulSet{}
	}

	// JSON转struct
	var statefulset apps_v1.StatefulSet
	err = json.Unmarshal(deployJson, &statefulset)
	if err != nil {
		fmt.Println(err)
		return apps_v1.StatefulSet{}
	}
	return statefulset
}

func ApplyStatefulSet(clientset kubernetes.Clientset, statefulset apps_v1.StatefulSet) {
	var namespace string
	if statefulset.Namespace != "" {
		namespace = statefulset.Namespace
	} else {
		namespace = "default"
	}

	// 查询k8s是否有该statefulset
	_, err := clientset.AppsV1().StatefulSets(namespace).Get(context.TODO(), statefulset.Name, meta_v1.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			fmt.Println(err)
			return
		}
		// 不存在则创建
		_, err = clientset.AppsV1().StatefulSets(namespace).Create(context.TODO(), &statefulset, meta_v1.CreateOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
	} else { // 已存在则更新
		_, err = clientset.AppsV1().StatefulSets(namespace).Update(context.TODO(), &statefulset, meta_v1.UpdateOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Apply statefulset %s success!\n", statefulset.Name)
}

func DeleteStatefulSet(clientset kubernetes.Clientset, statefulset apps_v1.StatefulSet) {
	var namespace string
	if statefulset.Namespace != "" {
		namespace = statefulset.Namespace
	} else {
		namespace = "default"
	}
	err := clientset.AppsV1().StatefulSets(namespace).Delete(context.TODO(), statefulset.Name, meta_v1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("delete statefulset %s success!\n", statefulset.Name)
}

func GetStatefulSetStatus(clientset kubernetes.Clientset, deployment apps_v1.Deployment) (success bool, reasons []string, err error) {
	return success, reasons, err
}
