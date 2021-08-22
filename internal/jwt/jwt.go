package jwt

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/linker-fan/gal-anonim-cli/internal/utils"
)

type tokenJSONFile struct {
	Token string `json:"token"`
}

func SaveTokenToFile(path string, token string) error {

	err := utils.ValidatePath(path)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func() {
		file.Close()
	}()

	data := tokenJSONFile{Token: token}
	dataBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	_, err = file.Write(dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func LoadTokenFromFile(path string) (string, error) {
	var token tokenJSONFile

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer func() {
		file.Close()
	}()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, &token)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}
