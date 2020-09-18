package primary

import (
	"encoding/json"
	"net/http"

	"github.com/lbernardo/slsui/pkg/slsui"
)

type ServiceSLSUI struct {
}

func NewService() *ServiceSLSUI {
	return &ServiceSLSUI{}
}

func (s *ServiceSLSUI) Build(w http.ResponseWriter, r *http.Request) {
	var reqBuild slsui.RequestBuild
	json.NewDecoder(r.Body).Decode(&reqBuild)
	n := slsui.NewSLSUI()
	res := n.Build(reqBuild)
	w.Header().Set("Content-Disposition", "attachment; filename=serverless.yml")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(res)
}
