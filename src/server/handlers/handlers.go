package handlers

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	pb "github.com/ziemowit141/random-jbp-quote-generator/src/quotes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func get_quote() string {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewQuotterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetQuote(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not quote: %v", err)
	}
	return r.GetQuote()
}

type NoMotivationHandler struct {}

func (h NoMotivationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf(`<h1> Random JBP Quote Generator</h1>
							 <h2> Your quote for Today is:</h2>
							 <h3> %s </h3>
							 <button onClick="window.location.reload();">New Quote</button>`, get_quote())

	io.WriteString(w, response)
}