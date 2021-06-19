package main

import (
	"fmt"
	"github/ByungHaLee/nomadcoin/blockchain"
	"log"
	"net/http"
	"text/template"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)

}
func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
