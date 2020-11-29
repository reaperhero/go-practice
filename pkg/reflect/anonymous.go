package main

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

//func main() {
//	m := Boy{User{1, "zs", 20}, "bj"}
//	t := reflect.TypeOf(m)
//	fmt.Println(t) //main.Boy
//	fmt.Printf("%#v\n", t.Field(0)) //reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x10bed40), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
//	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0)) //main.User{Id:1, Name:"zs", Age:20}
//}

type Student struct {
	Name string `json:"name1" db:"name2"`
}

//
//func main() {
//	var s Student
//	v := reflect.ValueOf(&s)
//	// 类型
//	t := v.Type()
//	// 获取字段
//	f := t.Elem().Field(0)
//	fmt.Println(f.Tag.Get("json")) // name1
//	fmt.Println(f.Tag.Get("db"))   // name2
//}
