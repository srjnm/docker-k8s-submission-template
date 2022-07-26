package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type apiController struct {
}

func NewApiController() ApiController {
	return &apiController{}
}

type ApiController interface {
	Root() string
	CCGet() string
	CCPost(*gin.Context) (int, *http.Response, error)
}

func (c *apiController) Root() string {
	return "CWiCS Assessment"
}

func (c *apiController) CCGet() string {
	return "POST to this endpoint with JSON to convert to YAML"
}

func (c *apiController) CCPost(cxt *gin.Context) (int, *http.Response, error) {
	var jsonData interface{}
	body, err := cxt.GetRawData()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if json.Unmarshal(body, &jsonData) != nil {
		return http.StatusBadRequest, nil, errors.New("invalid json!")
	}

	counterUrl := os.Getenv("COUNTER_URL")
	if counterUrl == "" {
		counterUrl = "http://127.0.0.1:8080/count"
	}

	req, err := http.NewRequest(http.MethodPost, counterUrl, nil)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if resp.StatusCode != 200 {
		return http.StatusInternalServerError, nil, errors.New("service-counter did not return 200")
	}

	converterUrl := os.Getenv("CONVERTER_URL")
	if converterUrl == "" {
		converterUrl = "http://127.0.0.1:3001/convert"
	}

	req, err = http.NewRequest(http.MethodPost, converterUrl, cxt.Request.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if resp.StatusCode != 200 {
		return http.StatusInternalServerError, nil, errors.New("service-connecter did not return 200")
	}

	return http.StatusOK, resp, nil
}
