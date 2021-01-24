// Package server...
//
// Description : 约定tcp服务需要实现的方法
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-24 4:10 下午
package server

// IServer TCP服务器的接口
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:13 下午 2021/1/24
type IServer interface {
	Start()
	Stop()
	Server()
}
