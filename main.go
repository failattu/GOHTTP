package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"log"
	"html"
	"io"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)
func server(){
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/mysqlw", func(w http.ResponseWriter, r *http.Request) {
		writeSQL()
	})
	http.HandleFunc("/mysqlr", func(w http.ResponseWriter, r *http.Request) {
		readSQL()

	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
func readSQL(){
	con, err := sql.Open("mysql", "test:test@/test")
	if err != nil { }
	var (
	    PersonID int
	    LastName string
			FirstName string
			Address string
			City string
	)
	rows, err := con.Query("SELECT PersonID, LastName , FirstName, Address, City FROM Persons")
	if err != nil {
	    log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
	    err := rows.Scan(&PersonID, &LastName,&FirstName,&Address,&City)
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(PersonID, LastName,FirstName,Address,City)
			fmt.Print(PersonID, LastName,FirstName,Address,City)
	}
	err = rows.Err()
	if err != nil {
	    log.Fatal(err)
	}

	defer con.Close()
}
func writeSQL(){
	con, err := sql.Open("mysql", "test:test@/test")
	if err != nil {}
	//Persons(PersonID int,LastName varchar(255),FirstName varchar(255),Address varchar(255),City varchar(255));"
	_, err = con.Exec("INSERT INTO Persons (PersonID, LastName , FirstName, Address, City) VALUES (0, 'test','vesa','pallo', 'forssa');")
	fmt.Printf( "Kirjotusta")
	defer con.Close()
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
    var buf io.Reader
	resp, err := http.Post(address,"text/plain", buf)
	if err != nil {
		fmt.Printf("ERROR IN GET")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	fmt.Printf( s)

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
