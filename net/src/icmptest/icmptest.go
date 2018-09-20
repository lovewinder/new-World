package main

import(
	/* 
		网络
		os
		字符操作
		标准输出
	*/
	"net"
	"os"
	"bytes"
	"fmt"
	"io"
)

func main(){
	// 判断是否有服务器参数
	if len(os.Args) != 2{
	// 提示正确用法
		fmt.Println("Usage:",os.Args[0],"host")
	// 退出程序
		os.Exit(1)
	}
	// addr
	service := os.Args[1]
	
	// 获取连接
	fmt.Println("start")
	conn,err := net.Dial("ip4:icmp",service)
	fmt.Println("step 1")
	// 检查
	checkError(err)
	
	// 声明消息变量
	var msg [512]byte
	
	// 为消息变量赋值
	msg[0] = 8 // echo
	msg[1] = 0 // code 0
	msg[2] = 0 // checksum
	msg[3] = 0 // checksum
	msg[4] = 0 // identifier[0]
	msg[5] = 13 // identifier[1]
	msg[6] = 0 // sequence[0]
	msg[7] = 37 // sequence[1]
	
	// 设置消息长度
	len := 8
	
	// 完成信息校验
	check := checkSum(msg[0:len])
	fmt.Println("step 2")
	
	// 根据校验信息修改消息的值
	msg[2] = byte(check>>8)
	msg[3] = byte(check & 255)
	
	// 发送消息
	_,err = conn.Write(msg[0:len])
	fmt.Println("step 3")
	checkError(err)
	
	// 接收消息
	_,err = conn.Read(msg[0:])
	fmt.Println("step 4") // 阻塞在了step 3和step 4之间，为啥www.baidu.com不给我回复消息。
	checkError(err)
	
	
	// 输出“Got response”
	fmt.Println("Got response")
	
	
	// 如果消息的第6个字符位13，输出“Identifier matches”
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	
	// 如果消息的第8个字符位37，输出“Sequence matches”
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}
	// 退出程序
	os.Exit(0)
}

// ICMP校验
func checkSum(msg []byte) uint16 {
	sum := 0
	
	// 先假设为偶数
	for n := 1;n < len(msg) -1 ; n+= 2{
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum>>16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

// 错误校验
func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}

// 读取返回信息
func readFully(conn net.Conn) ([]byte,error){
	defer conn.Close()
	
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for{
		n,err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil,err
		}
	}
	return result.Bytes(),nil
}
	// 默认关闭连接
	
	// 定义结果变量
	
	// 声明接收用变量
	
	// 在循环中接收信息，并发信息存到缓冲区，直到写到文件末尾为止
	
	// 返回存储的信息和err