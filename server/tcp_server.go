// Package server...
//
// Description : tcp 服务器
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-24 4:13 下午
package server

import (
	"fmt"
	"log"
	"net"
)

type tcpServer struct {
	Name      string
	IP        string
	IPVersion string
	Port      int
}

func (ts *tcpServer) Start() {
	log.Printf("%s %s:%d start...\n", ts.Name, ts.IP, ts.Port)
	addr, err := net.ResolveTCPAddr(ts.IPVersion, fmt.Sprintf("%s:%d", ts.IP, ts.Port))
	if err != nil {
		log.Println("resolve tcp addr err ", err)
		return
	}
	listener, err := net.ListenTCP(ts.IPVersion, addr)
	if err != nil {
		log.Println("listen tcp err ", err)
		return
	}

	go func() {
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("accept tcp err ", err)
				continue
			}

			go func() {

				for {
					b := make([]byte, 512)
					length, err := conn.Read(b)
					if err != nil {
						conn.Close()
						log.Println("read tcp err ", err)
						break
					}
					if _, err := conn.Write(b[:length]); err != nil {
						log.Println("write tcp err ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (ts *tcpServer) Stop() {
	fmt.Println("连接已断开")
}
func (ts *tcpServer) Server() {
	ts.Start()
	select{}
}
func NewServer(name string) IServer {
	s := &tcpServer{
		Name:      name,
		IP:        "0.0.0.0",
		IPVersion: "tcp4",
		Port:      19000,
	}
	return s
}