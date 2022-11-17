package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://localhost:8617/index"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjI1NDA0MjM5NDk3MTI0MzcyNDgsImV4cCI6MTY2NTEyNTIxNywiaWF0IjoxNjY0NTIwNDE3fQ.OC-35lkt405tL-8Oy7vgi4p7Y9WF5lYUAfLhOPzXAYU")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
