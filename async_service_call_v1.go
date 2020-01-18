package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//to make web requests
	"net/http"

	//read the response from call

	//unmarshall the XML response

	"time"
)

func main() {
	gihubUserName := []string{
		"SamsadSajid",
		"go",
		"elixir",
		"google",
	}

	start := time.Now()

	for _, name := range gihubUserName {
		resp, _ := http.Get("https://api.github.com/users/" + name)
		//close the connection
		defer resp.Body.Close()

		// fmt.Println(resp)

		// this should be done carefully because
		// if data is big it can blow up the system but
		// in this example data is small so we can read that in single byte stream
		body, _ := ioutil.ReadAll(resp.Body)

		gihubUserInfo := GithubUserInfo{}
		json.Unmarshal(body, &gihubUserInfo)

		fmt.Println(
			"Login Name: = ", gihubUserInfo.Login,
			"Followers:=", gihubUserInfo.Followers,
			"Email:=", gihubUserInfo.Email,
			"Bio:=", gihubUserInfo.Bio,
		)

		timeTakenForCall := time.Since(start)
		fmt.Println("Execution time is: ", timeTakenForCall)
	}
}

type GithubUserInfo struct {
	Login     string
	Followers int
	Email     string
	Bio       string
}
