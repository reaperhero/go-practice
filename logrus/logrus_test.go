package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/orandin/lumberjackrus"
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"

	"os"
	"time"
)

// 支持七种日志级别：Trace、Debug、Info、Warn、Error、Fatal、Panic
// 支持自定义日志格式，内置两种格式 JSONFormatter（JSON 格式） 和 TextFormatter（文本格式），并允许用户通过实现 Formatter 接口来自定义日志格式。

func init() {
	// 显示行号
	log.SetReportCaller(true)

	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	//log.SetFormatter(&log.TextFormatter{ // 输出格式 logfmt 风格
	//	DisableColors: true,
	//	FullTimestamp: true,
	//})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	//file, err := os.OpenFile(LOG_FILE, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0755)
	log.SetOutput(os.Stdout)

	// 同时写到多个输出
	//w1 := os.Stdout
	//w2, _ := os.OpenFile("demo.log", os.O_WRONLY|os.O_CREATE, 0644)
	//log.SetOutput(io.MultiWriter(w1, w2))

	// 只记录严重或以上警告
	log.SetLevel(log.WarnLevel)

	// 通过hock
	log.AddHook(&emailHook{})
	log.AddHook(newRotateHook())
}

type emailHook struct{}

func (hook *emailHook) Levels() []log.Level {
	return log.AllLevels // 所有日志级别都会执行 Fire 方法
}

func (hook *emailHook) Fire(entry *log.Entry) error {
	// 修改日志内容
	entry.Data["app"] = "email"
	// 发送邮件
	msg, _ := entry.String()
	fmt.Printf("fakeSendEmail: %s", msg)
	return nil
}

func newRotateHook() log.Hook {
	hook, _ := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{ // 通用日志配置
			Filename:   "general.log",
			MaxSize:    100,
			MaxBackups: 1,
			MaxAge:     1,
			Compress:   false,
			LocalTime:  false,
		},
		log.InfoLevel,
		&log.TextFormatter{DisableColors: true},
		&lumberjackrus.LogFileOpts{ // 针对不同日志级别的配置
			log.TraceLevel: &lumberjackrus.LogFile{
				Filename: "trace.log",
			},
			log.ErrorLevel: &lumberjackrus.LogFile{
				Filename:   "error.log",
				MaxSize:    10,    // 日志文件在轮转之前的最大大小，默认 100 MB
				MaxBackups: 10,    // 保留旧日志文件的最大数量
				MaxAge:     10,    // 保留旧日志文件的最大天数
				Compress:   true,  // 是否使用 gzip 对日志文件进行压缩归档
				LocalTime:  false, // 是否使用本地时间，默认 UTC 时间
			},
		},
	)
	return hook
}

func TestColorlog(t *testing.T) {
	log := log.New()
	log.Formatter = &CustomTextFormatter{
		ForceColors:   true,
		ColorInfo:     color.New(color.FgBlue),
		ColorWarning:  color.New(color.FgYellow),
		ColorError:    color.New(color.FgRed),
		ColorCritical: color.New(color.BgRed, color.FgWhite),
	}

	log.Info("This is an info message")     // 输出蓝色的信息日志
	log.Warn("This is a warning message")   // 输出黄色的警告日志
	log.Error("This is an error message")   // 输出红色的错误日志
	log.Fatal("This is a critical message") // 输出带有红色背景和白色文本的严重日志

	log.Panic("This is a critical message with panic") // 输出带有红色背景和白色文本的严重日志，并引发 panic
}

// 自定义格式化器，继承自 logrus.TextFormatter
type CustomTextFormatter struct {
	log.TextFormatter
	ForceColors   bool
	ColorInfo     *color.Color
	ColorWarning  *color.Color
	ColorError    *color.Color
	ColorCritical *color.Color
}

// 格式化方法，用于将日志条目格式化为字节数组
func (f *CustomTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	if f.ForceColors {
		switch entry.Level {
		case log.InfoLevel:
			f.ColorInfo.Println(entry.Message) // 使用蓝色打印信息日志
		case log.WarnLevel:
			f.ColorWarning.Println(entry.Message) // 使用黄色打印警告日志
		case log.ErrorLevel:
			f.ColorError.Println(entry.Message) // 使用红色打印错误日志
		case log.FatalLevel, log.PanicLevel:
			f.ColorCritical.Println(entry.Message) // 使用带有红色背景和白色文本的样式打印严重日志
		default:
			f.PrintColored(entry)
		}
		return nil, nil
	} else {
		return f.TextFormatter.Format(entry)
	}
}

// 自定义方法，用于将日志条目以带颜色的方式打印出来
func (f *CustomTextFormatter) PrintColored(entry *log.Entry) {
	levelColor := color.New(color.FgCyan, color.Bold)             // 定义蓝色和粗体样式
	levelText := levelColor.Sprintf("%-6s", entry.Level.String()) // 格式化日志级别文本

	msg := levelText + " " + entry.Message
	if entry.HasCaller() {
		msg += " (" + entry.Caller.File + ":" + strconv.Itoa(entry.Caller.Line) + ")" // 添加调用者信息
	}

	fmt.Fprintln(color.Output, msg) // 使用有颜色的方式打印消息到终端
}
