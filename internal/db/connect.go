package db

import (
	"fmt"
	"github.com/hecomp/catchall/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/sharding"

	"github.com/hecomp/catchall/internal/config"
	"gorm.io/gorm"
)

func NewConnect(l *log.Logger, dbConfig *config.Configurations) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBName, dbConfig.DBPort)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		l.Fatal(err)
		return nil, err
	}

	middleware := sharding.Register(sharding.Config{
		ShardingKey:         "name",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, dbConfig.DBSchema)
	db.Use(middleware)

	err = db.AutoMigrate(&models.DomainName{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
