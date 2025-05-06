package web

import (
	"errors"
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
		v := &Error{}
		if !errors.As(last.Err, v) {
			v = NewErrorWithCause(CommonError, last.Err)
		}
		status, err := strconv.Atoi(v.Code.String()[:3])
		if err != nil || status < 100 || status > 599 {
			status = 500
		}
		c.JSON(status, &ErrorMessage{
			Timestamp: time.Now().Unix(),
			Status:    status,
			Message:   v.Error(),
			Code:      v.Code.String(),
		})
	}
}
