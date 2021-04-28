// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"fmt"
	"formula/pkg/formula"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")

	fmt.Println("Stream started every 2s a msg is created.")

	formula.Formula{
		Token: token,
	}.Run()
}
