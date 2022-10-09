package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/ashishjuyal/banking/service"
)


type Customer struct {
	Name 	string `json:"full_name" xml:"name"`
	City 	string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World")
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer {
	// 	{Name: "Eduardo", City: "Santiago", Zipcode: "111111" },
	// 	{Name: "Rob", City: "Valdivia", Zipcode: "22222" },
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

