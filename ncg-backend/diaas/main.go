package main

import (
	. "DIaaS/diaas/api"
	. "DIaaS/diaas/constants"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(PORT, NewServer())
	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
