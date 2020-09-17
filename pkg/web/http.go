package web

type HttpConnector interface {
	DoGet(url string) []byte
	DownloadFile(url, filename string)
}
