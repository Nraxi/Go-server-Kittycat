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

// User struct which contains a name
// a type and a list of social links

type Cat struct {
	Name               string `json:"name"`
	Image              string `json:"image"`
	CutenessLevel      int    `json:"cutenessLevel"`
	AllergyInducingFur bool   `json:"allergyInducingFur"`
	LivesLeft          int    `json:"livesLeft"`
}

func Kittys(w http.ResponseWriter, r *http.Request) {

	fmt.Println(" Endpoint Hit: Endpoint hit (Api / Kittys) ")

	// Open our jsonFile
	jsonFile, err := os.Open("catdata.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened catdata.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var cats Cats

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	json.Unmarshal(byteValue, &cats)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(cats.Cats); i++ {
		fmt.Println(" Name: " + cats.Cats[i].Name)
		fmt.Println(" Image: " + cats.Cats[i].Image)
		fmt.Println(" CutenessLevel: " + strconv.Itoa(cats.Cats[i].CutenessLevel))
		fmt.Println(" LivesLeft: "+strconv.Itoa(cats.Cats[i].LivesLeft), "\n")
	}
	json.NewEncoder(w).Encode(cats.Cats)
}
