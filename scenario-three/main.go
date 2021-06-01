// Single Consumer Multiple Producer
package scenariothree

import (
	"fmt"
	"sync"
)

const producerCount int = 4

var messages = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

func produce(jobs chan<- string, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, msg := range messages[idx] {
		fmt.Printf("Producer %v sending message \"%v\"\n", idx, msg)
		jobs <- msg
	}
}

func consume(jobs <-chan string, done chan<- bool) {
	for msg := range jobs {
		fmt.Printf("Consumed message \"%v\"\n", msg)
	}
	done <- true
}

func main() {
	jobs := make(chan string)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	for i := 0; i < producerCount; i++ {
		wg.Add(1)
		go produce(jobs, i, &wg)
	}

	go consume(jobs, done)

	wg.Wait()
	close(jobs)
	<-done
}
