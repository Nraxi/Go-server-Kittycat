package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Cats struct {
	Cats []Cat `json:"cats"`
}

type Cat struct {
	Name               string `json:"name"`
	Image              string `json:"image"`
	CutenessLevel      int    `json:"cutenessLevel"`
	AllergyInducingFur bool   `json:"allergyInducingFur"`
	LivesLeft          int    `json:"livesLeft"`
}

func Kittys(w http.ResponseWriter, r *http.Request) {

	fmt.Println(" Endpoint Hit: Endpoint hit (Api / Kittys) ")

	jsonFile, err := os.Open("catdata.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened catdata.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cats Cats

	json.Unmarshal(byteValue, &cats)

	for i := 0; i < len(cats.Cats); i++ {
		fmt.Println(" Name: " + cats.Cats[i].Name)
		fmt.Println(" Image: " + cats.Cats[i].Image)
		fmt.Println(" CutenessLevel: " + strconv.Itoa(cats.Cats[i].CutenessLevel))
		fmt.Println(" LivesLeft: "+strconv.Itoa(cats.Cats[i].LivesLeft), "\n")
	}
	json.NewEncoder(w).Encode(cats.Cats)
}
