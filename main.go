package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Person struct{
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request){
   params:=mux.Vars(req)
   for _,item:=range people{
   	if item.ID == params["id"]{
   		json.NewEncoder(w).Encode(item)
   		return
	}
   }
	json.NewEncoder(w).Encode(&Person{})

}
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request){
  json.NewEncoder(w).Encode(people)

}
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request){
  params:=mux.Vars(req)
  var person Person
  _ = json.NewDecoder(req.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people,person)
  json.NewEncoder(w).Encode(people)

}
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request){

	params:=mux.Vars(req)

	for i,v:=range people{
		if v.ID == params["id"]{
			people = append(people[:i],people[i+1:]...)

			break
		}
	}
     json.NewEncoder(w).Encode(people)
}
func main(){

	router:=mux.NewRouter()

	people = append(people,Person{"1","Shubham","Garg",&Address{"","Delhi"}})
	people = append(people,Person{"2","Aman","Patel",&Address{"Noida","UP"}})
	people = append(people,Person{"3","Kiran","Kannade",&Address{"Pune","Maharashtra"}})

	router.HandleFunc("/people",GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}",GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}",CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}",DeletePersonEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",router))

}
