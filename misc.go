package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func executeCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	}

	defer cmdReader.Close()

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal(os.Stderr, "Error starting Cmd", err)
	}

	err = cmd.Wait()
	// go generate command will fail when no generate command find.
	if err != nil {
		if err.Error() != "exit status 1" {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
}
