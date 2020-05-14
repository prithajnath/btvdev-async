package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type factorResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

type factorResult struct {
	Result []int `json:"result"`
}

var wg sync.WaitGroup

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	if n < 0 {
		panic("N cannot be negative")
	}

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int, wg *sync.WaitGroup) {
			url := "https://0bdwnj5rj7.execute-api.us-east-1.amazonaws.com/prod"
			jsonStr := []byte(fmt.Sprintf(`{"number": %d}`, i))

			req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Connection", "close")

			if err != nil {
				panic(err)
			}

			client := http.Client{Timeout: time.Second * 2}
			req.Close = true
			resp, err := client.Do(req)

			if err != nil {
				panic(err)
			}

			var factorResp factorResponse

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			jsonErr := json.Unmarshal(body, &factorResp)

			if jsonErr != nil {
				panic(err)
			}

			var result factorResult

			factorResultErr := json.Unmarshal([]byte(factorResp.Body), &result)
			if factorResultErr != nil {
				panic(factorResultErr)
			}

			fmt.Printf("%d can be factored into %v\n", i, result.Result)
			resp.Body.Close()

			wg.Done()
		}(rand.Intn(10000-5000+1), &wg)
	}

	wg.Wait()
}
