package deps

import (
	"log"
	"sync"
	"time"
)

type Email struct {
	To, Subject, Content string
}

type EmailServer struct {
	queue                []Email
	mu                   sync.Mutex
	MaxConcurrentWorkers int
	ticker               *time.Ticker
}

func (es *EmailServer) sendEmailWorker(id int, jobsCh <-chan Email) {
	for range jobsCh {
		time.Sleep(300 * time.Millisecond)
		log.Printf("Email Sent by worker #%d!\n", id)
	}
}

func (es *EmailServer) SendEmail(email Email) {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.queue = append(es.queue, email)
}

func NewEmailServer(MaxConcurrentWorkers int, ticker *time.Ticker) *EmailServer {
	return &EmailServer{queue: []Email{}, MaxConcurrentWorkers: MaxConcurrentWorkers, ticker: ticker}
}

func (es *EmailServer) RunPostOffice() {

	for {
		<-es.ticker.C
		es.mu.Lock()
		jobsCh := make(chan Email, es.MaxConcurrentWorkers)
		numWorkers := min(es.MaxConcurrentWorkers, len(es.queue))
		for workerID := 1; workerID <= numWorkers; workerID++ {
			go es.sendEmailWorker(workerID, jobsCh)
		}
		for _, email := range es.queue {
			jobsCh <- email
		}
		es.queue = []Email{}
		close(jobsCh)
		es.mu.Unlock()
	}
}
