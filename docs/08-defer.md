# defer

defer规则

规则一：延迟函数的参数在defer语句出现时就已经确定下来了

规则二：延迟函数执行按后进先出顺序执行，即先出现的defer最后执行

规则三：延迟函数可能操作主函数的具名返回值

- 
```
func increaseA() int {
    var i int
    defer func() {
        i++
    }()
    return i
}

func increaseB() (r int) {
    defer func() {
        r++
    }()
    return r
}

func main() {
    fmt.Println(increaseA()) // 0
    fmt.Println(increaseB()) // 1
}
```

- 

```
type Person struct {
    age int
}

func main() {
    person := &Person{28}

    defer fmt.Println(person.age) //person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中

    defer func(p *Person) {
        fmt.Println(p.age) // defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29
    }(person)  

    defer func() {
        fmt.Println(person.age) // 闭包引用，输出 29；
    }()

    person.age = 29
}
```


- 

```
type Person struct {
    age int
}

func main() {
    person := &Person{28}

    defer fmt.Println(person.age)  // person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中

    defer func(p *Person) {
        fmt.Println(p.age)     // defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变
    }(person)

    defer func() {
        fmt.Println(person.age)  // 闭包引用，person 的值已经被改变，指向结构体 Person{29}，所以输出 29
    }()

    person = &Person{29} 
}
``` 