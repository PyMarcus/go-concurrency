package waitgroups

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
	fmt.Println(show())
}

func show() string {
	return msg
}

func Run() {
	var wg sync.WaitGroup

	msgs := []string{
		"text1",
		"text2",
		"text3",
		"text4",
		"text5",
	}

	wg.Add(len(msgs))
	for _, msg := range msgs {
		go updateMessage(msg, &wg)
	}

	wg.Wait()
}
