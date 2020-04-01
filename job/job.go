package job

import (
	"fmt"
)

func Run(param []interface{}) {
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
	c := cTest(paramMap)
	if resultChan != nil {
		resultChan <- c
	} else {
		fmt.Println(c)
	}

	//time.Sleep(time.Millisecond*10);
}
func cTest(paramMap map[string]interface{}) map[string]interface{} {
	paramMap["result"] = "success"
	return paramMap
}