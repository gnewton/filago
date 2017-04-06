package main

// See https://github.com/marpie/go-mjpeg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var prevOpenFiles map[string]struct{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if len(os.Args) != 2 {
		usage()
		return
	}
	pid := os.Args[1]

	_, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Must be an integer")
	}

	pidDevDir := "/proc/" + pid

	exists, err := exists(pidDevDir)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		return
	}
	//fmt.Println(pid)

	c := make(chan []string)
	go getOpenFiles(pidDevDir, c)

	prevOpenFiles = make(map[string]struct{})

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
		// Remove files no longer open
		for i, _ := range toBeRemoved {
			delete(prevOpenFiles, toBeRemoved[i])
			fmt.Printf("%s\tclose\t%s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), toBeRemoved[i])
		}

		// Add new files that have been opend
		for i, _ := range openFiles {
			of := openFiles[i]
			presentlyOpenFiles[of] = struct{}{}
			if _, ok := prevOpenFiles[of]; ok {
				continue
			} else {
				prevOpenFiles[of] = struct{}{}
				fmt.Printf("%s\topen\t%s\n", t.Format("2006-01-02T15:04:05.999999-07:00"), openFiles[i])

			}
		}

	}

}

func usage() {
	fmt.Println("usage")
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

func getOpenFiles(d string, c chan []string) {

	fdDir := d + "/fd"

	// Needs to be definable at command line
	ticker := time.NewTicker(time.Millisecond * 10)
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
			//info, err := os.Lstat(fdDir + "/" + f.Name())
			fullName := fdDir + "/" + f.Name()

			realFile, err := os.Readlink(fullName)
			if err != nil {
				//close(c)
				//log.Fatal(err)
				// It may no longer be open
				continue
			}
			if !strings.HasPrefix(realFile, "/") || strings.HasPrefix(realFile, "/dev/") {
				continue
			}
			openFiles = append(openFiles, realFile)
		}
		c <- openFiles
	}
	close(c)
}
