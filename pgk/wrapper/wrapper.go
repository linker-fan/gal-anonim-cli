package wrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Login user struct
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Takes username and string as arguments returns token on error response
func Login(username string, password string) []bytes {
	user := &LoginRequest{
		Username: username,
		Password: password,
	}
	userJSON, err := json.Marshal(user)

	resp, err := http.Post("http://localhost:8888/users/login", "application/json",
		bytes.NewBuffer(userJSON))

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body

}
