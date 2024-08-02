package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// type Joke struct {
// 	Joke string `json:"joke"`
// }

// type Words struct {
// 	Name       string `json:"name"`
// 	Age        int    `json:"age"`
// 	Occupation string `json:"occupation"`
// 	Device     string `json:"device"`
// 	BodyPart   string `json:"body_part"`
// 	Mood       string `json:"mood"`
// 	Action     string `json:"action"`
// }

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
	jokes := []string{
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

	// response := Joke{Joke: jokes[rand_i]}
	w.Header().Set("Content-Type", "text/plain")

	// Print the random joke
	fmt.Fprintf(w, jokes[rand_i])

}

func madlib_handler(w http.ResponseWriter, r *http.Request) {

	names := []string{"Alice", "Bob", "Charlie", "Diana", "Edward", "Fiona", "George", "Hannah", "Ian", "Jasmine"}
	ages := []int{25, 32, 45, 29, 51, 38, 22, 41, 27, 34}
	occupations := []string{"Software Engineer", "Teacher", "Doctor", "Nurse", "Artist", "Chef", "Mechanic", "Architect", "Plumber", "Musician"}
	devices := []string{"Laptop", "Smartphone", "Tablet", "Desktop", "Server", "Smartwatch", "Router", "Printer", "External Hard Drive", "Monitor"}
	body_parts := []string{"Head", "Arm", "Leg", "Hand", "Foot", "Eye", "Ear", "Nose", "Mouth", "Back"}
	moods := []string{"Happy", "Sad", "Excited", "Angry", "Bored", "Anxious", "Calm", "Nervous", "Content", "Irritated"}
	actions := []string{"Running", "Jumping", "Reading", "Writing", "Cooking", "Sleeping", "Driving", "Swimming", "Dancing", "Talking"}

	rand_name := rand.Intn(len(names))
	rand_age := rand.Intn(len(ages))
	rand_job := rand.Intn(len(occupations))
	rand_device := rand.Intn(len(devices))
	rand_bp := rand.Intn(len(body_parts))
	rand_mood := rand.Intn(len(moods))
	rand_actions := rand.Intn(len(actions))

	// words := Words{
	// 	Name: names[rand_name],
	// 	Age: ages[rand_age],
	// 	Occupation: occupations[rand_job],
	// 	Device: devices[rand_device],
	// 	BodyPart: body_parts[rand_bp],
	// 	Mood: moods[rand_mood],
	// 	Action: actions[rand_actions],
	// }

    paragraph := fmt.Sprintf(
        "%s is a %d-year-old %s who has been struggling with a lot of job-related stress. He/she decided to try a new application to relieve stress, which runs on a/an %s to help improve his/her mood.\n\nThe application senses his/her mood through a device he/she wears on his/her %s.\n\nWhen the device senses that he/she is %s, it responds by %s.", names[rand_name], ages[rand_age], occupations[rand_job], devices[rand_device], body_parts[rand_bp], moods[rand_mood], actions[rand_actions])

	w.Header().Set("Content-Type", "application/json")

	// writes to http response body which is sent bck to http
	fmt.Fprintln(w, paragraph)

	// Encoder that writes a response
	// encoder := json.NewEncoder(w)

	// if err := encoder.Encode(words); err != nil {
	// 	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	// 	return
	// }

}
