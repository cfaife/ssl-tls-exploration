package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func main(){

	http.HandleFunc("/",func(response http.ResponseWriter, request * http.Request){

		if request.Method == http.MethodGet {
			
			response.Header().Set("Content-type","application/json")
			response.WriteHeader(http.StatusOK)

			json.NewEncoder(response).Encode("{status:alive}")
		}

	})

	fmt.Println("The server is listening in port 8443")

	err := http.ListenAndServeTLS(
					":8443",
					"../certificates/certificate.pem",
					"../certificates/private_key.pem",
					nil)

	if err != nil {
		panic(err)
	}

}