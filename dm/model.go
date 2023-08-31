package main

import "time"

type AgentList struct {
	Id             string    `db:"id"`               //Agent ID (UUID)
	SidecarId      string    `db:"sidecar_id"`       //sidecar id
	Type           int8      `db:"type"`             //agent类型
	Name           string    `db:"name"`             //agent 名称
	Version        string    `db:"version"`          //agent版本
	IsUninstalled  int8      `db:"is_uninstalled"`   //是否已被卸载
	DeployDate     time.Time `db:"deploy_date"`      //agent部署时间
	AutoDeployment int8      `db:"auto_deployment"`  //是否是自动部署的
	LastUpdateDate time.Time `db:"last_update_date"` //最近更新时间
	AutoUpdated    int8      `db:"auto_updated"`     //是否是自动升级的
}

func (a AgentList) GetTableName() string {
	return "agent_list"
}

func (a AgentList) GetPKColumnName() string {
	return "id"
}

func (a AgentList) GetEntityMapPkSequence() string {
	return ""
}

func (a AgentList) GetDBFieldMap() map[string]interface{} {
	return map[string]interface{}{
		"id": "sadas",
	}
}

func (a AgentList) GetDBFieldMapKey() []string {
	return []string{"id"}
}

func (a AgentList) Set(key string, value interface{}) map[string]interface{} {
	return nil
}
