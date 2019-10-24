package main

import (
	"fmt"

	"github.com/my-ui/edgex"
)

func main() {
	config, err := edgex.LoadConfig("")
	if err != nil {
		return
	}
	fmt.Println(config)

}
