package handlers

import (
	"io"
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

func (h *ApiHandler) CCPostHandler(cxt *gin.Context) {
	status, resp, err := h.apiController.CCPost(cxt)
	if err != nil {
		cxt.AbortWithError(status, err)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		cxt.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	cxt.Writer.Write(respBody)
	return
}
