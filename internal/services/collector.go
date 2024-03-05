package services

import (
	"content_collector/internal/apperrors"
	"content_collector/internal/domain/model"
	"content_collector/internal/repository"
	"content_collector/internal/utils/scrappers"
	"content_collector/internal/utils/smartproxy"
)

type ICollectorService interface {
	Collect(url string) (*scrappers.ScrapperData, error)
	SetProxy(proxyIp string)
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

func (c *CollectorService) Collect(url string) (*scrappers.ScrapperData, error) {
	loadedCollector, err := c.CollectorRepo.GetByUrl(url)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.MongoCollectorRepositoryGetByIdErrNoDocuments.Code {
			return nil, apperrors.ServicesCollectorCollectGetByUrlError.AppendMessage(err)
		}
	}

	if loadedCollector != nil {
		return &scrappers.ScrapperData{
			Url:    loadedCollector.Url,
			Length: len(loadedCollector.Data),
			Data:   loadedCollector.Data,
			Code:   loadedCollector.DataCode,
			Status: loadedCollector.DataStatusText,
		}, nil
	}

	proxyIpSmart, err := c.SmartProxy.GetProxyRandomSmartProxy()
	if err != nil {
		return nil, apperrors.ServicesCollectorCollectGetProxyRandomError.AppendMessage(err)
	}

	c.Scrapper.SetProxy(proxyIpSmart.String())
	c.Scrapper.SetSmartProxy(proxyIpSmart)

	scrapStruct, err := c.Scrapper.Scrap(url)
	if err != nil {
		return nil, apperrors.ServicesCollectorCollectScrapError.AppendMessage(err)
	}

	collectorCreate := &model.CollectorRepository{
		Url:            url,
		Data:           scrapStruct.Data,
		DataCode:       scrapStruct.Code,
		DataStatusText: scrapStruct.Status,
		Status:         model.CollectorRepositoryStatusActive,
	}
	err = c.CollectorRepo.Create(collectorCreate)
	if err != nil {
		return nil, apperrors.ServicesCollectorCollectCreateError.AppendMessage(err)
	}

	return scrapStruct, nil
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
