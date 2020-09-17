package secondary

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/lbernardo/slsui/pkg/web"
)

const filenameZip = "latest.zip"

type WebUIController struct {
	httpConnector web.HttpConnector
}

func NewWebUI(httpConnector web.HttpConnector) web.UIWebController {
	return &WebUIController{
		httpConnector: httpConnector,
	}
}

func (w *WebUIController) DownloadZip(url string) {
	os.RemoveAll("webui")
	var b web.Github
	content := w.httpConnector.DoGet(url)
	json.Unmarshal(content, &b)
	w.httpConnector.DownloadFile(b.Assets[0].BrowserDownloadURL, filenameZip)
	fmt.Println("Download ", b.Assets[0].BrowserDownloadURL)
	Unzip(filenameZip)
	fmt.Println("Update webui")
	os.Remove(filenameZip)

}

func Unzip(src string) {

	r, _ := zip.OpenReader(src)
	defer r.Close()
	dest := ""

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			panic(err)
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		rc, err := f.Open()
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			panic(err)
		}
	}
}
