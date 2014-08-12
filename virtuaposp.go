// 模拟POSP，协议：ISO-8583
package virtuaposp

import (
	"fmt"
	"net"
	"time"
)

const (
	LLVAR  = 2
	LLLVAR = 3
)

const (
	FORMAT_BCD = iota
	FORMAT_ASC
	FORMAT_BIN
)

const (
	AS_A   = 0x00000001
	AS_N   = 0x00000002
	AS_S   = 0x00000004
	AS_AN  = AS_A | AS_N
	AS_AS  = AS_A | AS_S
	AS_NS  = AS_N | AS_S
	AS_ANS = AS_A | AS_N | AS_S
)

const (
	JUSTIFY_NONE = iota
	JUSTIFY_LEFT
	JUSTIFY_RIGHT
)

type Iso8583Field struct {
	name          string // 域的名称
	number        int    // 编号
	AdmissibleSet uint32 // 允许什么字符
	LengthVar     bool   // 是变长的吗？
	LengthSize    int    // 对于变长，用多少个字节表示长度（LLVAR或LLLVAR）；对于定长，无意义
	Length        int    // 长度（对于定长）或最大长度（对于变长）
	Format        int    // 格式 BCD ASC BIN
	duiqi         int    // 左靠还是右靠
	FillChar      byte   // 补充什么字符
}

type TransRuler struct {
	HaveHead bool // 是否要报文头
	Ruler    [128]Iso8583Field
}

func GetUnionPayRuler() *TransRuler {
	rlr := [128]Iso8583Field{
		{"交易处理码", 1, AS_N, false, 0, 6, FORMAT_BCD, 0, 0},
		{"交易处理", 2, AS_N, false, 0, 6, FORMAT_BCD, 0, 0},
		{"主账号", 2, AS_N, true, LLVAR, 19, FORMAT_BCD, JUSTIFY_LEFT, '0'},
	}

	tr := TransRuler{
		true,
		rlr,
	}

	return &tr
}

func BeginServe(port uint16) {
	var fWorking bool = false
	var N int

	var Addr string = fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", Addr) // l 的类型： net.Listener
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
