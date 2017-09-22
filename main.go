package main

import(
	"net/http"
	"log"

	"github.com/berryhill/gok/service"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := service.StringService{}

	uppercaseHandler := httptransport.NewServer(
		service.MakeUppercaseEndpoint(svc),
		service.DecodeUppercaseRequest,
		service.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)

	log.Fatal(http.ListenAndServe(":8008", nil))
}
