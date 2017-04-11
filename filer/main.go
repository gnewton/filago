package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println(os.Getpid())

	for {
		go func() {
			cmd := exec.Command("sleep", "40")
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Waiting for command to finish...")
			err = cmd.Wait()
			log.Printf("Command finished with error: %v", err)
		}()
		tmpfile, err := ioutil.TempFile("", "example")
		//fmt.Println(tmpfile.Name())
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)
		err = tmpfile.Close()
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

	}
}
