package web

func NewHandle(u UIWebController, url string) {
	u.DownloadZip(url)
}
