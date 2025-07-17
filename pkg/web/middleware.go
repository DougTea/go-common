package web

import (
	"errors"
	"log/slog"
	"strconv"
	"strings"
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
		var err error
		err = c.Errors.Last()
		v := &Error{}
		errMsg := new(strings.Builder)
		errMsg.WriteString(err.Error())
		for errors.As(err, &v) {
			err = v.Cause
			if err != nil {
				errMsg.WriteString("\n")
				errMsg.WriteString("Cause by:\n")
				errMsg.WriteString(err.Error())
				errMsg.WriteString("\n")
			}
		}
		slog.Error(errMsg.String())
		if !errors.As(c.Errors.Last(), &v) {
			v = NewErrorWithCause(CommonError, err)
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
