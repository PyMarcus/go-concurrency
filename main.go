package main

import (
	"github.com/PyMarcus/go-concurrency/mutex"
	waitgroups "github.com/PyMarcus/go-concurrency/wait_groups"
)

func main() {
	waitgroups.Run()
	mutex.Run()
}
