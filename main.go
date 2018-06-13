package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hypebeast/go-osc/osc"
)

var address string
var port int

func main() {
	f := flag.String("listen", "59000", "listen port of VRChat OSC Bridge.")
	g := flag.String("server", "127.0.0.1", "address of OSC receiver.")
	h := flag.String("port", "9000", "port of OSC receiver.")
	flag.Parse()

	address = *g
	port, _ = strconv.Atoi(*h)

	fmt.Println("Server listening at :" + *f)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+*f, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		send(r.URL.Path, body)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, "{\"vrchat-osc-bridge\":\"0.0.0\"}")
	}
}

func send(a string, d string) {
	client := osc.NewClient(address, port)
	msg := osc.NewMessage(a)
	msg.Append(d)
	client.Send(msg)
}
