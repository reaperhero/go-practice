package stand


// Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换。

// 一、string与int类型转换

// Atoi()函数用于将字符串类型的整数转换为int类型
// func Atoi(s string) (i int, err error)

// Itoa()函数用于将int类型数据转换为对应的字符串
// func Itoa(i int) string


// 二、Parse

// 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
// func ParseBool(str string) (value bool, err error)

// 返回字符串表示的整数值，接受正负号，
// base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
// bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
// func ParseInt(s string, base int, bitSize int) (i int64, err error)

// ParseUint类似ParseInt但不接受正负号，用于无符号整型。
// func ParseUint(s string, base int, bitSize int) (n uint64, err error)

// 解析一个表示浮点数的字符串并返回其值。
// func ParseFloat(s string, bitSize int) (f float64, err error)



// 三、 Format
//  FormatBool，FormatFloat，FormatInt 和  FormatUint将值转换为字符串
//s := strconv.FormatBool(true)
//s := strconv.FormatFloat(3.1415, 'E', -1, 64)
//s := strconv.FormatInt(-42, 16)
//s := strconv.FormatUint(42, 16)