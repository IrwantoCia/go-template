package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwantocia/go-template/pkg/logs"
	"github.com/irwantocia/go-template/pkg/validation"
	"go.uber.org/zap"
)

func Index(r *gin.Engine) {
	type getIndex struct {
		Age int `form:"age" validate:"omitempty,gte=0,lte=120"`
	}
	r.GET("/", validation.Validate[getIndex](), func(c *gin.Context) {
		param := c.MustGet("payload").(getIndex)
		traceID := logs.GetTraceID(c)

		logs.Logger.Info("Home route accessed", zap.String("trace_id", traceID), zap.Int("age", param.Age))
		c.HTML(http.StatusOK, "index", gin.H{})
	})

	type postIndex struct {
		Name  string `json:"name" validate:"required,min=2,max=50"`
		Email string `json:"email" validate:"required,email"`
		Age   int    `json:"age" validate:"required,gte=19,lte=100"`
	}
	r.POST("/", validation.Validate[postIndex](), func(c *gin.Context) {
		user := c.MustGet("payload").(postIndex)

		traceID := logs.GetTraceID(c)
		logs.Logger.Sugar().Infow("Home route accessed", zap.String("trace_id", traceID), zap.Any("user", user))
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
}
