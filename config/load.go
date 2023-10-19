package config

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

const (
	defaultConfigSuffix = "yaml"
)

var (
	configPathList = []string{
		"./",
		"./config",
		"../",
		"../config",
	}
)

func LoadDefaultConf[T any](conf *T, confFile, mergeFile string) error {
	v := viper.New()
	for _, path := range configPathList {
		v.AddConfigPath(path)
	}
	v.SetConfigName(confFile)
	v.SetConfigType(defaultConfigSuffix)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.SetConfigName(mergeFile)
	if err := v.MergeInConfig(); err != nil {
		klog.Error(err)
	}
	return v.Unmarshal(conf, viper.DecodeHook(StringRenderTextTemplateHookFunc()))
}
