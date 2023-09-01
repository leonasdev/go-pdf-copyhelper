package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"golang.design/x/clipboard"
)

func listenWayland() error {
	var toWrite []byte
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()
	for range ticker.C {
		input, err := exec.Command("wl-paste", "-n").Output()
		if err != nil {
			return err
		}
		toWrite = bytes.ReplaceAll(input, []byte("\n"), []byte(" "))
		err = exec.Command("wl-copy", "-n", string(toWrite)).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func listen() error {
	err := clipboard.Init()
	if err != nil {
		return err
	}

	changed := clipboard.Watch(context.TODO(), clipboard.FmtText)

	var toWrite []byte

	for input := range changed {
		if bytes.Equal(input, toWrite) {
			fmt.Printf("ToWrite string: %v\n", string(toWrite))
		} else {
			fmt.Printf("Copied string: %v, %v\n", string(input), string(toWrite))
		}
		toWrite = bytes.ReplaceAll(input, []byte("\r\n"), []byte(" "))
		clipboard.Write(clipboard.FmtText, []byte(toWrite))
	}

	return nil
}

func main() {
	fmt.Println("Start listening to clipboard...")
	if os.Getenv("WAYLAND_DISPLAY") != "" {
		err := listenWayland()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	err := listen()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
