package hystrix

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"testing"
	"time"
)

// Timeout: 执行 command 的超时时间。
// MaxConcurrentRequests：command 的最大并发量 。
// SleepWindow：当熔断器被打开后，SleepWindow 的时间就是控制过多久后去尝试服务是否可用了。
// RequestVolumeThreshold： 一个统计窗口 10 秒内请求数量。达到这个请求数量后才去判断是否要开启熔断
// ErrorPercentThreshold：错误百分比，请求数量大于等于 RequestVolumeThreshold 并且错误率到达这个百分比后就会启动熔断

func TestHystrix(t *testing.T) {
	hystrix.ConfigureCommand("wuqq", hystrix.CommandConfig{
		Timeout:                int(3 * time.Second),
		MaxConcurrentRequests:  10,
		SleepWindow:            5000,
		RequestVolumeThreshold: 10,
		ErrorPercentThreshold:  30,
	})
	_ = hystrix.Do("wuqq", func() error {
		// talk to other services
		_, err := http.Get("https://www.baidu.com/")
		if err != nil {
			return err
		}
		return nil
	}, func(err error) error {
		fmt.Printf("handle  error:%v\n", err)
		return nil
	})

}
