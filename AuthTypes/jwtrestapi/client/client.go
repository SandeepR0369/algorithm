package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go" //GT
)

//var mySigningKey = []byte("secretphrase")
var mySigningKey = []byte(os.Getenv("MY_JWT_TOKEN")) //GT

//generating tokens (GT)
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "sandeep"
	claims["expires"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Error1: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT() //GT
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:1001/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func main() {
	fmt.Println("Client Side")

	/*tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Issue with generation of tokenString", err)
	}*/
	//fmt.Println("token", tokenString)
	v, j := GenerateJWT()
	fmt.Println("VT", v, j)
	handleRequests()
}
