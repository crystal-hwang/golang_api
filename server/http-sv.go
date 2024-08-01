package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"time"
)

type Joke struct {
	Joke string `json:"joke"`
}

func main() {
	// Define routes
	http.HandleFunc("/joke", joke_handler)
	http.HandleFunc("/madlib", madlib_handler)

	// Start the server and start on port 8080
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func joke_handler(w http.ResponseWriter, r *http.Request) {
	// Create slice of jokes
	jokes := []string {
		"Why did the tomato blush? Because it saw the salad dressing.",
		"What do you call a bear with no teeth? A gummy bear.",
		"Why did the physicist go to the beach? To catch some waves!",
		"How does a moon cut his hair? Eclipse it!",
		"Why was the big cat disqualified from the race? Because it was a cheetah!",
		"Why did the chicken join a band? Because it had the drumsticks!",
		"How does a mathematician plow fields? With a pro-tractor.",
		"What do you get if you cross a snowman and a vampire? Frostbite.",
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

    // Get a random index
    rand_i := rand.Intn(len(jokes))

    // Print the random joke
    fmt.Println(jokes[rand_i])

	response := Joke{Joke: jokes[rand_i]}
	w.Header().Set("Content-Type", "application/json")

	// Encoder that writes a response
	encoder := json.NewEncoder(w)

	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)

	}	


}

func madlib_handler(w http.ResponseWriter, r *http.Request) {

}
