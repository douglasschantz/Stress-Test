package main

import (
	"gitlab.com/douglasschantz/stress-test/internal/infrastructure"
)

func main() {
	container := infrastructure.NewContainer()
	cli := container.GetCLI()
	cli.Execute()
}
