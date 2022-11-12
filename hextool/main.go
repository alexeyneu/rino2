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
type sierra map[string]string

func forth() {

	var tetr merv
	var fast sierra
	data, err := os.ReadFile("ready.json")
	json.Unmarshal(data, &tetr)
	for id,bomb_m := range tetr {
		fast[id] = on_green.Make_from(bomb_m.PrivateKey)
	}
	f = json.Marshal(fast)
	fmt.Println(string(f))
	os.WriteFile("western.json", f, 0777)
}