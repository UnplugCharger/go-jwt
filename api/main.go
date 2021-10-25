package main 

import(
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
)

var MySigningKey= []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter , r *http.Request)  {
	fmt.Fprintf(w,"This is a very secrete info")
}

func isAuthorised(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Header["Token"] != nil {
			fmt.Println("Mahinya")
		}
		fmt.Println("Byron")
	})


func handleRequests ()  {
	http.Handle("/",isAuthorised(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}





func main()  {
	 fmt.Println("server")
	 handleRequests()
}