package main

import (
	"bytes"
	"context"
	"fmt"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
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
}
