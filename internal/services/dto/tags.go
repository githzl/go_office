package dto

import "go_office/internal/pgmodel"

type Tags struct {
	pgmodel.Tags
	LogoUrl string `json:"logo_url"`
}
