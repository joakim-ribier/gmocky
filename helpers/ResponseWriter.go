package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/joakim-ribier/gmocky/models"
	"github.com/joakim-ribier/gmocky/utils"
)

// ResponseWriter represents the writer struct to write the response
type ResponseWriter struct {
	writer  http.ResponseWriter
	request *http.Request
}

// NewResponseWriter creates a new instance of ResponseWriter struct
func NewResponseWriter(w http.ResponseWriter, r *http.Request) ResponseWriter {
	return ResponseWriter{
		writer:  w,
		request: r,
	}
}

func (h ResponseWriter) writeResponse(config models.Config) {
	for key, value := range config.Headers {
		h.writer.Header().Set(key, value)
	}
	if contentType := config.GetHeaderContentType(); contentType != "" {
		h.writer.Header().Set("Content-Type", contentType)
	}

	if config.Status != 0 {
		h.writer.WriteHeader(config.Status)
	}

	if body := h.getBody(h.request); body != "" {
		h.writer.Write([]byte(body))
	}
}

func (h ResponseWriter) getBody(r *http.Request) string {
	values, ok := r.URL.Query()[utils.MockURLContentParam]
	if !ok || len(values[0]) < 1 {
		return ""
	}
	return values[0]
}

// Write writes the response depends on the mock URL configuration
func (h ResponseWriter) Write() {
	keys, ok := h.request.URL.Query()[utils.MockURLHeaderParam]
	if !ok || len(keys[0]) < 1 {
		h.writeBadRequest("no '" + utils.MockURLHeaderParam + "' (header configuration) query parameter")
		return
	}

	data := models.Config{}
	if error := json.Unmarshal([]byte(keys[0]), &data); error != nil {
		h.writeBadRequest("bad json format")
		return
	}

	fmt.Printf("%s - %s\n\r",
		time.Now().Format(time.RFC3339),
		h.request.URL)

	h.writeResponse(data)
	h.timeout()
}

func (h ResponseWriter) timeout() {
	keys, ok := h.request.URL.Query()[utils.MockURLDelayParam]
	if !ok || len(keys[0]) < 1 {
		return
	}
	if duration, error := time.ParseDuration(keys[0]); error == nil {
		time.Sleep(duration)
	}
}

func (h ResponseWriter) writeBadRequest(message string) {
	h.writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	h.writer.Header().Set("X-Content-Type-Options", "nosniff")

	h.writer.WriteHeader(400)
	h.writer.Write([]byte(`{"error":"` + message + `"}`))
}
