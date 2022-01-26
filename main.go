package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from home !")
}

func main() {
	// chain := blockchain.GetBlockChain()
	// chain.AddBlock("Second Block")
	// for _, block := range chain.GetAllBlocks() {
	// 	fmt.Printf("%s\n", block)
	// }
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}
