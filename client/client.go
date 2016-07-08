package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tracer/tracer"
)

type spanResponse struct {
	Error string         `json:"error"`
	Span  tracer.RawSpan `json:"span"`
}

type QueryClient struct {
	host   string
	client *http.Client
}

func NewQueryClient(host string) *QueryClient {
	return &QueryClient{
		host:   host,
		client: &http.Client{},
	}
}

func (q *QueryClient) SpanByID(id uint64) (tracer.RawSpan, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/span/?id=%016x", q.host, id), nil)
	if err != nil {
		panic(err)
	}

	resp, err := q.client.Do(req)
	if err != nil {
		return tracer.RawSpan{}, err
	}
	defer resp.Body.Close()
	//var sr spanResponse
	var sr tracer.RawSpan
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return tracer.RawSpan{}, err
	}
	// if sr.Error != "" {
	// 	return tracer.RawSpan{}, errors.New(sr.Error)
	// }
	// return sr.Span, nil
	return sr, nil
}

func (q *QueryClient) TraceByID(id uint64) (tracer.RawTrace, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/trace/?id=%016x", q.host, id), nil)
	if err != nil {
		panic(err)
	}

	resp, err := q.client.Do(req)
	if err != nil {
		return tracer.RawTrace{}, err
	}
	defer resp.Body.Close()
	//var sr spanResponse
	var tr tracer.RawTrace
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return tracer.RawTrace{}, err
	}
	// if sr.Error != "" {
	// 	return tracer.RawSpan{}, errors.New(sr.Error)
	// }
	// return sr.Span, nil
	return tr, nil
}
