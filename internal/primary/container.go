package primary

import (
	"github.com/lbernardo/slsui/internal/secondary"
	"github.com/lbernardo/slsui/pkg/web"
)

type Container struct {
	WebUIController web.UIWebController
	HttpConnector   web.HttpConnector
}

func (c *Container) GetWebUIController() web.UIWebController {
	if c.WebUIController == nil {
		c.WebUIController = secondary.NewWebUI(c.GetHttpConnector())
	}
	return c.WebUIController
}

func (c *Container) GetHttpConnector() web.HttpConnector {
	if c.HttpConnector == nil {
		c.HttpConnector = secondary.NewHttpConnector()
	}
	return c.HttpConnector
}
