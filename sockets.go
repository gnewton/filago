package main

import (
	"log"
	"strconv"
)

type SocketInfoX interface {
	GetType() string
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
	return getUnixSocketInfo(inodeInt)

}
