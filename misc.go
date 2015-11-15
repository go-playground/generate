package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func executeCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
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
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	// go generate command will fail when no generate command find.
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
	// 	os.Exit(1)
	// }
}
