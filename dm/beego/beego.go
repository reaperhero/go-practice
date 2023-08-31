package main

import (
	"fmt"
	_ "gitee.com/chunanyong/dm"
	"github.com/reaperhero/go-practice/dm/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(AgentList))
	orm.RegisterDriver("dm", orm.DRDM)
	// ?logLevel=all

	orm.RegisterDataBase("dmOrm", "dm", "dm://SYSDBA:SYSDBA001@172.16.104.165:5236")
	orm.Debug = true
}

//type AgentList struct {
//	Id             string    `db:"id" orm:"pk;column(id)"`                          //Agent ID (UUID)
//	SidecarId      string    `db:"sidecar_id" orm:"column(sidecar_id)"`             //sidecar id
//	Type           int8      `db:"type" orm:"column(type)"`                         //agent类型
//	Name           string    `db:"name" orm:"column(name)"`                         //agent 名称
//	Version        string    `db:"version" orm:"column(version)"`                   //agent版本
//	IsUninstalled  int8      `db:"is_uninstalled" orm:"column(is_uninstalled)"`     //是否已被卸载
//	DeployDate     time.Time `db:"deploy_date" orm:"column(deploy_date)"`           //agent部署时间
//	AutoDeployment int8      `db:"auto_deployment" orm:"column(auto_deployment)"`   //是否是自动部署的
//	LastUpdateDate time.Time `db:"last_update_date" orm:"column(last_update_date)"` //最近更新时间
//	AutoUpdated    int8      `db:"auto_updated" orm:"column(auto_updated)"`         //是否是自动升级的
//}

type AgentList struct {
	Id             string    `db:"id" orm:"pk;column(id)"` //Agent ID (UUID)
	SidecarId      string    `db:"sidecar_id"`             //sidecar id
	Type           int8      `db:"type"`                   //agent类型
	Name           string    `db:"name"`                   //agent 名称
	Version        string    `db:"version"`                //agent版本
	IsUninstalled  int8      `db:"is_uninstalled"`         //是否已被卸载
	DeployDate     time.Time `db:"deploy_date"`            //agent部署时间
	AutoDeployment int8      `db:"auto_deployment"`        //是否是自动部署的
	LastUpdateDate time.Time `db:"last_update_date"`       //最近更新时间
	AutoUpdated    int8      `db:"auto_updated"`           //是否是自动升级的
}

//func (m *AgentList) TableName() string {
//	return "agent_list"
//}

var o orm.Ormer

func main() {

	o = orm.NewOrmUsingDB("dmOrm")

	err := query()
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Minute * 5)
	builder()
	time.Sleep(time.Second * 10)
	delete()
	err := create()
	if err != nil {
		panic(err)
	}
	err = query()
	if err != nil {
		panic(err)
	}
	err = update()
	if err != nil {
		panic(err)
	}
	err = query()
	if err != nil {
		panic(err)
	}
	err = delete()
	if err != nil {
		panic(err)
	}
}

func create() error {
	a := AgentList{
		Id:             "06c2f181-4195-438c-a661-46d03dc98309",
		SidecarId:      "1e04d624-b831-4ae2-a028-23f06449e0c1",
		Type:           1,
		Name:           "hdfs_datanode",
		Version:        "v1",
		IsUninstalled:  1,
		DeployDate:     time.Now(),
		AutoDeployment: 0,
		LastUpdateDate: time.Now(),
		AutoUpdated:    0,
	}
	insert, err := o.Insert(&a)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(insert)
	return nil
}

func query() error {
	a := AgentList{}
	a.Id = "06c2f181-4195-438c-a661-46d03dc98307"
	err := o.Read(&a)
	if err != nil {
		return err
	}
	fmt.Println(a)

	err = o.QueryTable(&a).Filter("id", "06c2f181-4195-438c-a661-46d03dc9830as").One(&a)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func update() error {
	a := AgentList{Version: "v3"}
	_, err := o.QueryTable(&a).Filter("id", "06c2f181-4195-438c-a661-46d03dc98307").Update(orm.Params{
		"version": "v3",
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func delete() error {
	a := AgentList{}
	_, err := o.QueryTable(&a).Filter("id", "06c2f181-4195-438c-a661-46d03dc98307").Delete()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func builder() {
	//qb, _ := orm.NewQueryBuilder("mysql")
	//
	//// 构建查询对象
	//qb.Select("user.name",
	//    "profile.age").
	//    From("user").
	//    InnerJoin("profile").On("user.id_user = profile.fk_user").
	//    Where("age > ?").
	//    OrderBy("name").Desc().
	//    Limit(10).Offset(0)
	//
	//// 导出 SQL 语句
	//sql := qb.String()
	//
	//// 执行 SQL 语句
	//o := orm.NewOrm()
	//o.Raw(sql, 20).QueryRows(&users)
	//var age []AgentList
	//qb, _ := orm.NewQueryBuilder("dm")
	//qb.Select("deploy_cluster_host_rel.id", "name").From("deploy_cluster_host_rel").InnerJoin("deploy_host").On(`"deploy_cluster_host_rel"."sid" = "deploy_host"."sid"`).
	//	Where(`"deploy_host"."is_deleted" = 1`).And(`"deploy_host"."is_deleted" = 0`).OrderBy("id").Desc().Limit(10).Offset(0)
	qb, _ := orm.NewQueryBuilder("dm")
	qb.Select("t1.id", "t2.name").From("deploy_host01 t2").LeftJoin("deploy_host02 t1").On(`t1."instance_id" = t2."id"`).And(`"t1.instance_id = t2.id"`).
		Where(`"t1.id" = "12345678"`).And(`"t2"."is_deleted" = 0`).OrderBy("id").Desc().Limit(10).Offset(20)
	//o.Raw(qb.String()).QueryRows()
	//rows, err := o.Raw(qb.String()).QueryRows(&age)
	//if err != nil {
	//	return
	//}

	fmt.Println(qb.String())
}
