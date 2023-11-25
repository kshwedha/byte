package lib

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func addHeaders(req *http.Request) *http.Request {
	token := ""
	req.Header.Add("content-type", "application/json")
	req.Header.Add("token", token)
	return req
}

func client() *http.Client {
	client := &http.Client{}
	return client
}

func buildRequest(method string, url string, body *strings.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return req
}

func performRequest(req *http.Request) *http.Response {
	client := client()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer res.Body.Close()
	return res
}

func readResponse(res *http.Response) []byte {
	_body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return _body
}

func Request(method string, url string, body *strings.Reader) []byte {
	// body can be *string.Reader or io.Reader
	// io.Reader payload ex
	// payload := strings.NewReader(`{"sort":[{"field":"createdTime","order":"DESC"}],"searchText":"","pageInfo":{"start":0,"rows":50}}`)
	request := buildRequest(method, url, body)
	request = addHeaders(request)

	res := performRequest(request)
	_body := readResponse(res)
	return _body
}

func processRequest(body []byte) []byte {
	// Process the request body
	return body
}

func Respond(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Process the request and generate a response
	response := processRequest(body)

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
