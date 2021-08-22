package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/linker-fan/gal-anonim-cli/internal/jwt"
)

func SendRequest(addr string, method string, payload *map[string]interface{}) ([]byte, error) {
	var reqBody []byte
	var err error
	if payload != nil {
		reqBody, err = json.Marshal(*payload)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	//timeout set to two minutes
	timeout := time.Duration(2 * time.Minute)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(method, addr, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf("Response returned status code: %d\nrzad", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return body, nil

}

func prepareAuthHeaders(req *http.Request, path string) (*http.Request, error) {
	//get token from file
	token, err := jwt.LoadTokenFromFile(path)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return req, nil
}
