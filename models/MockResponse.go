package models

// MockResponse represents a mocked response struct
type MockResponse struct {
	Status      int
	ContentType string
	Charset     string
	Headers     map[string]string
	Delay       string
}
