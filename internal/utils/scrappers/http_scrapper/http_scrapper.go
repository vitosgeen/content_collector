package httpscrapper

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"content_collector/internal/utils/scrappers"
	"content_collector/internal/utils/smartproxy"
)

const (
	timeLimit = 30 * time.Second
)

type HttpScpaper struct {
	smartproxyIp *smartproxy.SmartProxy
	proxyIp      string
	userAgent    string
}

func NewHttpScpaper() scrappers.IScrappers {
	return &HttpScpaper{}
}

func (h *HttpScpaper) SetProxy(proxyIp string) {
	h.proxyIp = proxyIp
}

func (h *HttpScpaper) SetSmartProxy(smartProxy *smartproxy.SmartProxy) {
	h.smartproxyIp = smartProxy
}

func (h *HttpScpaper) SetUserAgent(userAgent string) {
	h.userAgent = userAgent
}

func (h *HttpScpaper) Scrap(urlTarget string) (*scrappers.ScrapperData, error) {
	proxyUrl := &url.URL{
		Scheme: "http",
		User:   url.UserPassword(h.smartproxyIp.Username, h.smartproxyIp.Password),
		Host:   h.smartproxyIp.Host + ":" + h.smartproxyIp.Port,
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	client := &http.Client{
		Timeout:   timeLimit,
		Transport: transport,
	}

	requestTarget, err := http.NewRequest("GET", urlTarget, nil)
	if err != nil {
		return nil, err
	}

	// modify request headers to avoid detection
	requestTarget.Header.Add("Accept-Charset", "UTF-8")
	requestTarget.Header.Add("User-Agent", h.userAgent)
	requestTarget.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	requestTarget.Header.Add("Accept-Language", "en-US,en;q=0.5")
	requestTarget.Header.Add("Connection", "keep-alive")
	requestTarget.Header.Add("Upgrade-Insecure-Requests", "1")
	requestTarget.Header.Add("Cache-Control", "no-cache")
	requestTarget.Header.Add("TE", "Trailers")

	// add utf-8 encoding
	requestTarget.Header.Add("Content-Type", "text/html; charset=utf-8")

	// add referer
	requestTarget.Header.Add("Referer", urlTarget)

	// add origin
	requestTarget.Header.Add("Origin", urlTarget)

	// add host
	requestTarget.Header.Add("Host", urlTarget)
	requestTarget.Header.Add("Alt-Used", urlTarget)

	// simulate browser
	requestTarget.Header.Add("Sec-Fetch-Dest", "document")
	requestTarget.Header.Add("Sec-Fetch-Mode", "navigate")
	requestTarget.Header.Add("Sec-Fetch-Site", "cross-site")
	requestTarget.Header.Add("Sec-Fetch-User", "?1")
	requestTarget.Header.Add("Pragma", "no-cache")

	// simulate javascript
	requestTarget.Header.Add("Sec-GPC", "1")

	requestTarget.Header.Add("Cookie", "__cf_bm=KEBcS5Kl48hyiVOLPsKYgU1IBmfz44LbdeLwVD9vCrw-1704568358-1-AcprFpyOg1tax/DPqtBvFQEH0WddTNiX+xK6wT/nftN4IKMAIOiSONVDThEvbivqCIhHh4uFDaQYSxNP6HNZRFI=; _ga_GZ1WW2KT3Q=GS1.1.1704568359.1.1.1704568376.43.0.0; _ga=GA1.2.80261033.1704568360; _gid=GA1.2.296099666.1704568360")

	response, err := client.Do(requestTarget)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// if response.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("status code: %d", response.StatusCode)
	// }
	bodyString := fmt.Sprintf("%s\n", body)

	scrapperData := &scrappers.ScrapperData{
		Url:    urlTarget,
		Length: len(bodyString),
		Data:   bodyString,
		Code:   response.StatusCode,
		Status: response.Status,
	}

	return scrapperData, nil
}

func (h *HttpScpaper) Decode(htmlString string) (string, error) {
	return htmlString, nil
}
