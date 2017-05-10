package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	//"net"
)

type UnixSocketInfo struct {
	Num      string `json:"num,omitempty"`
	Refcount string `json:"refcount,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Flags    string `json:"flags,omitempty"`
	Stype    string `json:"type,omitempty"`
	St       string `json:"st,omitempty"`
	Inode    int64  `json:"inode,omitempty"`
	Path     string `json:"path,omitempty"`
}

func NewUnixSocketInfo(s string) *UnixSocketInfo {
	var us UnixSocketInfo
	fmt.Sscanf(s, "%16s: %8s %8s %8s %4s %2s %7d %s", &us.Num, &us.Refcount, &us.Protocol, &us.Flags, &us.Stype, &us.St, &us.Inode, &us.Path)
	return &us
}

func getUnixSocketInfo(inode int64) *SocketInfo {

	f, err := os.Open(ProcNetUnix)
	if err != nil {
		log.Println("error opening file ", ProcNetUnix, err)
		return nil
	}
	defer f.Close()

	r := bufio.NewReader(f)
	// Throw away first line
	line, err := r.ReadString(10) // 0x0A separator = newline
	if err != nil {
		log.Println(err)
		return nil
	}
	for {
		line, err = r.ReadString(10) // 0x0A separator = newline

		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Println(err)
			return nil
		}
		us := NewUnixSocketInfo(line)

		if us == nil {
			return nil
		}
		//fmt.Println("++++++++++++", inode, us)
		if inode == us.Inode {
			var si SocketInfo
			si.Unix = us
			return &si
		}
	}
	return nil
}
