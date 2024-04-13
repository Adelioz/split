package logging

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger(devel bool) *zap.Logger {
	var cfg zap.Config
	if devel {
		cfg = zap.NewDevelopmentConfig()
		cfg.DisableStacktrace = true
	} else {
		cfg = zap.NewProductionConfig()
	}
	l, err := cfg.Build(zap.AddCaller())
	if err != nil {
		log.Fatalf("Failed to init logging: %s.", err)
	}
	return l
}
