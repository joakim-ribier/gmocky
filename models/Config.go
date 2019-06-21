package models

import (
	"encoding/json"

	"github.com/joakim-ribier/gmocky/utils"
)

// Config represents a struct to write the response
type Config struct {
	Status      int               `json:"status"`
	ContentType string            `json:"contentType"`
	Charset     string            `json:"charset"`
	Headers     map[string]string `json:"headers"`
}

// GetHeaderContentType builds Content-Type header value
func (config Config) GetHeaderContentType() string {
	if config.ContentType == "" {
		return ""
	}
	charset := ""
	if config.Charset != "" {
		charset = "; charset=" + config.Charset
	}
	return config.ContentType + charset
}

// AsJSON gets json string representation of the config struct
func (config Config) AsJSON() string {
	data, error := json.Marshal(config)
	if error != nil {
		return "json.Marshal error..."
	}
	var fc utils.FileContent = data
	return fc.AsString()
}

// NewConfig creates new config struct from a MockResponse struct
func NewConfig(r MockResponse) Config {
	charset := "utf-8"
	if v := r.Charset; v != "" {
		charset = v
	}
	return Config{
		Status:      r.Status,
		ContentType: r.ContentType,
		Charset:     charset,
		Headers:     r.Headers,
	}
}
