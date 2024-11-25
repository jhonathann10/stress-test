package entity

import "time"

type Report struct {
	TimeDuration        time.Duration
	TotalRequests       int
	TotalRequestsOK     int
	TotalRequestsErrors map[int]int
}

func (r *Report) SumTotalRequests() {
	r.TotalRequests++
}

func (r *Report) SumTotalRequestsOK() {
	r.TotalRequestsOK++
}

func (r *Report) MappingStatusErrors(status int) {
	r.TotalRequestsErrors[status]++
}
