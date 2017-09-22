package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"fmt"

)


type Employees struct {
	Name string `json:"name,omitempty"`
	ID string `json:"id,omitempty"`
	Position string `json:"position,omitempty"`
	Address *Location `json:"location,omitempty"`

}

type Location struct {
	State string `json:"state,omitempty"`
	City string `json:"city,omitempty"`
	PinCode string `json:"pincode,omitempty"`

}

var resources []Employees

func GetEmployeesDetails(w http.ResponseWriter,req *http.Request){

	json.NewEncoder(w).Encode(resources)

}
func GetEmployeeDetails(w http.ResponseWriter,req *http.Request){
  params:=mux.Vars(req)
  fmt.Println(params)
  flag:=0
  for _,item:=range resources{
  	if  item.Name == params["name"] && item.ID == params["id"]{
        json.NewEncoder(w).Encode(item)
           flag=1
	}
  }
  if flag==0 {
	  json.NewEncoder(w).Encode(&Employees{})
  }
}
func GetEmployeeDetailsPosition(w http.ResponseWriter,req *http.Request){
	params:=mux.Vars(req)
          flag:=0
	for _,item:=range resources{
		if item.Position == params["position"]{
			json.NewEncoder(w).Encode(item)
			flag=1
		}
	}
	if flag==0 {
		json.NewEncoder(w).Encode(&Employees{})
	}
}
func GetEmployeeDetailsByName(w http.ResponseWriter,req *http.Request){
	params:=mux.Vars(req)
	flag:=0
	for _,v:=range resources{
		if v.Name == params["name"]{
			json.NewEncoder(w).Encode(v)
			flag=1
		}
	}
	if flag ==0 {
		json.NewEncoder(w).Encode(&Employees{})
	}
}
func GetEmployeeDetailsByID(w http.ResponseWriter,req *http.Request){
	params:=mux.Vars(req)
	flag:=0
	for _,v:=range resources{
		if v.ID == params["id"]{
			json.NewEncoder(w).Encode(v)
			flag=1
		}
	}
	if flag ==0 {
		json.NewEncoder(w).Encode(&Employees{})
	}
}

func CreateEmployee(w http.ResponseWriter,req *http.Request){

	params:=mux.Vars(req)

	var NewEmployee Employees
	_ =json.NewDecoder(req.Body).Decode(&NewEmployee)
	NewEmployee.ID = params["id"]
	resources = append(resources,NewEmployee)
	json.NewEncoder(w).Encode(resources)

}

func DeleteEmployee(w http.ResponseWriter,req *http.Request){
	params:=mux.Vars(req)
	for index,value:=range resources{
		if value.ID == params["id"]{
			resources = append(resources[:index],resources[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(resources)
}


func main(){

   router := mux.NewRouter()

   resources = append(resources,Employees{"Shubham","17070002","GoDeveloper",
   &Location{"Delhi","Shahdara","110032"}})

	resources = append(resources,Employees{"Kiran","17070003","GoDeveloper",
	&Location{"Maharashtra","Pune","220031"}})

	resources = append(resources,Employees{"Ganesh","17070004","AndroidDeveloper",
	&Location{"Delhi","GTB Nagar","110029"}})

	resources = append(resources,Employees{"Nirbheek","17070005","AndroidDeveloper",
	&Location{"Delhi","Pitampura","110053"}})

	resources = append(resources,Employees{"Ambuj","17070006","AssociateConsultant",
	&Location{"Haryana","Bahadurgarh","170047"}})

	resources = append(resources,Employees{"Shubham","17070007","AssociateConsultant",
	&Location{"Delhi","Rohini","110059"}})


    router.HandleFunc("/employees",GetEmployeesDetails).Methods("GET")
	router.HandleFunc("/employees/{name}/{id}",GetEmployeeDetails).Methods("GET")
	router.HandleFunc("/employees/{position}",GetEmployeeDetailsPosition).Methods("GET")
	router.HandleFunc("/employees/{name}/",GetEmployeeDetailsByName).Methods("GET")
	router.HandleFunc("/employee/{id}",GetEmployeeDetailsByID).Methods("GET")
	router.HandleFunc("/employees/{id}",CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}",DeleteEmployee).Methods("DELETE")

   log.Fatal(http.ListenAndServe(":8080",router))

}