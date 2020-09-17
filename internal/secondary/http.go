package secondary

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/lbernardo/slsui/pkg/web"
)

type HttpConnector struct {
	client *http.Client
}

func NewHttpConnector() web.HttpConnector {
	c := &http.Client{}
	return &HttpConnector{
		client: c,
	}
}

func (h *HttpConnector) DoGet(url string) []byte {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, _ := h.client.Do(req)
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	return buf.Bytes()
}

func (h *HttpConnector) DownloadFile(url, filename string) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, _ := h.client.Do(req)
	defer res.Body.Close()
	out, _ := os.Create(filename)
	defer out.Close()
	io.Copy(out, res.Body)
}
