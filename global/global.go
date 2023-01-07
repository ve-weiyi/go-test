package global

import "gorm.io/gorm"

type config struct {
	DB *gorm.DB
}

var CONFIG = new(config)
