package model

type Content struct {
	Url           string
	Body          string
	StatusCode    int
	ContentType   string
	ContentLength int64
	Headers       map[string][]string
	ProxyIp       string
}
