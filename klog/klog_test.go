package klog

import (
	"flag"
	"k8s.io/klog/v2"
	"testing"
)

// klog提供了Info、InfoDepth、Infoln、Infof、InfoS方法；
//
//	Info使用logging.print打印info级别的日志，参数的处理跟fmt.Print类似，若没有换行则会追加换行；
//	InfoDepth可以指定要打印的call frame，Info使用的depth为0；
//	Infoln的参数处理与fmt.Println类似，总是会新添加换行；
//	Infof的参数处理与fmt.Printf类似，若没有换行则会追加换行；
//	InfoS用于打印结构化的日志，kv之间用=连接，总是会新添加换行
func TestName(t *testing.T) {
	klog.InitFlags(flag.CommandLine)
	defer klog.Flush()
	klog.Info("hello by Info")
	klog.InfoDepth(0, "hello by InfoDepth 0")
	klog.InfoDepth(1, "hello by InfoDepth 1")
	klog.Infoln("hello by Infoln")
	klog.Infof("hello by %s", "Infof")
	klog.InfoS("Pod status updated", "pod", "kubedns", "status", "ready")
}
