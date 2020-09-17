package slsui

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

const (
	apigateway = "apigateway"
	s3         = "s3"
	schedule   = "schedule"
	sqs        = "sqs"
)

type SLSUI struct{}

func NewSLSUI() *SLSUI {
	return &SLSUI{}
}

func (s *SLSUI) Build(r RequestBuild) {
	var sls ServerlessFramework
	sls.Service.Name = r.Provider.Name
	sls.Provider.Name = "aws"
	sls.Provider.Runtime = r.Provider.Runtime
	sls.Provider.Timeout = r.Provider.Timeout
	sls.Provider.MemorySize = r.Provider.MemorySize
	sls.Functions = map[string]SLSFunctions{}

	for _, n := range r.Lambda {

		event := make(map[string]interface{})
		switch n.Event {
		case apigateway:
			event["http"] = SLSHttpEvent{
				Path:   n.Apigateway.Path,
				Method: n.Apigateway.Method,
			}
		case s3:
			event["s3"] = SLSS3Event{
				Bucket: n.S3.Bucket,
				Event:  n.S3.Permission,
			}
		}

		sls.Functions[n.Name] = SLSFunctions{
			Handler: n.Handler,
			Events:  event,
		}
	}

	o, _ := yaml.Marshal(&sls)
	fmt.Println(string(o))
}
