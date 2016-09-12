// icmptest
package main

import (
	//	"bytes"
	//	"encoding/binary"
	"fmt"
	//"io"
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
	if len(os.Args) != 2 {
		fmt.Println("usage:", os.Args[0], "host")
		os.Exit(1)
	}
	//serv := os.Args[1]
	fmt.Println("os.Args[1] =", os.Args[1])

	laddr := net.IPAddr{IP: net.ParseIP("0.0.0.0")}
	raddr, _ := net.ResolveIPAddr("ip", os.Args[1])
	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)

	//conn, err := net.Dial("ip4:icmp", serv) //拨号
	checkError(err)
	fmt.Println("dial err =", err) //
	//var msg [512]byte
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

	ck := checkSum(msg[0:len])
	msg[2] = byte(ck >> 8)
	msg[3] = byte(ck & 255)

	_, err = conn.Write(msg[0:len])
	checkError(err)
	fmt.Println("write err =", err) //

	conn.SetReadDeadline((time.Now().Add(time.Second * 5)))
	_, err = conn.Read(msg[0:])
	fmt.Println("read err", err) //
	checkError(err)
	//fmt.Println("read:", rd)

	for i := 0; i < 30; i++ {
		if i%16 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%.2x ", msg[i])
	}
	fmt.Println("")

	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("5-13, Identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("7-37, Sequence matches")
	}
	os.Exit(0)

}

func checkSum(msg []byte) uint16 {
	/*
		var (
			sum    uint32
			length int = len(msg)
			index  int
		)
		for length > 1 {
			sum += uint32(msg[index])<<8 + uint32(msg[index+1])
			index += 2
			length -= 2
		}
		if length > 0 {
			sum += uint32(msg[index])
		}
		sum += (sum >> 16)
	*/

	sum := 0
	//先假设为偶数个字
	for n := 0; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
		//fmt.Println("******sum is ", sum)
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16) //>>跟+的优先级？

	var ans uint16 = uint16(^sum)
	fmt.Println("******data is ", ans)
	return ans

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}

/*
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	return result.Bytes(), nil
}
*/
