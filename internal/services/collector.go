package services

import (
	"fmt"

	"content_collector/internal/apperrors"
	"content_collector/internal/domain/model"
	"content_collector/internal/repository"
	"content_collector/internal/utils/scrappers"
	"content_collector/internal/utils/smartproxy"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const defaultUrlAddress = "https://www.google.com"

type ICollectorService interface {
	Collect(url string) (string, error)
	Close() error
	SetProxy(proxyIp string)
	CheckCollector() error
	DeleteOldCollectors() error
}

type CollectorService struct {
	Collector     model.Collector
	SmartProxy    smartproxy.ISmartProxyFile
	Scrapper      scrappers.IScrappers
	ProxyIp       string
	CollectorRepo repository.ICollectorRepository
}

func NewCollectorService(
	smartproxy smartproxy.ISmartProxyFile,
	pathChromeDriver string,
	portChromeDriver int,
	collectorRepo repository.ICollectorRepository,
	scrapper scrappers.IScrappers,
) ICollectorService {
	return &CollectorService{
		Collector: model.Collector{
			PathChromeDriver: pathChromeDriver,
			PortChromeDriver: portChromeDriver,
		},
		SmartProxy:    smartproxy,
		CollectorRepo: collectorRepo,
		Scrapper:      scrapper,
	}
}

func (c *CollectorService) Collect(url string) (string, error) {
	loadedCollector, err := c.CollectorRepo.GetByUrl(url)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.MongoCollectorRepositoryGetByIdErrNoDocuments.Code {
			return "", apperrors.ServicesCollectorCollectGetByUrlError.AppendMessage(err)
		}
	}

	if loadedCollector != nil {
		return loadedCollector.Data, nil
	}

	proxyIpSmart, err := c.SmartProxy.GetProxyRandomSmartProxy()
	if err != nil {
		return "", apperrors.ServicesCollectorCollectGetProxyRandomError.AppendMessage(err)
	}

	c.Scrapper.SetProxy(proxyIpSmart.String())
	c.Scrapper.SetSmartProxy(proxyIpSmart)

	html, err := c.Scrapper.Scrap(url)
	if err != nil {
		return "", apperrors.ServicesCollectorCollectScrapError.AppendMessage(err)
	}

	collectorCreate := &model.CollectorRepository{
		Url:    url,
		Data:   html,
		Status: model.CollectorRepositoryStatusActive,
	}
	err = c.CollectorRepo.Create(collectorCreate)
	if err != nil {
		return "", apperrors.ServicesCollectorCollectCreateError.AppendMessage(err)
	}

	return html, nil
}

func (c *CollectorService) Close() error {
	if err := c.Collector.WebDriver.Quit(); err != nil {
		return apperrors.ServicesCollectorCloseError.AppendMessage(err)
	}

	return nil
}

func (c *CollectorService) SetProxy(proxyIp string) {
	c.ProxyIp = proxyIp
}

func (c *CollectorService) CheckCollector() error {
	service, err := selenium.NewChromeDriverService(c.Collector.PathChromeDriver, c.Collector.PortChromeDriver)
	if err != nil {
		return apperrors.ServicesCollectorCheckCollectorNewChromeDriverServiceError.AppendMessage(err)
	}

	defer service.Stop() //nolint:errcheck

	caps := addCapabilities(c.ProxyIp)
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return apperrors.ServicesCollectorCheckCollectorNewRemoteError.AppendMessage(err)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		return apperrors.ServicesCollectorCheckCollectorMaximizeWindowError.AppendMessage(err)
	}

	if err := driver.Get(defaultUrlAddress); err != nil {
		return apperrors.ServicesCollectorCheckCollectorDriverGetError.AppendMessage(err)
	}

	return nil
}

func (c *CollectorService) DeleteOldCollectors() error {
	collectors, err := c.CollectorRepo.GetForDelete()
	if err != nil {
		return apperrors.ServicesCollectorDeleteOldCollectorsError.AppendMessage(err)
	}

	for _, collector := range collectors {
		err = c.CollectorRepo.Delete(collector.ID)
		if err != nil {
			return apperrors.ServicesCollectorDeleteOldCollectorsError.AppendMessage(err)
		}
	}

	return nil
}

func addCapabilities(proxyIp string) selenium.Capabilities {
	caps := selenium.Capabilities{}
	args := []string{}
	if proxyIp == "" {
		args = append(args, "--headless")
	} else {
		args = append(args, "--headless")
		args = append(args, "--proxy-server="+proxyIp)
		proxy := selenium.Proxy{
			Type: selenium.Manual,
			HTTP: "http://" + proxyIp,
		}
		caps.AddProxy(proxy)
	}

	// madify request headers to avoid detection
	args = append(args, []string{
		"--headless=new",
		"--user-agent=" + GetRandomUserAgent(),
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

	fmt.Println("proxyIp--------caps", caps)

	return caps
}
