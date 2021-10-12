package main

import (
	"fmt"
	"gopkg.in/ldap.v2"
	"log"
)

// Entry: 包括组，用户

// https://www.cnblogs.com/liu-shuai/articles/12698405.html
// ldapsearch -LLL -W -x -D "cn=admin,dc=ygyg,dc=cn" -H ldap://localhost -b "dc=ygyg,dc=cn"

var (
	ldapServer = "192.168.31.192:389"
	baseDn     = "dc=ygyg,dc=cn"          // 管理员账号
	ldapUser   = "cn=admin,dc=ygyg,dc=cn" // 管理员账号
	ldapPass   = "qwer456"                // 管理员密码
	conn, _    = ldap.Dial("tcp", ldapServer)
	domain     = "dc=ygyg,dc=cn"
)

func init() {
	conn.Bind(ldapUser, ldapPass)
}

func main() {
	addPeople()
	searchPeople()
}

// 添加组：cn=devGroup,dc=ygyg,dc=cn
func addGroup() {
	var (
		groupName = "cn=devGroup,"
	)
	sql := ldap.NewAddRequest(groupName + domain)
	sql.Attribute("objectClass", []string{"posixGroup"})
	sql.Attribute("cn", []string{"devGroup"})
	sql.Attribute("gidNumber", []string{"2000"})

	if err := conn.Add(sql); err != nil {
		log.Print(err)
	}
}

// 添加用户
func addPeople() {
	var (
		user = "uid=user03,cn=devGroup,dc=ygyg,dc=cn"
	)
	sql := ldap.NewAddRequest(user)
	sql.Attribute("objectClass", []string{"inetOrgPerson", "posixAccount", "shadowAccount"})
	sql.Attribute("homeDirectory", []string{"/home/user01"})
	sql.Attribute("userPassword", []string{"12345678"})
	sql.Attribute("cn", []string{"user03"})
	sql.Attribute("sn", []string{"linux"})
	sql.Attribute("uid", []string{"user03"})
	sql.Attribute("loginShell", []string{"/bin/bash"})
	sql.Attribute("uidNumber", []string{"10000"})
	sql.Attribute("gidNumber", []string{"2000"})
	if err := conn.Add(sql); err != nil {
		log.Println(err)
	}
}

// 修改用户
func updatePeople() {
	var (
		user = "uid=user03,cn=devGroup,dc=ygyg,dc=cn"
	)
	sql := ldap.NewModifyRequest(user)
	sql.Replace("userPassword", []string{"12345678qwer"})
	if err := conn.Modify(sql); err != nil {
		fmt.Println(err)
		return
	}
}

func searchPeople() {
	var (
		userDn   = "cn=devGroup,dc=ygyg,dc=cn"
		username = "user03"
	)
	sql := ldap.NewSearchRequest(
		userDn,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		fmt.Sprintf("(uid=%s)",username),
		nil,
		nil,
	)

	var cur *ldap.SearchResult

	cur, err := conn.Search(sql)
	if err != nil {
		log.Println(err)
		return
	}
	cur.Print()
	//DN: uid=user03,cn=devGroup,dc=ygyg,dc=cn
	//objectClass: [inetOrgPerson posixAccount shadowAccount]
	//homeDirectory: [/home/user01]
	//userPassword: [12345678]
	//cn: [user03]
	//sn: [linux]
	//uid: [user03]
	//loginShell: [/bin/bash]
	//uidNumber: [10000]
	//gidNumber: [2000]
}
