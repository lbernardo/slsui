package web

type Github struct {
	Assets []Asset `json:"assets"`
}

type Asset struct {
	BrowserDownloadURL string `json:"browser_download_url"`
}
