package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	//"net"
)

const ProcNetUnix = "/proc/net/unix"
const UnixSocket = "unix"

type UnixSocketInfo struct {
	num      string
	refcount string
	protocol string
	flags    string
	stype    string
	st       string
	inode    int64
	path     string
}

func NewUnixSocketInfo(s string) *UnixSocketInfo {
	var us UnixSocketInfo
	fmt.Sscanf(s, "%16s: %8s %8s %8s %4s %2s %7d %s", &us.num, &us.refcount, &us.protocol, &us.flags, &us.stype, &us.st, &us.inode, &us.path)
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
		if inode == us.inode {
			var si SocketInfo
			si.unix = us
			si.stype = UnixSocket
			return &si
		}
	}
	return nil
}
