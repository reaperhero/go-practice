package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/ldap.v2"
	"log"
)

// Entry: 包括组，用户

// DC: Domain Component. LDAP目录类似文件系统目录dc=ygyg,dc=cn相当于/cn/ygyg
// CN: Common Name. CN有可能代表一个用户名，例如cn=admin,dc=ygyg,dc=cn表示在/cn/ygyg域下的管理员用户admin
// OU: Organizational Unit. 例如ou=People,dc=ygyg,dc=cn表示在/cn/ygyg域下的一个组织单元People

// https://www.cnblogs.com/liu-shuai/articles/12698405.html
// https://blog.csdn.net/wyansai/article/details/99416987

// ldapsearch -LLL -W -x -D "cn=admin,dc=ygyg,dc=cn" -H ldap://localhost -b "dc=ygyg,dc=cn"

var (
	ldapServer = "172.16.115.208:389"
	baseDC     = "dc=faysongg,dc=com"                                 // 根dc
	userDC     = "cn=users,cn=accounts"                               // 管理员账号
	groupDC    = "cn=groups,cn=accounts"                              // 管理员账号
	ldapUser   = fmt.Sprintf("%s,%s,%s", "uid=admin", userDC, baseDC) // 管理员账号
	ldapPass   = "12345678"                                           // 管理员密码
)

type ldapConnManagement struct {
	userSearch  string
	groupSearch string
	con         *ldap.Conn
}

func NewLdapConnManagent() *ldapConnManagement {
	conn, err := ldap.Dial("tcp", ldapServer)
	if err != nil {
		logrus.Errorf("[ldapConnManagement.NewLdapConnManagent] ldap.Dial %v", err)
		return nil
	}
	fmt.Println(ldapUser)
	err = conn.Bind(ldapUser, ldapPass)
	if err != nil {
		logrus.Errorf("[ldapConnManagement.NewLdapConnManagent] conn.Bind %v", err)
		return nil
	}
	return &ldapConnManagement{
		userSearch:  fmt.Sprintf("%s,%s", userDC, baseDC),
		groupSearch: fmt.Sprintf("%s,%s", groupDC, baseDC),
		con:         conn,
	}
}

func (l *ldapConnManagement) ListUser() {
	sql := ldap.NewSearchRequest(l.userSearch, ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false, "(uid=*)", nil, nil)

	var cur *ldap.SearchResult

	cur, err := l.con.Search(sql)
	if err != nil {
		logrus.Errorf("[ldapConnManagement.ListUser] l.con.Search %v", err)
		return
	}
	cur.Entries[2].PrettyPrint(2)
}

func (l *ldapConnManagement) ListGroup() {
	sql := ldap.NewSearchRequest(l.groupSearch, ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false, "(cn=*)", nil, nil)

	var cur *ldap.SearchResult

	cur, err := l.con.Search(sql)
	if err != nil {
		logrus.Errorf("[ldapConnManagement.ListUser] l.con.Search %v", err)
		return
	}
	cur.Entries[2].PrettyPrint(2)
}

// 添加组：cn=devGroup,dc=ygyg,dc=cn
func (l *ldapConnManagement) addGroup() {
	var (
		groupName = "cn=devGroup,"
	)
	sql := ldap.NewAddRequest(groupName + baseDC)
	sql.Attribute("objectClass", []string{"posixGroup"})
	sql.Attribute("cn", []string{"devGroup"})
	sql.Attribute("gidNumber", []string{"2000"})

	if err := l.con.Add(sql); err != nil {
		log.Print(err)
	}
}

// 添加用户
func (l *ldapConnManagement) addPeople() {
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
	//sql.Attribute("displayName", []string{username})
	//sql.Attribute("mobile", []string{params.Phone})
	//sql.Attribute("employeeID", []string{params.EmployeeID})
	//sql.Attribute("mail", []string{params.mail})
	if err := l.con.Add(sql); err != nil {
		log.Println(err)
	}
}

// 修改用户
func (l *ldapConnManagement) updatePeople() {
	var (
		user = "uid=user03,cn=devGroup,dc=ygyg,dc=cn"
	)
	sql := ldap.NewModifyRequest(user)
	sql.Replace("userPassword", []string{"12345678qwer"})
	if err := l.con.Modify(sql); err != nil {
		fmt.Println(err)
		return
	}
}

func (l *ldapConnManagement) searchPeople() {
	sql := ldap.NewSearchRequest(baseDC, ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false, "(dc=faysongg)", nil, nil)

	var cur *ldap.SearchResult

	cur, err := l.con.Search(sql)
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
