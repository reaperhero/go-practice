package cue

import (
	"cuelang.org/go/cue"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"testing"
)

const config = `
output: {
    apiVersion: "apps/v1"
    kind:       "Deployment"
    spec: {
        selector: matchLabels: {
            "app.oam.dev/component": parameter.name
        }
        template: {
            metadata: labels: {
                "app.oam.dev/component": parameter.name
            }
            spec: {
                containers: [{
                    name:  parameter.name
                    image: parameter.image
                }]
            }}}
}
parameter: {
    name: string
    image: string
}
`

type Parameter struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func TestName(t *testing.T) {
	cueTest1()
}

func cueTest1() {
	var r cue.Runtime
	instance, err := r.Compile("test", config)
	if err != nil {
		fmt.Println(err)
		return
	}
	p := Parameter{
		Name:  "mytest",
		Image: "nginx:v1",
	}
	instanceNew, err := instance.Fill(p, "parameter")
	if err != nil {
		fmt.Println(err)
		return
	}

	dataNew, _ := instanceNew.Value().MarshalJSON()
	fmt.Println(string(dataNew))
	d := new(appsv1.Deployment)
	v := instanceNew.Lookup("output")
	err = v.Decode(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.Marshal(d)
	fmt.Println(string(data))
}
