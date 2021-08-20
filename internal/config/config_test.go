package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c, err := NewConfig("./../../config.yml")
	if err != nil {
		t.Fail()
	}

	fmt.Println(c)
}
