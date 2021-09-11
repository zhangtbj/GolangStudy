package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`
		{
			"crn:v1:bluemix:public:containers-kubernetes:eu-de:a/xxx:xxx": {
				"type":"containers-kubernetes",
				"subtype":"k8s-cluster"
			},
			"crn:v1:bluemix:public:cf:us-south:a/xxx:xxx": {
				"type":"cf",
				"subtype":"cf-application"
			}
		}`)

	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		fmt.Printf(err.Error())
	}

	resourceMap := f.(map[string]interface{})

	crnName := "crn:v1:bluemix:public:containers-kubernetes:eu-de:a/xxx:xxx"
	resource := resourceMap[crnName]
	if resource != nil {
		types := resource.(map[string]interface{})
		fmt.Printf("Type is %s\n", types["type"])
		fmt.Printf("Subtype is %s\n", types["subtype"])
	} else {
		fmt.Printf("No this resource")
	}
}
