package slsui

type RequestBuild struct {
	Provider Provider      `json:"provider"`
	Lambda   []Lambda      `json:"lambda"`
	Dynamodb []interface{} `json:"dynamodb"`
	Sqs      []interface{} `json:"sqs"`
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
