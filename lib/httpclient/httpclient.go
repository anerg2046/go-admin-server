package httpclient

import (
	"go-app/config"
	"net/http"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	client *resty.Client
	once   sync.Once
)

func C() *resty.Client {
	once.Do(func() {
		client = resty.New()
		setup()
	})
	return client
}

func setup() {
	client.SetDebug(config.HTTPCLIENT.Debug)
	client.SetRetryCount(config.HTTPCLIENT.RetryCount)
	client.SetTimeout(time.Duration(config.HTTPCLIENT.Timeout) * time.Second)
	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			if r.StatusCode() != http.StatusOK || r == nil {
				return true
			}
			return false
		},
	)
}
