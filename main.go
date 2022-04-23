package main

import (
	"fmt"
	"math/rand"
	"time"
)

var JBP_QUOTES [10]string = [...]string{
	"You do not get not to make your sacrifice, you can pick which one will you make",
	"We need men",
	"There is no better predictor of life success than IQ",
	"It has been rather uncomfortable",
	"Rescue your father from the belly of the whale",
	"It is not a thrivial thing",
	"Accept responsibility",
	"Grow the hell up!",
	"Qlue in, bucko!",
	"Defeat the chaos",
}

func get_random_quote_number() int {
	seed := rand.NewSource(time.Now().UnixNano())
    random_number_generator := rand.New(seed)

	return random_number_generator.Intn(10)
}

func get_quote() string {
	return JBP_QUOTES[get_random_quote_number()]
}

func main() {
	fmt.Println("Random JBP Quote Generator")
	fmt.Println("Your quote for Today is:")
	fmt.Println(get_quote())
}