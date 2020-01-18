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
	//tracking variables for go routine
	numOfCompletedRoutine := 0

	gihubUserName := []string{
		"SamsadSajid",
		"google",
	}

	start := time.Now()

	for _, name := range gihubUserName {
		//
		/**
		* Following is a self executing function which is a g routine
		* It takes a string param symbol. We can't use the `name` from
		* the `for` loop because the following is a concurrent function
		* to main function. Main func will continue to execute the for loop
		* not knwoing about the following function. If we kept the `name` variable
		* it might so happen that after all the iterations have been done, the concurrency
		* has begun execution and right then the `name` variable points to only the
		* `SamsadSajid` from the above list which we definitely don't want. We want
		* for each of the name to make a network call!
		* So each goroutines are decoupled from for loop
		 */

		go func(symbol string) {
			fmt.Println(symbol)
			resp, _ := http.Get("https://api.github.com/users/" + symbol)
			//close the connection
			defer resp.Body.Close()

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

			numOfCompletedRoutine++

			fmt.Println("numOfCompletedRoutine:=", numOfCompletedRoutine)
		}(name)
	}
	/*
	* As long as numOfCompletedRoutine is not equal to the
	* len of the gihubUserName list, make the main function
	* sleep for 10 Millisecond
	 */
	for numOfCompletedRoutine < len(gihubUserName) {
		time.Sleep(10 * time.Millisecond)
	}

	timeTakenForCall := time.Since(start)
	fmt.Println("Execution time is: ", timeTakenForCall)
}

type GithubUserInfo struct {
	Login     string
	Followers int
	Email     string
	Bio       string
}
