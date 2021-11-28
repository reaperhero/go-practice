package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"testing"
)

func TestApolloClinet(t *testing.T) {
	c := &config.AppConfig{
		AppID:          "appid",
		Cluster:        "PRO",
		IP:             "http://192.168.50.24:8080",
		NamespaceName:  "namespace1,namespace1",
		IsBackupConfig: true,
		Secret:         "",
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	cache := client.GetConfigCache("namespace1")
	cache.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return true
	})
}