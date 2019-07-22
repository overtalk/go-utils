package pprof_test

import (
	"net/http"
	"testing"

	"github.com/rs/cors"

	"github.com/qinhan-shu/go-utils/pprof"
)

func TestAddPprof(t *testing.T) {
	mux := http.NewServeMux()
	pprof.AddPprof(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	server := http.Server{
		Addr:    ":9091",
		Handler: c.Handler(mux),
	}

	if err := server.ListenAndServe(); err != nil {
		t.Fatal(err)
	}
}
