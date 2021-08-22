package services

import (
	"encoding/json"
	"log"

	"github.com/linker-fan/gal-anonim-cli/internal/config"
	"github.com/linker-fan/gal-anonim-cli/internal/jwt"
	"github.com/linker-fan/gal-anonim-cli/internal/requests"
)

var c config.Config

func Login(username, password, addr, pathToFile string) error {
	payload := map[string]interface{}{
		"username": username,
		"password": password,
	}

	data, err := requests.SendRequest(addr, "POST", &payload)
	if err != nil {
		log.Printf("Login: utils.SendRequest failed: %v\n", err)
		return err
	}

	type loginResponse struct {
		Token string `json:"token"`
	}
	var resp loginResponse

	err = json.Unmarshal(data, &resp)
	if err != nil {
		log.Printf("Login: json.Unmarshal failed: %v\n", err)
		return err
	}

	err = jwt.SaveTokenToFile(pathToFile, resp.Token)
	if err != nil {
		log.Printf("Login: SaveTokenToFile failed: %v\n", err)
		return err
	}

	return nil
}
