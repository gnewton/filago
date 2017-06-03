package main

import (
	"log"
	"strconv"
)

type SocketInfoX interface {
	GetType() string
}

type SocketInfo struct {
	Inode int64           `json:"inode,omitempty"`
	Tcp   *TCPSocketInfo  `json:"tcp_socket,omitempty"`
	Unix  *UnixSocketInfo `json:"unix_socket,omitempty"`
}

func getSocketInfo(inode string) *SocketInfo {

	inodeInt, err := strconv.ParseInt(inode, 10, 32)

	info := getTCPSocketInfo(inodeInt)
	if info != nil {
		return info
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	info = getUnixSocketInfo(inodeInt)
	if info != nil {
		return info
	}
	return nil
}
