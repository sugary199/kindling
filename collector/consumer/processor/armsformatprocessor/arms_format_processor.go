package armsformatprocessor

import (
	"github.com/Kindling-project/kindling/collector/component"
	"github.com/Kindling-project/kindling/collector/consumer"
	"github.com/Kindling-project/kindling/collector/consumer/processor"
	"github.com/Kindling-project/kindling/collector/model"
)

const ProcessorName = "armsformatprocessor"

// ArmsFormatProcessor generates new model.GaugeGroup according to the documentation.
type ArmsFormatProcessor struct {
	cfg          *Config
	nextConsumer consumer.Consumer
	telemetry    *component.TelemetryTools
}

func (r *ArmsFormatProcessor) Consume(gaugeGroup *model.GaugeGroup) error {
	// ARMS logic

	return nil
}

func NewArmsFormatProcessor(cfg interface{}, telemetry *component.TelemetryTools, nextConsumer consumer.Consumer) processor.Processor {
	processorCfg := cfg.(*Config)
	if processorCfg.SamplingRate == nil {
		processorCfg.SamplingRate = &SampleConfig{
			NormalData: 0,
			SlowData:   100,
			ErrorData:  100,
		}
	}
	return &ArmsFormatProcessor{
		cfg:          processorCfg,
		nextConsumer: nextConsumer,
		telemetry:    telemetry,
	}
}
