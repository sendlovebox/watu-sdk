package model

type (
	// ErrorPayload schema
	ErrorPayload struct {
		HasError   bool        `mapstructure:"has_error" json:"has_error"`
		StatusCode interface{} `mapstructure:"status_code" json:"status_code"`
		Message    string      `mapstructure:"message" json:"message"`
		Data       struct {
			Status int `mapstructure:"status" json:"status"`
		} `mapstructure:"data" json:"data"`
	}

	// Request schema
	Request struct {
		IsFavourite    int     `json:"is_favourite"`
		ShouldPaginate int     `json:"should_paginate"`
		Country        Country `json:"country"`
	}

	// Country string representation of country
	Country string
)

const (
	// LogStrRequest log string key
	LogStrRequest = "request"
	// LogStrKeyEndpoint log endpoint name value
	LogStrKeyEndpoint = "endpoint"

	// BaseURL string
	BaseURL = "https://api.watupay.com/v1"

	// CountryNigeria string representation of Nigeria
	CountryNigeria Country = "NG"

	// RandomString random string
	RandomString = "0000000000"
)
