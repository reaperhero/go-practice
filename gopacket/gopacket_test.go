package gopacket

import (
	"fmt"
	"github.com/google/gopacket"
	_ "github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
	"testing"
	"time"
)

// go 抓包工具
// https://studygolang.com/articles/35581


// 获取网卡名称
func TestName(t *testing.T) {
	devices, _ := pcap.FindAllDevs()
	for _, device := range devices {
		fmt.Println("\nname:", device.Name)
		fmt.Println("describe:", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("IP:", address.IP)
			fmt.Println("mask:", address.Netmask)
		}
	}
}

var (
	device   string        = `en0`            //网卡名
	snapshot int32         = 65535            //读取一个数据包的最大值，一般设置成这65535即可
	promisc  bool          = true             //是否开启混杂模式
	timeout  time.Duration = time.Second * -1 //抓取数据包的超时时间，负数表示立即刷新，一般都设为负数
)

// 抓包
func TestName1(t *testing.T) {
	//获取一个网卡句柄
	handle, err := pcap.OpenLive(device, snapshot, promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}
	//别忘了释放句柄
	defer handle.Close()

	//NewPacketSource新建一个数据包数据源
	//捕捉一个数据包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packet, err := packetSource.NextPacket() //返回一个数据包
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(packet)

	//一直捕捉数据包
	p := packetSource.Packets() //返回一个channel
	for data := range p {
		fmt.Println(data)
	}
}

// 包分析
func TestName2(t *testing.T) {
	handle, err := pcap.OpenLive(device, snapshot, promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}

	err = handle.SetBPFFilter("port 80 and tcp")//过滤规则
	if err != nil {
		return 
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetChan := packetSource.Packets()

	//分析每一个网络报文
	for packet := range packetChan {
		fmt.Println("---------Layer Type------------")
		for _, layer := range packet.Layers() {
			fmt.Println(layer.LayerType()) //打印这个包里面每一层的类型
		}

		//IP协议
		ipLayer := packet.Layer(layers.LayerTypeIPv4) //返回layer接口
		if ipLayer != nil {
			fmt.Println("-----------IP Layer-----------")
			ip, _ := ipLayer.(*layers.IPv4) //通过类型断言，得到对应协议的结构体
			fmt.Println("protocol:", ip.Protocol)
			fmt.Println("dstIP:", ip.DstIP)
			fmt.Println("srcIP:", ip.SrcIP)
		} else {
			fmt.Println("No IPv4 layer")
		}

		//TCP协议
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			fmt.Println("-----------TCP Layer-----------")
			tcp, _ := tcpLayer.(*layers.TCP)
			fmt.Println("ack:", tcp.Ack)
			fmt.Println("checksum:", tcp.Checksum)
		}

		//应用层
		appLayer := packet.ApplicationLayer()
		if appLayer == nil {
			continue
		}

		b := appLayer.Payload()
		if b == nil {
			continue
		}
		if strings.Contains(string(b), "HTTP") {
			fmt.Println("http packer ......")
		}

		fmt.Printf("\n\n")
	}
}
