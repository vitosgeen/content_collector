package scrappers

import "content_collector/internal/utils/smartproxy"

type IScrappers interface {
	SetProxy(proxyIp string)
	SetSmartProxy(smartProxy *smartproxy.SmartProxy)
	SetUserAgent(userAgent string)
	Scrap(url string) (*ScrapperData, error)
}

type ScrapperData struct {
	Url    string
	Length int
	Data   string
	Code   int
	Status string
}
