package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type TestConfig struct {
	TestTokenID string `json: "testTokenID"`
	TestRoomID  string `json: "testRoomID"`
}

var (
	TConfig *TestConfig
)

func TestMain(m *testing.M) {
	file, err := ioutil.ReadFile("../.secrets")
	if err != nil {
		log.Fatal(err)
	}
	TConfig = &TestConfig{}

	err = json.Unmarshal([]byte(file), TConfig)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())

}
