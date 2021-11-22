package web

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Timestamp int64  `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Code      string `json:"code"`
}

func ErrorFormatMiddleWare(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		last := c.Errors.Last()
		v, ok := last.Err.(*Error)
		if !ok {
			v = NewErrorWithCause(CommonError, last.Err)
		}
		status, _ := strconv.Atoi(v.Code.String()[:3])
		c.JSON(status, &ErrorMessage{
			Timestamp: time.Now().Unix(),
			Status:    status,
			Message:   v.Error(),
			Code:      v.Code.String(),
		})
	}
}
