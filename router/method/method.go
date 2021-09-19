package method

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetResponse struct {
	Args    gin.Params  `json:"args,omitempty"`
	Headers http.Header `json:"headers"`
	Origin  string      `json:"origin"`
	Url     string      `json:"url"`
}

func Get(c *gin.Context) {
	getResponse := &GetResponse{
		Args:    c.Params,
		Headers: c.Request.Header,
		Origin:  c.Request.RemoteAddr,
		Url:     c.Request.RequestURI,
	}

	c.JSON(http.StatusOK, getResponse)
}

func Post(c *gin.Context) {

}

func Put(c *gin.Context) {

}

func Patch(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
