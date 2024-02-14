package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destiation string, port string) string {
	address := destiation + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[Down]%v is unreachable, \n Error: %v", destiation, err)
	} else {
		status = fmt.Sprintf("[Up]%v is reachable, \n From: %v \n To %v ", destiation, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
