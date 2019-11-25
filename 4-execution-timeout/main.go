package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func NewRemoteApiClient(timeout time.Duration) *remoteApiClient {
	return &remoteApiClient{
		timeoutOption: timeout,
		random:        rand.New(rand.NewSource(time.Now().Unix())),
	}
}

type remoteApiClient struct {
	timeoutOption time.Duration
	random        *rand.Rand
}

func (rac *remoteApiClient) GetLastVersion() (string, error) {
	return "", nil
}

func (rac *remoteApiClient) GetCategoriesList() ([]string, error) {

	results := make(chan []byte, 0)

	go func() {
		// call remote API
		
		randPart := rac.random.Intn(5000)
		duration := time.Duration(2500 + randPart)
		fmt.Println("DEBUG:", "will wait for", duration*time.Millisecond)
		time.Sleep(duration * time.Millisecond)

		results <- []byte(`["hello", "world"]`)
	}()

	select {
	case <-time.After(rac.timeoutOption):
		return nil, errors.New("timeout exceeded")
	case answer := <-results:
		var responseBody []string
		_ = json.Unmarshal(answer, &responseBody)
		return responseBody, nil
	}

}

func main() {
	client := NewRemoteApiClient(5 * time.Second)
	list, err := client.GetCategoriesList()

	if err != nil {
		panic(err)
	}

	for _, str := range list {
		fmt.Println("Category:", str)
	}
}
