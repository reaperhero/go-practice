// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orm

import (
	"fmt"
	"strconv"
	"strings"
)

// dmCommaSpace is the separation
const dmCommaSpace = ", "
const dmCommaFiled = "\""

// DmQueryBuilder is the SQL build
type DmQueryBuilder struct {
	tokens []string
}

// Select will join the fields
// Select("*") Select("id","name")  Select(`deploy_cluster_host_rel.id`, "name")
func (qb *DmQueryBuilder) Select(fields ...string) QueryBuilder {
	if len(fields) == 1 && fields[0] == "*" {
		qb.tokens = append(qb.tokens, "SELECT", strings.Join(fields, dmCommaSpace))
		return qb
	}
	for i, field := range fields {
		if len(strings.Split(field, ".")) == 2 {
			fields[i] = dmCommaFiled + strings.Split(field, ".")[0] + dmCommaFiled + "." + dmCommaFiled + strings.Split(field, ".")[1] + dmCommaFiled
			continue
		}
		fields[i] = dmCommaFiled + field + dmCommaFiled
	}
	qb.tokens = append(qb.tokens, "SELECT", strings.Join(fields, dmCommaSpace))
	return qb
}

// ForUpdate add the FOR UPDATE clause
func (qb *DmQueryBuilder) ForUpdate() QueryBuilder {
	qb.tokens = append(qb.tokens, "FOR UPDATE")
	return qb
}

// From join the tables
// From("deploy_cluster_host_rel") From("deploy_cluster_host_rel t1")
func (qb *DmQueryBuilder) From(tables ...string) QueryBuilder {
	if len(strings.Split(tables[0], " ")) > 1 {
		qb.tokens = append(qb.tokens, "FROM", dmCommaFiled+strings.Split(tables[0], " ")[0]+dmCommaFiled+" as "+strings.Split(tables[0], " ")[1])
		return qb
	}
	qb.tokens = append(qb.tokens, "FROM", dmCommaFiled+strings.Join(tables, dmCommaSpace)+dmCommaFiled)
	return qb
}

// InnerJoin INNER JOIN the table
// InnerJoin("deploy_host")
func (qb *DmQueryBuilder) InnerJoin(table string) QueryBuilder {
	if len(strings.Split(table, " ")) > 1 {
		qb.tokens = append(qb.tokens, "INNER JOIN", dmCommaFiled+strings.Split(table, " ")[0]+dmCommaFiled+" as "+strings.Split(table, " ")[1])
		return qb
	}
	qb.tokens = append(qb.tokens, "INNER JOIN", dmCommaFiled+table+dmCommaFiled)
	return qb
}

// LeftJoin LEFT JOIN the table
// LeftJoin("deploy_host")
func (qb *DmQueryBuilder) LeftJoin(table string) QueryBuilder {
	if len(strings.Split(table, " ")) > 1 {
		qb.tokens = append(qb.tokens, "LEFT JOIN", dmCommaFiled+strings.Split(table, " ")[0]+dmCommaFiled+" as "+strings.Split(table, " ")[1])
		return qb
	}

	qb.tokens = append(qb.tokens, "LEFT JOIN", dmCommaFiled+table+dmCommaFiled)
	return qb
}

// RightJoin RIGHT JOIN the table
// InnerJoin("deploy_host")
func (qb *DmQueryBuilder) RightJoin(table string) QueryBuilder {
	if len(strings.Split(table, " ")) > 1 {
		qb.tokens = append(qb.tokens, "RIGHT JOIN", dmCommaFiled+strings.Split(table, " ")[0]+dmCommaFiled+" as "+strings.Split(table, " ")[1])
		return qb
	}

	qb.tokens = append(qb.tokens, "RIGHT JOIN", dmCommaFiled+table+dmCommaFiled)
	return qb
}

// On join with on cond
// On(`"deploy_cluster_host_rel"."sid" = "deploy_host"."sid"`)
func (qb *DmQueryBuilder) On(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "ON", cond)
	return qb
}

// Where join the Where cond
// Where(`"deploy_host"."is_deleted" = 1`)
func (qb *DmQueryBuilder) Where(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "WHERE", cond)
	return qb
}

// And join the and cond
// And(`"deploy_host"."is_deleted" = 0`)
func (qb *DmQueryBuilder) And(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "AND", cond)
	return qb
}

// Or join the or cond
func (qb *DmQueryBuilder) Or(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "OR", cond)
	return qb
}

// In join the IN (vals)
func (qb *DmQueryBuilder) In(vals ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "IN", "(", strings.Join(vals, dmCommaSpace), ")")
	return qb
}

// OrderBy join the Order by fields
// OrderBy("id")
func (qb *DmQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	for i, field := range fields {
		fields[i] = dmCommaFiled + field + dmCommaFiled
	}
	qb.tokens = append(qb.tokens, "ORDER BY", strings.Join(fields, dmCommaSpace))
	return qb
}

// Asc join the asc
func (qb *DmQueryBuilder) Asc() QueryBuilder {
	qb.tokens = append(qb.tokens, "ASC")
	return qb
}

// Desc join the desc
func (qb *DmQueryBuilder) Desc() QueryBuilder {
	qb.tokens = append(qb.tokens, "DESC")
	return qb
}

// Limit join the limit num
func (qb *DmQueryBuilder) Limit(limit int) QueryBuilder {
	qb.tokens = append(qb.tokens, "LIMIT", strconv.Itoa(limit))
	return qb
}

// Offset join the offset num
func (qb *DmQueryBuilder) Offset(offset int) QueryBuilder {
	qb.tokens = append(qb.tokens, "OFFSET", strconv.Itoa(offset))
	return qb
}

// GroupBy join the Group by fields
func (qb *DmQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "GROUP BY", strings.Join(fields, dmCommaFiled+dmCommaSpace+dmCommaFiled))
	return qb
}

// Having join the Having cond
func (qb *DmQueryBuilder) Having(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "HAVING", dmCommaFiled+cond+dmCommaFiled)
	return qb
}

// Update join the update table
func (qb *DmQueryBuilder) Update(tables ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "UPDATE", strings.Join(tables, dmCommaSpace))
	return qb
}

// Set join the set kv
func (qb *DmQueryBuilder) Set(kv ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "SET", strings.Join(kv, dmCommaSpace))
	return qb
}

// Delete join the Delete tables
func (qb *DmQueryBuilder) Delete(tables ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "DELETE")
	if len(tables) != 0 {
		qb.tokens = append(qb.tokens, strings.Join(tables, dmCommaSpace))
	}
	return qb
}

// InsertInto join the insert SQL
func (qb *DmQueryBuilder) InsertInto(table string, fields ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "INSERT INTO", table)
	if len(fields) != 0 {
		fieldsStr := strings.Join(fields, dmCommaFiled+dmCommaSpace+dmCommaFiled)
		qb.tokens = append(qb.tokens, "(", fieldsStr, ")")
	}
	return qb
}

// Values join the Values(vals)
func (qb *DmQueryBuilder) Values(vals ...string) QueryBuilder {
	valsStr := strings.Join(vals, dmCommaSpace)
	qb.tokens = append(qb.tokens, "VALUES", "(", valsStr, ")")
	return qb
}

// Subquery join the sub as alias
func (qb *DmQueryBuilder) Subquery(sub string, alias string) string {
	return fmt.Sprintf("(%s) AS %s", sub, alias)
}

// String join all tokens
func (qb *DmQueryBuilder) String() string {
	s := strings.Join(qb.tokens, " ")
	qb.tokens = qb.tokens[:0]
	return s
}
