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

type UnixSocketInfo struct {
	a        string
	refcount string
	protocol string
	flags    string
	stype    string
	st       string
	path     string
}

func NewUnixSocketInfo(s string) *UnixSocketInfo {
	var us UnixSocketInfo
	fmt.Sscanf(s, "%17s: %08s %08s %08s %04s %02s %5d %s", &us.a, &us.refcount, &us.protocol, &us.flags, &us.stype, &us.st, &us.path)
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
	_, err = r.ReadString(10) // 0x0A separator = newline
	if err != nil {
		log.Println(err)
		return nil
	}
	for {
		//line, err = r.ReadString(10) // 0x0A separator = newline
		_, err = r.ReadString(10) // 0x0A separator = newline
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Println(err)
			return nil
		}
		//fmt.Println("=== ", line)

		// if si == nil {
		// 	return nil
		// }
		// if i == si.inode {
		// 	return si
		// }
	}
	return nil

}
