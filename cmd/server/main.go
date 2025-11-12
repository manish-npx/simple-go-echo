package main

import (
	"fmt"

	"github.com/manish-npx/simple-go-echo/internal/config"
)

func main() {

	//config done
	cfg := config.LoadConfig()

	fmt.Printf("config value is %v", cfg)

	//db connection

	//routes
}
