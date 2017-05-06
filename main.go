package main

// Author: Glen Newton
// BSD 3-Clause License

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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

var delayInMillis uint64 = 10
var realFilesOnly = false

var prevOpenFiles map[string]*SocketInfo

type Config interface{}

func init() {
	flag.Uint64Var(&delayInMillis, "d", delayInMillis, "Time granularity for checking files, in milliseconds")
	flag.BoolVar(&realFilesOnly, "r", realFilesOnly, "Show only real files, i.e. no pipes, sockets, etc.")
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
	c := make(chan []string)
	go getOpenFiles(pidDevDir, c)

	prevOpenFiles = make(map[string]*SocketInfo)

	for openFiles := range c {
		t := time.Now()
		presentlyOpenFiles := make(map[string]struct{})

		//Make hash of open files
		for i, _ := range openFiles {
			presentlyOpenFiles[openFiles[i]] = struct{}{}
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
			delete(prevOpenFiles, toBeRemoved[i])
			printFile(toBeRemoved[i], t, false, nil)
			//fmt.Printf("%s close %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), toBeRemoved[i])
		}

		// Add new files that have been opened & print them out

		for i, _ := range openFiles {
			of := openFiles[i]
			presentlyOpenFiles[of] = struct{}{}
			if _, ok := prevOpenFiles[of]; ok {
				continue
			} else {
				var socketInfo *SocketInfo = nil
				if strings.HasPrefix(openFiles[i], SocketInodePrefix) {
					inode := openFiles[i][lenSocketInodePrefix:]
					inode = inode[0 : len(inode)-1]
					socketInfo = getSocketInfo(inode)
				}

				prevOpenFiles[of] = socketInfo
				printFile(openFiles[i], t, true, socketInfo)

				//fmt.Printf("%s open %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), openFiles[i])

			}
		}

	}

}

const CLOSED = "close"
const OPEN = "open"

func printFile(filename string, t time.Time, isOpen bool, si *SocketInfo) {
	status := CLOSED
	if isOpen {
		status = OPEN
	}
	if si == nil {
		fmt.Printf("%s "+status+" %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename)
	} else {
		fmt.Printf("%s "+status+" %s %s:%d\n", t.Format("2006-01-02T15:04:05.999999-07:00"), filename, si.remoteIP, si.remotePort)
	}
}

const ProcTcp = "/proc/net/tcp"

func getSocketInfo(inode string) *SocketInfo {
	i, err := strconv.ParseInt(inode, 10, 32)
	tcpFile := ProcTcp

	f, err := os.Open(tcpFile)
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	line, err := r.ReadString(10) // 0x0A separator = newline
	for {
		line, err = r.ReadString(10) // 0x0A separator = newline
		if err == io.EOF {
			return nil
			break
		} else if err != nil {
			return nil
		}
		//fmt.Println("--------", line)
		si := NewSocketInfo(line)

		if i == si.inode {
			return si
		}
	}
	return nil

}

func getOpenFiles(d string, c chan []string) {

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
		openFiles := make([]string, 0)
		for _, f := range files {
			fullName := fdDir + "/" + f.Name()

			realFile, err := os.Readlink(fullName)
			if err != nil {
				continue
			}
			if realFilesOnly && !strings.HasPrefix(realFile, SLASH) || strings.HasPrefix(realFile, DEV) {
				continue
			}
			openFiles = append(openFiles, realFile)
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
