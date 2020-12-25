package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PushoverController struct{}

func (pc PushoverController) Check(c *gin.Context) {
	c.JSON(http.StatusOK, "checkPushover")
	return
}
