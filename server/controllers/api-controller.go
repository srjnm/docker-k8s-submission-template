package controllers

type apiController struct {
}

func NewApiController() ApiController {
	return &apiController{}
}

type ApiController interface {
	Root() string
	CCGet() string
}

func (c *apiController) Root() string {
	return "CWiCS Assessment"
}

func (c *apiController) CCGet() string {
	return "POST to this endpoint with JSON to convert to YAML"
}
