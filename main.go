package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"log"
	"html"
)
func server(){
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getReq(address string){
	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("ERROR IN GET")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	fmt.Printf( s)
}

func postReq(address string){


}

func main() {
	fmt.Printf("hello, world\n")
	
	if len(os.Args) > 1 {
		if os.Args[1] == "get" {
			var address = os.Args[2]
			getReq(address)
		}else if os.Args[1] == "post"{
			var address = os.Args[2]
			postReq(address)
		}else if os.Args[1] == "server"{
			server()
		}else{
			fmt.Printf("Goodbye\n")
		}
	}
}