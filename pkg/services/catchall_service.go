package services

import (
	"github.com/hecomp/catchall/internal/models"
	"github.com/hecomp/catchall/pkg/repository"
)

type CatchAllService interface {
	PutDeliveredEvent(domainName string) error
	PutBouncedEvent(domainName string) error
	GetDomains(domainName string) (*models.DomainName, error)
}

type catchAll struct {
	repo repository.CatchAllRepository
}

func NewCatchAllService(repo repository.CatchAllRepository) CatchAllService {
	return &catchAll{repo: repo}
}

func (c catchAll) PutDeliveredEvent(domainName string) error {

	if err := c.repo.SaveDeliveredEvent(domainName); err != nil {
		return err
	}
	return nil
}

func (c catchAll) PutBouncedEvent(domainName string) error {
	if err := c.repo.SaveBouncedEvent(domainName); err != nil {
		return err
	}
	return nil
}

func (c catchAll) GetDomains(domainName string) (*models.DomainName, error) {
	resp, err := c.repo.FindDomains(domainName)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
