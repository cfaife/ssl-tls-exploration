package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	Black      = "\033[30m"
    Reset      = "\033[0m"
    Red        = "\033[31m"
    Green      = "\033[32m"
    Yellow     = "\033[33m"
    Blue       = "\033[34m"
    Magenta    = "\033[35m"
    Cyan       = "\033[36m"
    White      = "\033[37m"
	BrightBlack   = "\033[90m"
    BrightRed  = "\033[91m"
    BrightGreen = "\033[92m"
    BrightYellow = "\033[93m"
    BrightBlue  = "\033[94m"
    BrightMagenta = "\033[95m"
    BrightCyan  = "\033[96m"
    BrightWhite = "\033[97m"
)


func main(){

	loggerInfo := log.New(os.Stdout,Magenta+"INFO ",log.Ldate|log.Ltime)

	http.HandleFunc("/",handle)
	loggerInfo.Println("Server is listening")

	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		loggerRrro:= log.New(os.Stdout,Red+"INFO ",log.Ldate|log.Ltime)

		loggerRrro.Println(err)
		panic(err)
	}
}


func  handle(response http.ResponseWriter, request *http.Request)  {
	if request.Method == http.MethodGet {
		response.Header().Set("Content-type","application/json")
		response.WriteHeader(http.StatusOK)

		err:= json.NewEncoder(response).Encode("{'status':'up'}")
		if err != nil{
			panic(err)
		}
	}
}