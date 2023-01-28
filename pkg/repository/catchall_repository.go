package repository

import (
	"fmt"
	"github.com/hecomp/catchall/internal/constants"
	"gorm.io/gorm"
	"sync"

	"github.com/hecomp/catchall/internal/models"
)

type CatchAllRepository interface {
	SaveDeliveredEvent(domainName string) error
	SaveBouncedEvent(domainName string) error
	FindDomains(domainName string) (*models.DomainName, error)
}

type catchAllRepository struct {
	db   *gorm.DB
	lock *sync.RWMutex
}

func NewCatchAllRepository(db *gorm.DB) CatchAllRepository {
	var repositoory CatchAllRepository
	repositoory = &catchAllRepository{
		db:   db,
		lock: &sync.RWMutex{},
	}
	return repositoory
}

func (c catchAllRepository) SaveDeliveredEvent(domainName string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	var domainN models.DomainName
	err := c.db.Model(&models.DomainName{}).Where("name", domainName).Find(&domainN).Error
	if err != nil {
		return err
	}
	if domainN.Name == "" {
		domainN.Name = domainName
	}
	domainN.DeliveredEvent += 1
	if result := c.db.Save(&domainN); result.Error != nil {
		return result.Error
	}
	return nil
}

func (c catchAllRepository) SaveBouncedEvent(domainName string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	var domainN models.DomainName
	err := c.db.Model(&models.DomainName{}).Where("name", domainName).Find(&domainN).Error
	if err != nil {
		return err
	}
	if domainN.Name == "" {
		domainN.Name = domainName
	}
	domainN.BouncedEvent += 1
	if result := c.db.Save(&domainN); result.Error != nil {
		return result.Error
	}
	return nil
}

func (c catchAllRepository) FindDomains(domainName string) (*models.DomainName, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	var domainN models.DomainName
	err := c.db.Model(&models.DomainName{}).Where("name", domainName).Find(&domainN).Error
	if err != nil {
		return nil, err
	}
	if domainN.ID == 0 {
		return nil, fmt.Errorf("not found")
	}
	switch event := domainN.DeliveredEvent; {
	case event > constants.EventsThreshHold:
		domainN.Status = constants.CatchAll.String()
	case event < constants.EventsThreshHold:
		var unknown constants.DomainStatus = 99
		domainN.Status = unknown.String()
	default:
		domainN.Status = constants.NotCatchAll.String()
	}
	return &domainN, nil
}
