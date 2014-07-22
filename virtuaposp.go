// 银联模拟POSP，协议：ISO-8583
package virtuaposp

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func dlm_main() {
	var fWorking bool = false
	var N int

	l, err := net.Listen("tcp", "127.0.0.1:8583") // l 的类型： net.Listener
	if err == nil {
		fmt.Println("监听成功")
	} else {
		fmt.Println("监听失败")
	}

	for fWorking {
		con, err := l.Accept() // 等待客户端连接   con 的类型： net.Conn
		if err == nil {
			fmt.Println("Accept成功")
			N++
			if N >= 10 {
				fWorking = false
			} else {
				go Server(con, N)
			}
		} else {
			fmt.Println("Accept失败")
		}
	}

	time.Sleep(10 * time.Second)

	fmt.Println("main函数返回")
}

func Server(con net.Conn, N int) {
	const SIZE = 10
	var b [SIZE]byte
	sb := b[:]
	n, err := con.Read(sb) // 等待数据
	if err == nil {
		fmt.Println(N, "号协程收到了", n, "字节数据")
		fmt.Println(sb[:n])
	} else {
		fmt.Println("Read失败")
	}
}
