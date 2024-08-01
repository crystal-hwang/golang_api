package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type api_response struct {
    Message string `json:"message"`
}

// // get request to api endpoint
// func getRequest(url string) (*ApiResponse, error) {
//     resp, err := http.Get(url)
//     if err != nil {
//         return nil, err
//     }
//     defer resp.Body.Close()

//     body, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         return nil, err
//     }

//     var apiResponse ApiResponse
//     err = json.Unmarshal(body, &apiResponse)
//     if err != nil {
//         return nil, err
//     }

//     return &apiResponse, nil
// }

// server than returns json in response to requests
// allows us to see what happens when we get json from a server 
const url = "https://jsonplaceholder.typicode.com"

func main() {
    // Define the API endpoint

    resp, err := http.Get("http://localhost:8080/" + os.Args[1])

    fmt.Println(os.Args[1])


    if err != nil {
        log.Fatalf("Error making GET request: %v", err)
    }

    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        body, err := ioutil.ReadAll(resp.Body)

        if err != nil {
            log.Fatalf("Error reading response body: %v", err)

        }

        fmt.Println(string(body))
    }

}
