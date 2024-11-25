package usecase_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/jhonathann10/stress-test/internal/usecase"
	"github.com/jhonathann10/stress-test/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestStartRequests_Execute(t *testing.T) {
	clientInterface := new(mocks.ClientInterface)
	clientInterface.On("Get").Return(func() (int, error) {
		statuses := []int{200, 404, 429, 500}
		rand.Seed(time.Now().UnixNano())
		return statuses[rand.Intn(len(statuses))], nil
	})
	requests := usecase.NewStartRequests(100000, 1000, clientInterface)

	err := requests.Execute()

	assert.Nil(t, err)
}
