package component

import (
	"github.com/Kindling-project/kindling/collector/observability"
	"github.com/Kindling-project/kindling/collector/observability/logger"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
	"log"
)

const (
	ObservabilityConfig = "observability"
	LogKey              = "logger"
	MetricKey           = "opentelemetry"
)

type TelemetryManager struct {
	Telemetry *TelemetryTools
}

func NewTelemetryManager() *TelemetryManager {
	return &TelemetryManager{
		Telemetry: NewDefaultTelemetryTools(),
	}
}

func (t *TelemetryManager) ConstructConfig(viper *viper.Viper) {
	t.initLogger(viper)
	t.initProvider(viper)
}

func (t *TelemetryManager) initLogger(viper *viper.Viper) {
	var loggerConfig = logger.Config{}
	key := ObservabilityConfig + "." + LogKey
	err := viper.UnmarshalKey(key, &loggerConfig)
	if err != nil {
		log.Printf("Error happened when reading logger config, and default config will be used: %v", err)
	}
	t.Telemetry.Logger = logger.InitLogger(loggerConfig)
}

func (t *TelemetryManager) initProvider(viper *viper.Viper) {
	var config = &observability.DefaultConfig
	key := ObservabilityConfig + "." + MetricKey
	err := viper.UnmarshalKey(key, config)
	if err != nil {
		log.Printf("Error happened when reading observability config, and default config will be used: %v", err)
	}
	meterProvider, err := observability.InitTelemetry(t.Telemetry.Logger, config)
	if err != nil {
		log.Printf("Error happened when initializing meter provider: %v", err)
		return
	}
	t.Telemetry.MeterProvider = meterProvider
}

type TelemetryTools struct {
	Logger        *zap.Logger
	MeterProvider metric.MeterProvider
}

func NewDefaultTelemetryTools() *TelemetryTools {
	return &TelemetryTools{
		Logger:        logger.CreateDefaultLogger(),
		MeterProvider: metric.NewNoopMeterProvider(),
	}
}
