package helpers

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/joakim-ribier/gmocky/models"
	"github.com/joakim-ribier/gmocky/utils"
)

// MockURLWriter represents the writer struct to generate a mock URL
type MockURLWriter struct {
	contentGMockyResource string
	headerGMockyResource  string
	port                  string
}

// NewMockURLWriter creates a new instance of MockURLWriter struct
func NewMockURLWriter(port string) MockURLWriter {
	return MockURLWriter{
		contentGMockyResource: utils.ContentGMockyResource,
		headerGMockyResource:  utils.HeaderGMockyResource,
		port:                  port,
	}
}

// Print creates and prints the new mock URL (from resources files folder)
func (h MockURLWriter) Print() {
	mr := h.getMockResponse()
	cr := utils.GetFileContent(h.contentGMockyResource).AsString()

	// Create request config from the mocked response
	config := models.NewConfig(mr).AsJSON()

	params := url.Values{}
	params.Add(utils.MockURLContentParam, cr)
	params.Add(utils.MockURLHeaderParam, config)
	params.Add(utils.MockURLDelayParam, mr.Delay)

	h.printLn(params, config, mr)
}

func (h MockURLWriter) getMockResponse() models.MockResponse {
	data := models.MockResponse{}
	json.Unmarshal(utils.GetFileContent(h.headerGMockyResource), &data)

	return data
}

func (h MockURLWriter) printLn(params url.Values, config string, mr models.MockResponse) {
	fmt.Println("# Generate mocked URL with encoded params")
	fmt.Println(":" + h.port + "/?" + params.Encode())

	fmt.Println("")

	fmt.Println("# Generate mocked URL with only encoded 'content response' param")
	params.Del(utils.MockURLHeaderParam)
	params.Del(utils.MockURLDelayParam)
	fmt.Println(
		":" + h.port + "/?" +
			utils.MockURLHeaderParam + "=" + config + "&" +
			utils.MockURLDelayParam + "=" + mr.Delay + "&" +
			params.Encode())
}
