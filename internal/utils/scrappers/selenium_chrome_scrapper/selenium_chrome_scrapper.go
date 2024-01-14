package seleniumchromescrapper

// TODO: add support for proxyip with authentication like as smartproxy

import (
	"fmt"
	"net/http"

	"content_collector/internal/apperrors"
	"content_collector/internal/utils/scrappers"
	"content_collector/internal/utils/smartproxy"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type SeleniumChromeScrapper struct {
	PathChromeDriver string
	PortChromeDriver int
	smartproxy       *smartproxy.SmartProxy
	proxyIp          string
	userAgent        string
}

func NewSeleniumChromeScrapper(pathChromeDriver string, portChromeDriver int) scrappers.IScrappers {
	return &SeleniumChromeScrapper{
		PathChromeDriver: pathChromeDriver,
		PortChromeDriver: portChromeDriver,
	}
}

func (s *SeleniumChromeScrapper) SetProxy(proxyIp string) {
	s.proxyIp = proxyIp
}

func (s *SeleniumChromeScrapper) SetSmartProxy(smartProxy *smartproxy.SmartProxy) {
	s.smartproxy = smartProxy
}

func (s *SeleniumChromeScrapper) SetUserAgent(userAgent string) {
	s.userAgent = userAgent
}

func (s *SeleniumChromeScrapper) Scrap(url string) (*scrappers.ScrapperData, error) {
	service, err := selenium.NewChromeDriverService(s.PathChromeDriver, s.PortChromeDriver)
	if err != nil {
		return nil, apperrors.SeleniumChromeScrapperScrapNewChromeDriverServiceError.AppendMessage(err)
	}
	defer service.Stop() //nolint:errcheck

	caps := s.addCapabilities()

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, apperrors.SeleniumChromeScrapperScrapNewRemoteError.AppendMessage(err)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		return nil, apperrors.SeleniumChromeScrapperScrapMaximizeWindow.AppendMessage(err)
	}

	err = driver.Get(url)
	if err != nil {
		return nil, apperrors.SeleniumChromeScrapperScrapDriverGet.AppendMessage(err)
	}

	html, err := driver.PageSource()
	if err != nil {
		return nil, apperrors.SeleniumChromeScrapperScrapPageSource.AppendMessage(err)
	}

	scrapperData := &scrappers.ScrapperData{
		Url:    url,
		Length: len(html),
		Data:   html,
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	return scrapperData, nil
}

func (s SeleniumChromeScrapper) addCapabilities() selenium.Capabilities {
	caps := selenium.Capabilities{}
	args := []string{}
	if s.proxyIp == "" {
		args = append(args, "--headless")
	} else {
		args = append(args, "--headless")
		args = append(args, "--proxy-server="+s.smartproxy.Host+":"+s.smartproxy.Port)
		proxy := selenium.Proxy{
			Type: selenium.Manual,
			HTTP: s.smartproxy.String(),
		}
		caps.AddProxy(proxy)
	}

	// madify request headers to avoid detection
	args = append(args, []string{
		"--headless=new",
		"--user-agent=" + s.userAgent,
		"--disable-blink-features=AutomationControlled",
		"--disable-dev-shm-usage",
		"--disable-gpu",
		"--no-sandbox",
		"--disable-features=VizDisplayCompositor",
		"--disable-features=IsolateOrigins,site-per-process",
		"--disable-site-isolation-trials",
		"--disable-extensions",
		"--disable-web-security",
		"--disable-features=site-per-process",
		"--disable-features=NetworkService",
		"--disable-features=NetworkServiceInProcess",
	}...)

	caps.AddChrome(chrome.Capabilities{
		Args: args,
	})

	fmt.Println("caps", caps)

	return caps
}

func (s *SeleniumChromeScrapper) Decode(htmlString string) (string, error) {
	return htmlString, nil
}
