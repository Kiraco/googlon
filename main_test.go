package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		main()
		wg.Done()
	}(wg)
}
func TestMain(t *testing.T) {
	cases := []struct {
		Input  string
		Output string
	}{
		{Input: TestCaseA, Output: ExpectedOutputA},
		{Input: TestCaseB, Output: ExpectedOutputB},
		{Input: TestCaseC, Output: ExpectedOutputC},
		{Input: TestCaseD, Output: ExpectedOutputD},
		{Input: TestCaseE, Output: ExpectedOutputE},
	}

	for _, singleCase := range cases {

		_, err := http.Get("http://localhost:8080/ping")
		assert.Nil(t, err)

		response, err := http.Post("http://localhost:8080/process", "text", strings.NewReader(singleCase.Input))
		assert.Nil(t, err)

		bytes, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		assert.Equal(t, singleCase.Output, string(bytes))
	}
}
