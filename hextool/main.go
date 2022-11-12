package forth

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/alexeyneu/rino2/on_green"
	)

type bomb struct {
	PrivateKey string `json:"privateKey"`
	Chain      string `json:"chain"`     
	Testnet    bool   `json:"testnet"`
}
type merv map[string]bomb

func forth() {

	data, err := ioutil.ReadFile("ready.json")
	var b merv
	json.Unmarshal(data, &b)

 }