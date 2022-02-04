package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/estellechoi/sunkisscoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "explorer/templates"
)

var templates *template.Template

type homeData struct {
	Title  string
	Blocks []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml")) // template.Must does ... if err != nil { log.Fatal(err) }
	data := homeData{Title: "Home", Blocks: blockchain.GetBlockChain().GetAllBlocks()}
	// tmpl.Execute(rw, data)
	templates.ExecuteTemplate(rw, "home", data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := homeData{Title: "Add Blocks!", Blocks: blockchain.GetBlockChain().GetAllBlocks()}
		templates.ExecuteTemplate(rw, "blocks", data)
	case "POST":
		r.ParseForm()
		blockData := r.Form.Get("blockData")
		blockchain.GetBlockChain().AddBlock(blockData)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start() {
	// chain := blockchain.GetBlockChain()
	// chain.AddBlock("Second Block")
	// for _, block := range chain.GetAllBlocks() {
	// 	fmt.Printf("%s\n", block)
	// }
	templates = template.Must(template.ParseGlob(templateDir + "/pages/*.gohtml"))     // **/* does not work in Go
	templates = template.Must(templates.ParseGlob(templateDir + "/partials/*.gohtml")) // see templates.ParseGlob, not template.ParseGlob

	// routes
	http.HandleFunc("/", home)
	http.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}
