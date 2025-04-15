package channels

import "fmt"

// pong: listen the channel
func pong(wordSentCh, wordReceivedCh chan string) {
	for {
		word := <-wordSentCh
		fmt.Println("ping: " + word)
		wordReceivedCh <- word
	}
}

func main() {
	// ping-pong with channels
	wordSentCh := make(chan string)
	wordReceivedCh := make(chan string)

	go pong(wordSentCh, wordReceivedCh)

	for {
		var input string
		fmt.Print("->: ")
		_, _ = fmt.Scanln(&input)

		if input == "q" {
			break
		}

		wordSentCh <- input

		fmt.Println("pong: " + <-wordReceivedCh)
	}

	close(wordReceivedCh)
	close(wordSentCh)
}
