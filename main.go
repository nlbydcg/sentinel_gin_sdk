package sentinel_sdk

import (
	"errors"
	"fmt"
	"os"

	sentinel "github.com/alibaba/sentinel-golang/api"
	sentinelFlow "github.com/alibaba/sentinel-golang/core/flow"
	sentinelSystem "github.com/alibaba/sentinel-golang/core/system"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

// 初始化sentinel 第一个文件是sentinel 配置项，第二规则配置文件
func InitSentinel(filepath string, rulePath string) error {
	if err := sentinel.InitWithConfigFile(filepath); err != nil {
		return err
	}

	if err := InitSentinelRulePath(rulePath); err != nil {
		return err
	}
	return nil
}

// 初始化规则配置项
func InitSentinelRulePath(rulePath string) error {
	if rulePath == "" {
		return errors.New("rule path is empty")
	}
	_, err := os.Stat(rulePath)
	if err != nil {
		if os.IsExist(err) {
			return errors.New("rule path is not exist")
		} else {
			return err
		}
	}

	content, err := os.ReadFile(rulePath)
	if err != nil {
		return err
	}

	var ruleConfig RulesConfig
	err = yaml.Unmarshal(content, &ruleConfig)
	if err != nil {
		return err
	}

	SetSentinelRules(&ruleConfig)
	return nil
}

// SetSentinelRules 更新过滤规则
func SetSentinelRules(ruleConfig *RulesConfig) error {
	if ruleConfig == nil {
		return errors.New("rule config is nil")
	}

	_, err := sentinelFlow.LoadRules(ruleConfig.Rules.ToSentinelFlow())
	if err != nil {
		return err
	}

	_, err = sentinelSystem.LoadRules(ruleConfig.System)
	if err != nil {
		return err
	}
	return nil
}

// 初始化规则yaml文件 InitRuleYamlByGin 生成rule文件
func InitRuleYamlByGin(routes []*gin.RouteInfo, baseOptions FlowOptions, filepath string) error {
	rules := make([]*FlowRule, len(routes))
	for i, v := range routes {
		fr := &FlowRule{
			Resource:    fmt.Sprintf("%s:%s", v.Method, v.Path),
			FlowOptions: baseOptions,
		}
		rules[i] = fr
	}

	config := RulesConfig{
		Rules: rules,
	}
	if err := SetYamlByRulesConfig(config, filepath); err != nil {
		return err
	}

	return nil
}

// 根据配置文件生成对应的app
func SetYamlByRulesConfig(config RulesConfig, filePath string) error {
	if filePath == "" {
		return errors.New("rule path is empty")
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	output, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	_, err = file.Write(output)
	if err != nil {
		return err
	}

	return nil
}
