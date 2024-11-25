package usecase

import (
	"fmt"
	"sync"
	"time"

	"github.com/jhonathann10/stress-test/internal/entity"
	"github.com/jhonathann10/stress-test/internal/infra/client"
)

var (
	report = &entity.Report{
		TotalRequests:       0,
		TotalRequestsOK:     0,
		TotalRequestsErrors: make(map[int]int),
	}
	mu sync.Mutex
)

type StartRequests struct {
	requests    int
	concurrency int

	clientInterface client.ClientInterface
}

func NewStartRequests(requests, concurrency int, clientInterface client.ClientInterface) *StartRequests {
	return &StartRequests{
		requests:        requests,
		concurrency:     concurrency,
		clientInterface: clientInterface,
	}
}

func (s *StartRequests) Execute() error {
	var wg sync.WaitGroup
	concurrencyCh := make(chan struct{}, s.concurrency)

	start := time.Now()
	for i := 0; i < s.requests; i++ {
		wg.Add(1)
		concurrencyCh <- struct{}{}
		go s.StartRequests(&wg, concurrencyCh)
	}
	wg.Wait()

	report.TimeDuration = time.Since(start)

	fmt.Println("Total Requests: ", report.TotalRequests)
	fmt.Println("Total Requests [200]: ", report.TotalRequestsOK)
	if len(report.TotalRequestsErrors) > 0 {
		for key, value := range report.TotalRequestsErrors {
			fmt.Printf("Total Requests Error [%d]: %d\n", key, value)
		}
	} else {
		fmt.Println("Total Requests Errors: ", 0)
	}
	fmt.Println("Time Duration: ", report.TimeDuration)

	return nil
}

func (s *StartRequests) StartRequests(wg *sync.WaitGroup, concurrencyCh chan struct{}) {
	defer wg.Done()
	defer func() {
		<-concurrencyCh
	}()

	mu.Lock()
	report.SumTotalRequests()
	mu.Unlock()

	status, err := s.clientInterface.Get()
	if err != nil {
		return
	}

	s.countRequestsStatus(status, report)
}

func (s *StartRequests) countRequestsStatus(status int, report *entity.Report) {
	mu.Lock()
	defer mu.Unlock()

	if status == 200 {
		report.SumTotalRequestsOK()
		return
	}

	report.MappingStatusErrors(status)
}
