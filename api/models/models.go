package models

import (
	"time"

	"gorm.io/gorm"
)

type ResponseURL struct {
	gorm.Model

	URL         string
	CustomShort string
	Expiry      time.Duration
}
