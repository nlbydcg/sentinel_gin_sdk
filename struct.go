package sentinel_sdk

import (
	sentinelFlow "github.com/alibaba/sentinel-golang/core/flow"
	sentinelSystem "github.com/alibaba/sentinel-golang/core/system"
)

type RulesConfig struct {
	System []*sentinelSystem.Rule `json:"system" yaml:"system"` // 系统级限流
	Rules  FlowRules              `json:"rules" yaml:"rules"`   // 某一接口限流策略
}

type FlowRule struct {
	Resource string `yaml:"resource" json:"resource" binding:"required" form:"resource"`
	FlowOptions
}

type FlowRules []*FlowRule

func (list FlowRules) ToSentinelFlow() []*sentinelFlow.Rule {
	result := make([]*sentinelFlow.Rule, len(list))
	for i, v := range list {
		result[i] = &sentinelFlow.Rule{
			Resource:               v.Resource,
			TokenCalculateStrategy: v.TokenCalculateStrategy,
			ControlBehavior:        v.ControlBehavior,
			Threshold:              v.Threshold,
			RelationStrategy:       v.RelationStrategy,
			RefResource:            v.RefResource,
			MaxQueueingTimeMs:      v.MaxQueueingTimeMs,
			WarmUpPeriodSec:        v.WarmUpPeriodSec,
			WarmUpColdFactor:       v.WarmUpColdFactor,
			StatIntervalInMs:       v.StatIntervalInMs,
			LowMemUsageThreshold:   v.LowMemUsageThreshold,
			HighMemUsageThreshold:  v.HighMemUsageThreshold,
			MemLowWaterMarkBytes:   v.MemLowWaterMarkBytes,
			MemHighWaterMarkBytes:  v.MemHighWaterMarkBytes,
		}
	}
	return result
}

type FlowOptions struct {
	TokenCalculateStrategy sentinelFlow.TokenCalculateStrategy `yaml:"tokenCalculateStrategy" json:"tokenCalculateStrategy" form:"tokenCalculateStrategy"`
	ControlBehavior        sentinelFlow.ControlBehavior        `yaml:"controlBehavior" json:"controlBehavior" form:"controlBehavior"`
	Threshold              float64                             `yaml:"threshold" json:"threshold" form:"threshold"`
	RelationStrategy       sentinelFlow.RelationStrategy       `yaml:"relationStrategy" json:"relationStrategy" form:"relationStrategy"`
	RefResource            string                              `yaml:"refResource" json:"refResource" form:"refResource"`
	MaxQueueingTimeMs      uint32                              `yaml:"maxQueueingTimeMs" json:"maxQueueingTimeMs" form:"maxQueueingTimeMs"`
	WarmUpPeriodSec        uint32                              `yaml:"warmUpPeriodSec" json:"warmUpPeriodSec" form:"warmUpPeriodSec"`
	WarmUpColdFactor       uint32                              `yaml:"warmUpColdFactor" json:"warmUpColdFactor" form:"warmUpColdFactor"`
	StatIntervalInMs       uint32                              `yaml:"statIntervalInMs" json:"statIntervalInMs" form:"statIntervalInMs"`
	LowMemUsageThreshold   int64                               `yaml:"lowMemUsageThreshold" json:"lowMemUsageThreshold" form:"lowMemUsageThreshold"`
	HighMemUsageThreshold  int64                               `yaml:"highMemUsageThreshold" json:"highMemUsageThreshold" form:"highMemUsageThreshold"`
	MemLowWaterMarkBytes   int64                               `yaml:"memLowWaterMarkBytes" json:"memLowWaterMarkBytes" form:"memLowWaterMarkBytes"`
	MemHighWaterMarkBytes  int64                               `yaml:"memHighWaterMarkBytes" json:"memHighWaterMarkBytes" form:"memHighWaterMarkBytes"`
}
