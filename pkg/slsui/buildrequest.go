package slsui

type RequestBuild struct {
	Provider Provider      `json:"provider"`
	Lambda   []Lambda      `json:"lambda"`
	Dynamodb []Dynamodb    `json:"dynamodb"`
	Sqs      []SqsResource `json:"sqs"`
}

type Dynamodb struct {
	Name                string              `json:"name"`
	TableName           string              `json:"tableName"`
	AttributeDefinition AttributeDefinition `json:"attributeDefinition"`
	KeySchema           KeySchema           `json:"keySchema"`
	Throughput          Throughput          `json:"throughput"`
}

type AttributeDefinition struct {
	AttributeName string `json:"attributeName"`
	AttributeType string `json:"attributeType"`
}

type KeySchema struct {
	AttributeName string `json:"attributeName"`
	KeyType       string `json:"keyType"`
}

type Throughput struct {
	ReadCapacityUnits  int64 `json:"readCapacityUnits"`
	WriteCapacityUnits int64 `json:"writeCapacityUnits"`
}

type Lambda struct {
	Event      string     `json:"event"`
	Name       string     `json:"name"`
	Handler    string     `json:"handler"`
	Apigateway Apigateway `json:"apigateway"`
	S3         S3         `json:"s3"`
	Schedule   Schedule   `json:"schedule"`
	Sqs        Sqs        `json:"sqs"`
}

type Apigateway struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type S3 struct {
	Bucket     string `json:"bucket"`
	Permission string `json:"permission"`
}

type Schedule struct {
	Type       string `json:"type"`
	RateNum    int64  `json:"rateNum"`
	RatePeriod string `json:"ratePeriod"`
	CronPeriod string `json:"cronPeriod"`
}

type Sqs struct {
	Value string `json:"value"`
}

type Provider struct {
	Name       string `json:"name"`
	Region     string `json:"region"`
	Runtime    string `json:"runtime"`
	MemorySize int64  `json:"memorySize"`
	Timeout    int64  `json:"timeout"`
}

type SqsResource struct {
	Name      string `json:"name"`
	QueueName string `json:"queueName"`
}
