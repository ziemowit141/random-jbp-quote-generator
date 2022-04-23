package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ziemowit141/random-jbp-quote-generator/src/quotes"
)

type NoMotivationHandler struct {}

func (h NoMotivationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	inner_wisdom := make(chan string)
	go quotes.GetQuote(inner_wisdom)

	response := fmt.Sprintf(`<h1> Random JBP Quote Generator</h1>
							 <h2> Your quote for Today is:</h2>
							 <h3> %s </h3>
							 <button onClick="window.location.reload();">New Quote</button>`, <-inner_wisdom)

	io.WriteString(w, response)
}
