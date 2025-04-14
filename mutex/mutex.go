package mutex

import (
	"fmt"
	"sync"
)

var couter int
var mutex sync.Mutex
var wg sync.WaitGroup

func sum(wg *sync.WaitGroup, m *sync.Mutex, i int) {
	defer wg.Done()

	m.Lock() // antes da escrita e da leitura de uma variavel.
	couter += i
	fmt.Printf("Counter received : %d and now is: %d\n", i, couter)
	m.Unlock()
}

func Run() {
	couter = 0
	fmt.Println("Counter initial value: ", couter)

	size := 10
	wg.Add(size)
	for i := 1; i <= size; i++ {
		go sum(&wg, &mutex, i)
	}

	wg.Wait()

	fmt.Println("Counter value: ", couter)
}
