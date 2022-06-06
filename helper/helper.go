package main

import (
	"git.jetbrains.space/artdecoction/wt/tower/helper/contracts"
	"os"
)

func main() {
	serviceName := os.Args[1]

	contracts.GenerateContracts(serviceName)
}
