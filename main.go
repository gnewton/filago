package main

// Author: Glen Newton
// BSD 3-Clause License

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const PROC = "/proc/"

const FD = "/fd"
const ProcNetUnix = "/proc/net/unix"
const ProcNetTcp = "/proc/net/tcp"

const SLASH = string(filepath.Separator)
const DEV = "/dev/"
const EmptyValue = "-"

// Types of files:
const File = "file"
const Pipe = "pipe"
const TCPSocket = "tcp"
const UnixSocket = "unix"
const AnonInode = "anon_inode"

//Inode prefixes
const SocketInodePrefix = "socket:["
const PipeInodePrefix = "pipe:["
const AnonInodePrefix = "anon_inode:"

var delayInMillis uint64 = 100
var realFilesOnly = false
var lookupHostnames = false
var jsonOut = false

var prevOpenFiles map[string]*FDInfo

type Config interface{}

type FDInfo struct {
	Filename   string      `json:"filename,omitempty"`
	Type       string      `json:"type,omitempty"`
	SocketInfo *SocketInfo `json:"socket_info,omitempty"`
	Status     string      `json:"status,omitempty"`
}

func init() {
	flag.Uint64Var(&delayInMillis, "d", delayInMillis, "Time granularity for checking files, in milliseconds")
	flag.BoolVar(&realFilesOnly, "r", realFilesOnly, "Show only real files, i.e. no pipes, sockets, etc.")
	flag.BoolVar(&lookupHostnames, "l", lookupHostnames, "Turn on hostname lookup (default is a \""+EmptyValue+"\"")
	flag.BoolVar(&jsonOut, "j", jsonOut, "Output json (complete json per line)")

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
			presentlyOpenFiles[openFiles[i].Filename] = openFiles[i]
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
				printFile(toBeRemoved[i], &t, false, fdInfo)
			} else {
				printFile(toBeRemoved[i], &t, false, nil)
			}

			delete(prevOpenFiles, toBeRemoved[i])

			//fmt.Printf("%s close %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), toBeRemoved[i])
		}

		// Add new files that have been opened & print them out

		for i, _ := range openFiles {
			of := openFiles[i]
			presentlyOpenFiles[of.Filename] = of
			if _, ok := prevOpenFiles[of.Filename]; ok {
				continue
			} else {
				newFd := FDInfo{Filename: openFiles[i].Filename}

				if strings.HasPrefix(openFiles[i].Filename, SocketInodePrefix) {
					// This should be altered so only run once, outside of loop, not for each file
					// as we read /proc/net/[tcp|unux] completely for each file...BAD FIXXX
					// Expensive
					//
					newFd.SocketInfo = getSocketInfo(extractInode(openFiles[i].Filename))
					if newFd.SocketInfo != nil {
						if newFd.SocketInfo.Tcp != nil {
							newFd.Type = "tcp"
						} else {
							if newFd.SocketInfo.Unix != nil {
								newFd.Type = "unix"
							} else {
								newFd.Type = "other"
							}
						}
					} else {
						newFd.Type = "other"
					}
				} else {
					if strings.HasPrefix(openFiles[i].Filename, PipeInodePrefix) {
						newFd.Type = Pipe
					} else {
						if strings.HasPrefix(openFiles[i].Filename, AnonInodePrefix) {
							newFd.Type = AnonInode
						} else {
							newFd.Type = File
						}
					}

				}

				prevOpenFiles[of.Filename] = &newFd
				printFile(openFiles[i].Filename, &t, true, &newFd)

				//fmt.Printf("%s open %s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), openFiles[i])

			}
		}

	}

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
			fdInfo.Filename = realFile
			if strings.HasPrefix(realFile, SocketInodePrefix) {
				fdInfo.SocketInfo = getSocketInfo(extractInode(realFile))
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
