package main

// Author: Glen Newton
// BSD 3-Clause License

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const CLOSED = "close"
const OPEN = "open"

func printFile(filename string, t *time.Time, isOpen bool, fi *FDInfo) {
	if jsonOut {
		printJson(filename, t, isOpen, fi)
	} else {
		printText(filename, t, isOpen, fi)
	}

}

func printText(filename string, t *time.Time, isOpen bool, fi *FDInfo) {
	si := fi.SocketInfo
	status := CLOSED
	if isOpen {
		status = OPEN
	}
	if si == nil {
		fmt.Printf("%s "+status+" %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename)
	} else {
		if fi.Type == TCPSocket {
			fmt.Printf("%s "+status+" %s %s %s:%d %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename, fi.Type, si.Tcp.RemoteIP, si.Tcp.RemotePort, getRemoteHostname(si.Tcp.RemoteIP.String()))
		} else {
			if si.Unix != nil && si.Unix.Path == "" {
				si.Unix.Path = EmptyValue
				fmt.Printf("%s "+status+" %s %s %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename, fi.Type, si.Unix.Path)
			} else {
				fmt.Printf("%s "+status+" %s %s %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename, fi.Type)
			}

		}
	}
}

func printJson(filename string, t *time.Time, isOpen bool, fi *FDInfo) {

	si := fi.SocketInfo

	if si != nil && si.Tcp != nil {
		si.Tcp.RemoteHostName = getRemoteHostname(si.Tcp.RemoteIP.String())
	}
	if isOpen {
		fi.Status = OPEN
	} else {
		fi.Status = CLOSED
	}
	jsonBytes, err := json.Marshal(fi)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(jsonBytes))
}
