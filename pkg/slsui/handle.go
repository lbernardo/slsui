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

func (s *SLSUI) Build(r RequestBuild) []byte {
	var sls ServerlessFramework
	sls.Service.Name = r.Provider.Name
	sls.Provider.Name = "aws"
	sls.Provider.Runtime = r.Provider.Runtime
	sls.Provider.Timeout = r.Provider.Timeout
	sls.Provider.MemorySize = r.Provider.MemorySize
	sls.Functions = map[string]SLSFunctions{}
	sls.Resources = map[string]interface{}{}

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
		case schedule:
			if n.Schedule.Type == "rate" {
				event["schedule"] = fmt.Sprintf("rate(%v %v)", n.Schedule.RateNum, n.Schedule.RatePeriod)
			} else {
				event["schedule"] = fmt.Sprintf("cron(%v)", n.Schedule.CronPeriod)
			}
		case sqs:
			event["sqs"] = n.Sqs.Value
		}

		sls.Functions[n.Name] = SLSFunctions{
			Handler: n.Handler,
			Events:  event,
		}
	}

	for _, n := range r.Dynamodb {

		sls.Resources[n.Name] = SLSDynamodbResource{
			Type: "AWS::DynamoDB::Table",
			Properties: map[string]interface{}{
				"TableName": n.TableName,
				"AttributeDefinitions": SLSAttributeDefinitions{
					AttributeName: n.AttributeDefinition.AttributeName,
					AttributeType: n.AttributeDefinition.AttributeType,
				},
				"KeySchema": SLSKeySchema{
					AttributeName: n.KeySchema.AttributeName,
					KeyType:       n.KeySchema.KeyType,
				},
				"ProvisionedThroughput": SLSProvisionedThroughput{
					ReadCapacityUnits:  n.Throughput.ReadCapacityUnits,
					WriteCapacityUnits: n.Throughput.WriteCapacityUnits,
				},
			},
		}
	}

	for _, n := range r.Sqs {
		sls.Resources[n.Name] = SLSSQSResource{
			Type: "AWS::SQS::Queue",
			Properties: SLSSQSProperties{
				QueueName: n.QueueName,
			},
		}
	}

	o, _ := yaml.Marshal(&sls)
	return o
}
