package providers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IApi interface {
	Home(c *gin.Context)
}

type Api struct {
	Config   IConfig
	Database IDatabase
}

func NewApi(config IConfig, database IDatabase) IApi {
	return &Api{
		Config:   config,
		Database: database,
	}
}

// Home is an HTTP handler that returns the current version of the API.
// It responds with a JSON object containing the version string.
func (api *Api) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "v0.0.1",
	})
}
