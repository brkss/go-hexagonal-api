package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/brkss/dextrace-server/internal/adapter/config"
)




func main() {

	// load envirenement variables 
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environement variables", "error", err)
		os.Exit(1);
	}

	fmt.Printf("%v : ", config.App);

	//fmt.Println("Hello, World!")
}