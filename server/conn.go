// Package server...
//
// Description : 包装原始的连接信息
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-24 4:35 下午
package server

import (
	"fmt"
	"net"
)

// conn 原始的连接信息
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:35 下午 2021/1/24
type Connect struct {
	net.TCPConn
}

func (c *Connect) Print() {
	fmt.Println("包装了连接信息")
}
