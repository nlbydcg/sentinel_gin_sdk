package sentinel_gin_sdk

import (
	sentinelPlugin "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
)

// 抽离展示出来
var SentinelMiddleware = sentinelPlugin.SentinelMiddleware
var WithResourceExtractor = sentinelPlugin.WithResourceExtractor
var WithBlockFallback = sentinelPlugin.WithBlockFallback
