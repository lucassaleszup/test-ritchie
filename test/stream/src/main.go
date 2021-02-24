// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"fmt"
	"formula/pkg/formula"
	"os"
)

func main() {
	streamTime := os.Getenv("RIT_INPUT_STREAM_TIME")

	fmt.Println("Stream started every 1s a msg is created.")

	formula.Formula{
		StreamTime: streamTime,
	}.Run()
}
