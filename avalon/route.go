package avalon

import (
	"Avalon/avalon/handle"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	handle.Router(r)
}
