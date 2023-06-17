package main

import (
	"fmt"
	"github.com/Sidharth-Shankar-Bihary/go-http-client/gohttp"
	"net/http"
)

var githubHttpClient = getGithubClient()

type User struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	getRequestAPIGithub()
	user := &User{
		Firstname: "John",
		LastName:  "Travolta",
	}
	createUser(*user)
}

func getGithubClient() gohttp.HttpClient {
	client := gohttp.NewClient()
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(headers)

	return client
}

func getRequestAPIGithub() {
	getRequestHeaders := make(http.Header)
	getRequestHeaders.Set("Accept", "Application/xml")
	response, err := githubHttpClient.Get("https://api.github.com", getRequestHeaders)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
}

func createUser(user User) {
	postRequestHeaders := make(http.Header)
	postRequestHeaders.Set("Accept", "Application/xml")
	response, err := githubHttpClient.Post("https://api.github.com", postRequestHeaders, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
}
