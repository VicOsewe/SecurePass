package main

import (
	"fmt"

	"github.com/VicOsewe/secure-pass/pkg"
	"github.com/VicOsewe/secure-pass/pkg/config"
)

func main() {
	pkg.SetUpRouter()
	env := config.NewEnv()
	fmt.Println("env:", env)
}
