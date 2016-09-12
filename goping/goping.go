// goping
package main

import (
	//	"bytes"
	//	"container/list"
	//	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func main() {
	fmt.Println("Hello World!")
	arg_num := len(os.Args)

	if arg_num < 2 {
		fmt.Print(
			"Please runAs [super user] in [terminal].\n",
			"Usage:\n",
			"\tgoping url\n",
			"\texample: goping www.baidu.com",
		)
		time.Sleep(5e9)
		return
	}

	var (
		//		icmp     ICMP
		laddr    = net.IPAddr{IP: net.ParseIP("0.0.0.0")}
		raddr, _ = net.ResolveIPAddr("ip", os.Args[1])
	)

	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer conn.Close()
	/*
		icmp.Type = 8
		icmp.Code = 0
		icmp.Checksum = 0
		icmp.Identifier = 0
		icmp.SequenceNum = 0
	*/
	msg := make([]byte, 512)
	msg[0] = 8  //echo
	msg[1] = 0  //code 0
	msg[2] = 0  //checksum
	msg[3] = 0  //checksum
	msg[4] = 0  //identifier[0]
	msg[5] = 13 //identifier[1]
	msg[6] = 0  //sequence[0]
	msg[7] = 37 //sequence[1]
	len := 8

	ck := CheckSum(msg[0:len])
	msg[2] = byte(ck >> 8)
	msg[3] = byte(ck & 255)

	_, err = conn.Write(msg[0:len])
	fmt.Println("write err =", err)
	_, err = conn.Read(msg[0:])
	fmt.Println("read err", err)
	/*
		var buffer bytes.Buffer
		binary.Write(&buffer, binary.BigEndian, icmp)
		icmp.Checksum = CheckSum(buffer.Bytes())
		buffer.Reset()
		binary.Write(&buffer, binary.BigEndian, icmp)
	*/
	fmt.Printf("\n正在 Ping %s 具有 0 字节的数据:\n", raddr.String())
	//	recv := make([]byte, 512)

	//	statistic := list.New()
	//	sended_packets := 0

	/*
		for i := 4; i > 0; i-- {

			if _, err := conn.Write(buffer.Bytes()); err != nil {
				fmt.Println(err.Error())
				return
			}
			sended_packets++
			t_start := time.Now()

			conn.SetReadDeadline((time.Now().Add(time.Second * 5)))
			_, err := conn.Read(recv)

			if err != nil {
				fmt.Println("请求超时")
				continue
			}

			t_end := time.Now()

			dur := t_end.Sub(t_start).Nanoseconds() / 1e6

			fmt.Printf("来自 %s 的回复: 时间 = %dms\n", raddr.String(), dur)

			statistic.PushBack(dur)
	*/
	//打印出接收到的数据
	for i := 0; i < 30; i++ {
		if i%16 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%.2x ", msg[i])
	}
	fmt.Println("")

	//}
	/*
		defer func() {
			fmt.Println("")
			//信息统计
			var min, max, sum int64
			if statistic.Len() == 0 {
				min, max, sum = 0, 0, 0
			} else {
				min, max, sum = statistic.Front().Value.(int64), statistic.Front().Value.(int64), int64(0)
			}

			for v := statistic.Front(); v != nil; v = v.Next() {

				val := v.Value.(int64)

				switch {
				case val < min:
					min = val
				case val > max:
					max = val
				}

				sum = sum + val
			}
			recved, losted := statistic.Len(), sended_packets-statistic.Len()
			fmt.Printf("%s 的 Ping 统计信息：\n  数据包：已发送 = %d，已接收 = %d，丢失 = %d (%.1f%% 丢失)，\n往返行程的估计时间(以毫秒为单位)：\n  最短 = %dms，最长 = %dms，平均 = %.0fms\n",
				raddr.String(),
				sended_packets, recved, losted, float32(losted)/float32(sended_packets)*100,
				min, max, float32(sum)/float32(recved),
			)
		}()
	*/
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}
