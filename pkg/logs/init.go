package logs

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init() {
	env := os.Getenv("GIN_MODE")

	logConfig := zap.NewProductionConfig()

	// Set output paths for both file and console
	logConfig.OutputPaths = []string{"stdout", "./logs/app.log"}
	logConfig.ErrorOutputPaths = []string{"stderr", "./logs/error.log"}

	// Ensure JSON format
	logConfig.Encoding = "json"

	// Create directory for logs if it doesn't exist
	if err := os.MkdirAll("./logs", 0755); err != nil {
		log.Fatalf("error creating log directory: %s", err.Error())
	}

	// Configure log level based on environment
	if env == "debug" {
		logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// Build the logger
	logger, err := logConfig.Build()
	if err != nil {
		log.Fatalf("error creating logger: %s", err.Error())
	}

	Logger = logger
}

func GetTraceID(c *gin.Context) string {
	if traceID, exists := c.Get("traceID"); exists {
		return traceID.(string)
	}
	return ""
}
