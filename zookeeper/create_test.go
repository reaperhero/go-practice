package zookeeper

import (
	"fmt"
	"testing"
	"time"

	zk "github.com/samuel/go-zookeeper/zk"
)

var (
	conn *zk.Conn
)

func init() {
	zkList := []string{"localhost:2181"}
	conn, _, _ = zk.Connect(zkList, 10*time.Second)
	//defer conn.Close()
}

// create node
func Test_create_node(t *testing.T) {
	var flags int32 = 0
	//flags有4种取值：
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	conn.Create("/node", nil, flags, zk.WorldACL(zk.PermAll))               // zk.WorldACL(zk.PermAll)控制访问权限模式
	conn.Create("/node/1", []byte("data1"), flags, zk.WorldACL(zk.PermAll)) // zk.WorldACL(zk.PermAll)控制访问权限模式
}

// list node
func Test_list_node(t *testing.T) {
	childrens, _, _ := conn.Children("/node")
	fmt.Printf("%v \n", childrens) //[1]
}

func Test_get_node(t *testing.T) {
	data, stat, err := conn.Get("/node/1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data), stat) // data1 &{4294967303 4294967303 1604825087278 1604825087278 0 0 0 0 5 0 4294967303}
}
