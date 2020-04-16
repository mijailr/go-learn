package main

import (
	"fmt"
	"github.com/mijailr/go-learn/pkg/database"
	"github.com/mijailr/go-learn/pkg/server"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	d := database.Connect()
	defer d.Close()

	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}


