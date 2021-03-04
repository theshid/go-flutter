package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const port = ":5550"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/products/{fetchCountPercentage}", products).Methods("GET")

	fmt.Println("Serving @ http:127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is root page"))
}

func products(w http.ResponseWriter, r *http.Request) {
fetchCountPercentage,errInput := strconv.ParseFloat(mux.Vars(r)["fetchCountPercentage"],64)

fetchCount := 0
if errInput != nil{
	fmt.Println(errInput.Error())
	}else{
		fetchCount = int(float64(len(productList)) * fetchCountPercentage /100)
		if fetchCount > len(productList){fetchCount = len(productList)}
	}

jsonList,err := json.Marshal(productList[0:fetchCount])
if err != nil {
	http.Error(w,err.Error(),http.StatusInternalServerError)
	} else{
		w.Header().Set("content-type","application/json")
		w.Write(jsonList)
	}
}

type product struct {
	Name string
	Price float64
	Count int
}

var productList = []product{
	{"p1",25.0,30},
	{"p2",45.0,50},
	{"p3",95.0,90},
	{"p4",35.0,80},
	{"p5",55.0,50},
	{"p6",95.0,40},
	{"p7",85.0,70},
	{"p8",75.0,20},
	{"p9",35.0,10},
	{"p10",5.0,70},
	{"p11",15.0,830},
}
