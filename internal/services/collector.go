package services

import (
	"content_collector/internal/apperrors"
	"content_collector/internal/domain/model"
	"content_collector/internal/repository"

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
	ProxyIp       string
	CollectorRepo repository.ICollectorRepository
}

func NewCollectorService(pathChromeDriver string, portChromeDriver int, collectorRepo repository.ICollectorRepository) ICollectorService {
	return &CollectorService{
		Collector: model.Collector{
			PathChromeDriver: pathChromeDriver,
			PortChromeDriver: portChromeDriver,
		},
		CollectorRepo: collectorRepo,
	}
}

func (c *CollectorService) Collect(url string) (string, error) {
	loadedCollector, err := c.CollectorRepo.GetByUrl(url)
	if err != nil {
		return "", apperrors.ServicesCollectorCollectGetByUrlError.AppendMessage(err)
	}

	if loadedCollector != nil {
		return loadedCollector.Data, nil
	}

	service, err := selenium.NewChromeDriverService(c.Collector.PathChromeDriver, c.Collector.PortChromeDriver)
	if err != nil {
		return "", apperrors.ServicesCollectorCollectNewChromeDriverServiceError.AppendMessage(err)
	}

	defer service.Stop() //nolint:errcheck

	caps := addCapabilities(c.ProxyIp)
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return "", apperrors.ServicesCollectorCollectNewRemoteError.AppendMessage(err)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		return "", apperrors.ServicesCollectorCollectMaximizeWindow.AppendMessage(err)
	}

	err = driver.Get(url)
	if err != nil {
		return "", apperrors.ServicesCollectorCollectDriverGet.AppendMessage(err)
	}

	html, err := driver.PageSource()
	if err != nil {
		return "", apperrors.ServicesCollectorCollectPageSource.AppendMessage(err)
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
	if proxyIp == "" {
		caps.AddChrome(chrome.Capabilities{Args: []string{
			"--headless",
		}})
	} else {
		caps.AddChrome(chrome.Capabilities{Args: []string{
			"--headless",
			"--proxy-server=" + proxyIp,
		}})
	}

	// madify request headers to avoid detection
	caps.AddChrome(chrome.Capabilities{Args: []string{
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
	}})

	return caps
}
