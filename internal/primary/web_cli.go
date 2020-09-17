package primary

import "github.com/lbernardo/slsui/pkg/web"

func NewWebCli(url string) {
	c := new(Container)
	u := c.GetWebUIController()
	web.NewHandle(u, url)
}
