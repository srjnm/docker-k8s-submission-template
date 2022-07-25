package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srjnm/docker-k8s-submission-template/server/controllers"
)

type ApiHandler struct {
	apiController controllers.ApiController
}

func NewApiHandler(apiController controllers.ApiController) *ApiHandler {
	return &ApiHandler{
		apiController: apiController,
	}
}

func (h *ApiHandler) RootHandler(cxt *gin.Context) {
	str := h.apiController.Root()

	cxt.Data(http.StatusOK, "text/plain", []byte(str))
	return
}

func (h *ApiHandler) CCGetHandler(cxt *gin.Context) {
	str := h.apiController.CCGet()

	cxt.Data(http.StatusOK, "text/plain", []byte(str))
	return
}
