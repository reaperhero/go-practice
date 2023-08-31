package main

import (
	"fmt"
	"strings"
)

const dmCommaFiled = "\""

func Split(source string, split string) []string {
	data := strings.Split(source, split)
	if len(data) == 0 && data[0] == "" {
		return nil
	}
	return data
}

func StringContain(a []string, s string) bool {
	for _, str := range a {
		if str == s {
			return true
		}
	}
	return false
}

func initSql(query string) []string {
	// init
	query = strings.ReplaceAll(query, "\t", " ")
	query = strings.ReplaceAll(query, "\n", " ")
	query = strings.ReplaceAll(query, "(", " ( ")
	query = strings.ReplaceAll(query, ")", " ) ")

	querySplits := strings.FieldsFunc(query, func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	})
	return querySplits
}

// }
func getTableAlias(querySplits []string) []string {
	var (
		asStart    bool
		aliasTable = []string{}
	)
	for i := 0; i < len(querySplits); i++ {
		if strings.ToLower(querySplits[i]) == "from" || strings.ToLower(querySplits[i]) == "join" {
			asStart = true
		}
		if asStart && strings.ToLower(querySplits[i]) == "as" {
			aliasTable = append(aliasTable, querySplits[i+1])
			i++
			asStart = false
			continue
		}
	}
	return aliasTable
}

func main() {

	query := "SELECT DISTINCT deploy_host.group FROM deploy_host\nLEFT JOIN deploy_cluster_host_rel ON deploy_host.sid = deploy_cluster_host_rel.sid\nLEFT JOIN deploy_instance_list ON deploy_cluster_host_rel.sid = deploy_instance_list.sid\nLEFT JOIN deploy_product_list ON deploy_instance_list.pid = deploy_product_list.id\nLEFT JOIN sidecar_list ON sidecar_list.id = deploy_host.sid\nWHERE deploy_host.isDeleted=0 AND deploy_cluster_host_rel.is_deleted=0 AND deploy_host.sid != '' AND (deploy_host.hostname LIKE '%dsada%' OR deploy_host.ip LIKE '%asdsa%')"

	// 找到表名  及  别名
	// 字符串分割后，替换旧的sql

	// 分割
	querySplits := initSql(query)
	fmt.Println(strings.Join(querySplits, "\n"))
	// 获取表别名
	tableAlias := getTableAlias(querySplits)

	// 别名字段处理双引号
	mark := replacaAliasAddMark(querySplits, tableAlias)

	fmt.Println(strings.Join(mark, " "))
}

func replacaAliasAddMark(querySplits, tableAlias []string) []string {
	oldquerySplits := querySplits
	for i := 0; i < len(querySplits); i++ {
		// 遇到as 跳到下下个
		if strings.ToLower(querySplits[i]) == "as" || strings.ToLower(querySplits[i]) == "like" {
			i++
			continue
		}

		// 系统关键字
		if _, ok := sss[strings.ToLower(querySplits[i])]; ok {
			continue
		}

		// 处理条件语句
		if isonditions(querySplits[i]) {
			conditions := handleConditions(querySplits[i], tableAlias)
			oldquerySplits[i] = conditions
		} else {
			// 处理非条件语句
			if strings.Contains(querySplits[i], ",") && strings.Contains(querySplits[i], ".") { // u.user_name, user_list.user_name,
				field := strings.Split(querySplits[i], ".")
				if StringContain(tableAlias, field[0]) {
					inFields := strings.Split(field[1], ",")
					oldquerySplits[i] = field[0] + "." + dmCommaFiled + inFields[0] + dmCommaFiled + ","
				} else {
					inFields := strings.Split(field[1], ",")
					oldquerySplits[i] = dmCommaFiled + field[0] + dmCommaFiled + "." + dmCommaFiled + inFields[0] + dmCommaFiled + ","
				}
				continue
			}
			if !strings.Contains(querySplits[i], ",") && strings.Contains(querySplits[i], ".") { // user_list.user_name u.user_name
				field := strings.Split(querySplits[i], ".")
				if StringContain(tableAlias, field[0]) {
					oldquerySplits[i] = field[0] + "." + dmCommaFiled + field[1] + dmCommaFiled
				} else {
					oldquerySplits[i] = dmCommaFiled + field[0] + dmCommaFiled + "." + dmCommaFiled + field[1] + dmCommaFiled
				}
				continue
			}
			if strings.Contains(querySplits[i], ",") && !strings.Contains(querySplits[i], ".") { // user_name,
				inFields := strings.Split(querySplits[i], ",")
				oldquerySplits[i] = dmCommaFiled + inFields[0] + dmCommaFiled + ","
				continue
			}
			if !strings.Contains(querySplits[i], ",") && !strings.Contains(querySplits[i], ".") { // user_name
				oldquerySplits[i] = dmCommaFiled + querySplits[i] + dmCommaFiled
				continue
			}
		}

	}
	return oldquerySplits
}

// 表别名  关键字不用替换
var sss = map[string]*struct{}{
	"select":   nil,
	"distinct": nil,
	"as":       nil,
	"where":    nil,
	"and":      nil,
	"on":       nil,
	"or":       nil,
	"=":        nil,
	"!=":       nil,
	"''":       nil,
	">":        nil,
	"(":        nil,
	")":        nil,
	"<":        nil,
	">=":       nil,
	"<=":       nil,
	"from":     nil,
	"left":     nil,
	"right":    nil,
	"inner":    nil,
	"join":     nil,
}

func isonditions(query string) bool {
	split := strings.FieldsFunc(query, func(r rune) bool {
		if r == '!' || r == '=' || r == '>' || r == '<' {
			return true
		}
		return false
	})
	return len(split) >= 2
}
func handleConditions(query string, tableAlias []string) string {
	split := strings.FieldsFunc(query, func(r rune) bool {
		if r == '!' || r == '=' || r == '>' || r == '<' {
			return true
		}
		return false
	})

	operator := strings.Replace(query, split[0], "", -1)
	operator = strings.Replace(operator, split[len(split)-1], "", -1)

	querySplits := []string{split[0], split[len(split)-1]}

	for i := 0; i < len(querySplits); i++ {
		if i == 0 {
			if !strings.Contains(querySplits[i], ",") && strings.Contains(querySplits[i], ".") { // user_list.user_name u.user_name
				field := strings.Split(querySplits[i], ".")
				if StringContain(tableAlias, field[0]) {
					querySplits[i] = field[0] + "." + dmCommaFiled + field[1] + dmCommaFiled
				} else {
					querySplits[i] = dmCommaFiled + field[0] + dmCommaFiled + "." + dmCommaFiled + field[1] + dmCommaFiled
				}
				continue
			}
			if !strings.Contains(querySplits[i], ",") && !strings.Contains(querySplits[i], ".") { // user_name
				querySplits[i] = dmCommaFiled + querySplits[i] + dmCommaFiled
				continue
			}
		}

		if i == 1 {
			// 处理非条件语句
			if !strings.Contains(querySplits[i], ",") && strings.Contains(querySplits[i], ".") { // user_list.user_name u.user_name
				field := strings.Split(querySplits[i], ".")
				if StringContain(tableAlias, field[0]) {
					querySplits[i] = field[0] + "." + dmCommaFiled + field[1] + dmCommaFiled
				} else {
					querySplits[i] = dmCommaFiled + field[0] + dmCommaFiled + "." + dmCommaFiled + field[1] + dmCommaFiled
				}
				continue
			}
			if !strings.Contains(querySplits[i], ",") && !strings.Contains(querySplits[i], ".") { // user_name
				querySplits[i] = "'" + querySplits[i] + "'"
				continue
			}
		}
	}
	return strings.Join([]string{querySplits[0], operator, querySplits[1]}, " ")
}
