package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sample_svc/models"
	"github.com/sample_svc/service"
)
// RectangleArea is used to receive a request and send the response back
func RectangleArea(w http.ResponseWriter, r *http.Request) {

	params := &models.Rectangle{}

	//read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal body into the object called profile
	err = json.Unmarshal(body, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := fmt.Sprintf("The area of rectangle is : %v\n", service.Area(models.Rectangle{Length: params.Length, Breadth: params.Breadth}))

	b, _ := json.Marshal(result)

	w.WriteHeader(http.StatusCreated)
	w.Write(b)

}
