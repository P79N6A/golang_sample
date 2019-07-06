package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// simple
	{
		resp, err := http.Get("http://127.0.0.1:8080/hello/handler?name=lxb31")
		check(err)
		defer resp.Body.Close()
		readBody(resp)
	}

	// timeout
	{
		client := &http.Client{
			Timeout: 1 * time.Second,
		}

		resp, err := client.Get("http://127.0.0.1:8080/hello/handler?name=lxb31")
		check(err)
		defer resp.Body.Close()
		readBody(resp)

		req, err := http.NewRequest("GET", "http://127.0.0.1:8080/hello/handler", nil)
		check(err)
		req.Header.Add("Accept-Encoding", "gzip, deflate")
		resp, err = client.Do(req)
		check(err)
		defer resp.Body.Close()
		readBody(resp)
	}

}

func readBody(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Body: %v\n", string(body))
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
