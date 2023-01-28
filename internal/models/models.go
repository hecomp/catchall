package models

// DomainName is the models type for user object
type DomainName struct {
	ID             int64  `json:"id" gorm:"primarykey"`
	Name           string `json:"name"`
	DeliveredEvent int64  `json:"deliveredEvent"`
	BouncedEvent   int64  `json:"bouncedEvent"`
	Status         string `json:"status"`
}

type DbConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
}

// CatchAllResponse collects the response values for the catchall method.
type CatchAllResponse struct {
	Message string      `json:",omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Err     string      `json:"err,omitempty"`
}
