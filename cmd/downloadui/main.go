package main

import (
	"github.com/lbernardo/slsui/internal/primary"
)

const url = "https://api.github.com/repos/lbernardo/slsui-front/releases/latest"

func main() {
	primary.NewWebCli(url)
}
