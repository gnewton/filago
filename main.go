package main

// Author: Glen Newton
// BSD 3-Clause License

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const PROC = "/proc/"
const FD = "/fd"
const SLASH = "/"
const DEV = "/dev/"

var delayInMillis uint64 = 100
var realFilesOnly = false

var prevOpenFiles map[string]*FDInfo

type Config interface{}

func init() {
	flag.Uint64Var(&delayInMillis, "d", delayInMillis, "Time granularity for checking files, in milliseconds")
	flag.BoolVar(&realFilesOnly, "r", realFilesOnly, "Show only real files, i.e. no pipes, sockets, etc.")

	initCache()
}

func handleParameters() {
	flag.Parse()

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	handleParameters()

	if len(flag.Args()) != 1 {
		log.Println(flag.Args())
		return
	}
	pid := flag.Args()[0]

	_, err := strconv.Atoi(pid)
	if err != nil {
		log.Fatal("Must be an process number (integer): " + pid)
	}

	listOpenFiles(pid, nil)
}

const SocketInodePrefix = "socket:["

var lenSocketInodePrefix = len(SocketInodePrefix)

func listOpenFiles(pid string, config Config) {
	pidDevDir := PROC + pid

	exists, err := exists(pidDevDir)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		return
	}
	c := make(chan []*FDInfo)
	go getOpenFiles(pidDevDir, c)

	prevOpenFiles = make(map[string]*FDInfo)

	for openFiles := range c {
		t := time.Now()
		presentlyOpenFiles := make(map[string]*FDInfo)

		//Make hash of open files
		for i, _ := range openFiles {
			presentlyOpenFiles[openFiles[i].filename] = openFiles[i]
		}

		//Find files no longer open
		toBeRemoved := make([]string, 0)
		for f, _ := range prevOpenFiles {
			if _, ok := presentlyOpenFiles[f]; !ok {
				toBeRemoved = append(toBeRemoved, f)
			}
		}

		// Remove files no longer open & print them out
		for i, _ := range toBeRemoved {
			fdInfo, ok := prevOpenFiles[toBeRemoved[i]]
			if ok {
				printFile(toBeRemoved[i], &t, false, fdInfo.socketInfo)
			} else {
				printFile(toBeRemoved[i], &t, false, nil)
			}

			delete(prevOpenFiles, toBeRemoved[i])

			//fmt.Printf("%s close %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), toBeRemoved[i])
		}

		// Add new files that have been opened & print them out

		for i, _ := range openFiles {
			of := openFiles[i]
			presentlyOpenFiles[of.filename] = of
			if _, ok := prevOpenFiles[of.filename]; ok {
				continue
			} else {
				var socketInfo *SocketInfo = nil
				if strings.HasPrefix(openFiles[i].filename, SocketInodePrefix) {
					// This should be altered so only run once, outside of loop, not for each file
					// as we read /proc/net/[tcp|unux] completely for each file...BAD FIXXX
					// Expensive
					//
					socketInfo = getSocketInfo(extractInode(openFiles[i].filename))
				}

				prevOpenFiles[of.filename] = &(FDInfo{filename: openFiles[i].filename, socketInfo: socketInfo})
				printFile(openFiles[i].filename, &t, true, socketInfo)

				//fmt.Printf("%s open %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), openFiles[i])

			}
		}

	}

}

const CLOSED = "close"
const OPEN = "open"

func printFile(filename string, t *time.Time, isOpen bool, si *SocketInfo) {

	status := CLOSED
	if isOpen {
		status = OPEN
	}
	if si == nil {
		fmt.Printf("%s "+status+" %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename)
	} else {
		if si.stype == TCPSocket {
			fmt.Printf("%s "+status+" %s %s %s:%d %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename, si.stype, si.tcp.remoteIP, si.tcp.remotePort, getRemoteHostname(si.tcp.remoteIP.String()))
		} else {
			if si.unix.path == "" {
				si.unix.path = "-"
			}
			fmt.Printf("%s "+status+" %s %s %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename, si.stype, si.unix.path)
		}
	}
}

type FDInfo struct {
	filename   string
	socketInfo *SocketInfo
}

//func getOpenFiles(d string, c chan []string) {
func getOpenFiles(d string, c chan []*FDInfo) {

	fdDir := d + FD

	ticker := time.NewTicker(time.Millisecond * time.Duration(delayInMillis))
	for _ = range ticker.C {

		exists, err := exists(fdDir)
		if err != nil {
			close(c)
			log.Fatal(err)
		}
		if !exists {
			close(c)
			return
		}
		files, _ := ioutil.ReadDir(fdDir)
		openFiles := make([]*FDInfo, 0)
		for _, f := range files {
			fullName := fdDir + "/" + f.Name()

			realFile, err := os.Readlink(fullName)
			if err != nil {
				continue
			}
			if realFilesOnly && !strings.HasPrefix(realFile, SLASH) || strings.HasPrefix(realFile, DEV) {
				continue
			}
			fdInfo := new(FDInfo)
			fdInfo.filename = realFile
			if strings.HasPrefix(realFile, SocketInodePrefix) {
				fdInfo.socketInfo = getSocketInfo(extractInode(realFile))
			}
			//openFiles = append(openFiles, realFile)
			openFiles = append(openFiles, fdInfo)
		}
		c <- openFiles
	}
	close(c)
}

// From: https://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func extractInode(s string) string {
	s = s[lenSocketInodePrefix:]
	return s[0 : len(s)-1]
}
