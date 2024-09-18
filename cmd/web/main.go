package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pobute/subscription-service/cmd/pkg/config"
	"github.com/pobute/subscription-service/cmd/pkg/handlers"
)

const webPort = ":8080"
var session *scs.SessionManager
var app config.AppConfig

// // myMap := make(map[string]string)
// myMap["dog"] = "Fido"
// myMap["cat"] = "Fluffy"

// mySlice := []int {1,2,3,4}

var mySlice2 []int

var myVar string

type Human interface{
	Speak() string

	Walk() string
}

type Person struct {
	FirstName string
	LastName  string
	HairColor string
	HasDog    bool
}



func main() {

	result, err := divide(100.0, 0.5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The result of division is", result)

	

	

	
	myJson := `
	[
	{
	"first_name": "Clark",
	"last_name": "Kent",
	"Hair_color": "Black"
	"has_dog": true
	},
	{
	"first_name": "Bruce",
	"last_name": "Wayne",
	"Hair_color": "Black"
	"has_dog": false
	}
	]`

	var unmarshalled []Person

	err = json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}
	fmt.Printf("unmarshalled: %v", unmarshalled)

	var mySlice []Person

	person1 := Person{
		FirstName: "Peace",
		LastName:  "Obute",
		HairColor: "Black",
		HasDog:    false,
	}

	person2 := Person{
		FirstName: "Aishat",
		LastName:  "Moshood",
		HairColor: "Black",
		HasDog:    false,
	}

	mySlice = append(mySlice, person1)
	mySlice = append(mySlice, person2)

	newJson, err := json.MarshalIndent(mySlice, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(newJson))

	// var user [] int

	myMap := make(map[string]string)
	myMap["dog"] = "Fido"
	myMap["cat"] = "Fluffy"

	myVar = "dog"

	switch myVar {
	case "cat":
		fmt.Println("Cat is set to cat")

	case "dog":
		fmt.Println("Cat is set to dog")
	case "fish":
		fmt.Println("Cat is set to fish")
	default:
		fmt.Println("Cat is set to something else ")
	}

	fmt.Println(sayHello("Peace"))

	////////////////////////////////////////////
	///////////////////////////////////////////
	/**********for building apps**************/



	//change this to true when in production
	app.InProduction = false

	//sessions are used to store user  authentication details accross multiple requests for a time so the page can know who the user is 
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction


    app.Session = session
	
	app.UseCache = false
	repo:= handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port 8080")

	srv:= &http.Server{
		Addr: webPort,
		Handler: routes(&app),
	}
	
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}





func sayHello(word string) string {
	myName := word + "a"

	return myName
}

func divide(x, y float32) (float32, error) {

	var result float32
	if y <= 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil
}

func (p *Person) Speak() string {
	return "Hello"
}


func (p *Person) Walk() string {
	return "Walk"
}
