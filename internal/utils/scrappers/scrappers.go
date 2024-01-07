package scrappers

import "content_collector/internal/utils/smartproxy"

type IScrappers interface {
	SetProxy(proxyIp string)
	SetSmartProxy(smartProxy *smartproxy.SmartProxy)
	SetUserAgent(userAgent string)
	Scrap(url string) (string, error)
}
