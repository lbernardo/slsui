package slsui

type ServerlessFramework struct {
	Service   SLSService              `yaml:"service"`
	Provider  SLSProvider             `yaml:"provider"`
	Functions map[string]SLSFunctions `yaml:"functions"`
}

type SLSService struct {
	Name string `yaml:"name"`
}

type SLSProvider struct {
	Name       string `yaml:"name"`
	Runtime    string `yaml:"runtime"`
	MemorySize int64  `yaml:"memorySize"`
	Timeout    int64  `yaml:"timeout"`
}

type SLSFunctions struct {
	Handler string                 `yaml:"handler"`
	Events  map[string]interface{} `yaml:"events,omitempty"`
}

type SLSHttpEvent struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type SLSS3Event struct {
	Bucket string `yaml:"bucket"`
	Event  string `yaml:"event"`
}
