package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

const (
	PORT = 8888
	ADDR = "127.0.0.1"
	SIZE = 100
)

func main() {
	// 1. create fd
	socketFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil || socketFd < 0 {
		fmt.Println("fail to create socketfd", err)
		os.Exit(-1)
	}

	// 2. bind addr and port
	ipv4 := net.ParseIP(ADDR).To4()
	if ipv4 == nil {
		fmt.Println("net.parseIp err")
		os.Exit(-1)
	}
	sa := &syscall.SockaddrInet4{Port: PORT}
	copy(sa.Addr[:], ipv4)

	err = syscall.Bind(socketFd, sa)
	if err != nil {
		fmt.Println("fail to bind sockaddr")
		os.Exit(-1)
	}

	// 3. listen
	err = syscall.Listen(socketFd, 128)
	if err != nil {
		fmt.Println("fail to listen")
		os.Exit(-1)
	}

	// 4. accept
	acceptFd, _, err := syscall.Accept(socketFd)
	if err != nil {
		fmt.Println("fail to accept")
		os.Exit(-1)
	}

	var (
		buf   = make([]byte, SIZE)
		readn int
		err2  error
	)

	for {
		// 5. read data
		readn, err2 = syscall.Read(acceptFd, buf)
		if err2 != nil {
			break
		}

		if readn > 0 {
			// 5. write data
			writen, _ := syscall.Write(acceptFd, buf[:readn])
			if writen < 0 {
				fmt.Println("fail to write acceptFd")
				break
			}
		} else if readn == 0 {
			fmt.Println("客户端已关闭")
			break
		} else {
			fmt.Println("fail to read acceptFd")
			break
		}
	}

	// 7. close socket
	_ = syscall.Close(socketFd)
	_ = syscall.Close(acceptFd)
}
