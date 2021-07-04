package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	router := GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/convert", strings.NewReader(`
		{
			"value": 0
		}
	`))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var result struct {
		Value float64 `json:"value"`
	}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, float64(32), result.Value)
}

func TestConvertNegative(t *testing.T) {
	router := GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/convert", strings.NewReader(`
		{
			"value": -10
		}
	`))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var result struct {
		Value float64 `json:"value"`
	}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, float64(14), result.Value)
}

func TestConvertFail(t *testing.T) {
	router := GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/convert", strings.NewReader(`
		{
			"value": -10
		}
	`))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var result struct {
		Value float64 `json:"value"`
	}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, float64(15), result.Value)
}
