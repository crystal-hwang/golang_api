package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	// "os"
)

type Words struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Device     string `json:"device"`
	BodyPart   string `json:"bodyPart"`
	Mood       string `json:"mood"`
	Action     string `json:"action"`
}

const url = "http://localhost:8080/madlib"

func main() {
	// Define the API endpoint

	resp, err := http.Get(url)

	// fmt.Println(os.Args[1])

	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		var words Words

		err = json.Unmarshal(body, &words)
		if err != nil {
			log.Fatalf("Error unmarshalling response body: %v", err)
		}
		// this prints out raw json text

		// fmt.Println(string(body))

		fmt.Println(words)
	}

}
