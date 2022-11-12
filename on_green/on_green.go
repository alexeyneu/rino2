package on_green
import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"crypto/elliptic"
	"reflect"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"sync"
  "strings"
	"github.com/btcsuite/btcd/btcutil/base58"
 	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)



type brand struct {
	Balance      int   `json:"balance"`
	CreateTime   int64 `json:"createTime"`
	FreeNetUsage int   `json:"freeNetUsage"`
	Trc10        []struct {
		Value int    `json:"value"`
		Key   string `json:"key"`
	} `json:"trc10"`
	Trc20 []struct {
		TR7NHqjeKQxGTCi8Q8ZY4PL8OtSzgjLj6T string `json:"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t,omitempty"`
	} `json:"trc20"`
}

type brazer struct {
	TxID string `json:"txId"`
}

//Bitcore
func generateKeyPair() (pubkey, privkey []byte) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	pubkey = elliptic.Marshal(secp256k1.S256(), key.X, key.Y)
	return pubkey, math.PaddedBigBytes(key.D, 32)
}
func hash(s []byte) []byte {
        h := sha256.New()
        h.Write(s)
        return h.Sum(nil)
}

func Make_c() (string, string) {
	fog, seckey := generateKeyPair()
	dm := fog[1:]
    bt := crypto.Keccak256Hash(dm)
	t := bt[11:]
	t[0] = 0x41
	b := hash(t);
	b = hash(b);
	ch := append(t, b[:4]...)
	tdk := base58.Encode(ch)
	f := hex.EncodeToString(seckey);
	fmt.Println(f)
	fmt.Println(tdk)
	return tdk, f
}

func Make_from(wessst string) string {
	pr, _ := crypto.HexToECDSA(wessst)
	edcpub := pr.PublicKey
	fog := elliptic.Marshal(secp256k1.S256(), edcpub.X, edcpub.Y)
	dm := fog[1:]
  bt := crypto.Keccak256Hash(dm)
	t := bt[11:]
	t[0] = 0x41
	b := hash(t);
	b = hash(b);
	ch := append(t, b[:4]...)
	tdk := base58.Encode(ch)
	fmt.Println(tdk)
	return tdk
}


func Made(tdk string) string {
  esUrl := "https://api-eu1.tatum.io/v3/tron/account/" + tdk
  es, _ := http.NewRequest("GET", esUrl, nil)
  es.Header.Add("x-api-key", os.ExpandEnv("$API_GUN"))
  var rp *http.Response
  var wg sync.WaitGroup
 	wg.Add(1)
  go func() {
  	defer wg.Done()
		rp, _ = http.DefaultClient.Do(es)
	}()
	wg.Wait()
	defer rp.Body.Close()
  stuff, _ := ioutil.ReadAll(rp.Body)
  var x brand 
  json.Unmarshal(stuff, &x)
	for _, elm := range x.Trc20 {
		if reflect.ValueOf(elm).IsZero()	{ 

		} else {
			fmt.Println(elm.TR7NHqjeKQxGTCi8Q8ZY4PL8OtSzgjLj6T)
			return 	elm.TR7NHqjeKQxGTCi8Q8ZY4PL8OtSzgjLj6T
		}

	}
	return *new(string)
}

func Made_from(target string) string {
  esUrl := "https://api.tatum.io/v3/tron/trc20/transaction"
  es, _ := http.NewRequest("POST", esUrl, strings.NewReader(target))
  es.Header.Add("x-api-key", os.ExpandEnv("$API_GUN"))
  es.Header.Add("Content-Type","application/json")

  var rp *http.Response
  var wg sync.WaitGroup
 	wg.Add(1)
  go func() {
  	defer wg.Done()
		rp, _ = http.DefaultClient.Do(es)
	}()
	wg.Wait()
	defer rp.Body.Close()
  stuff, _ := ioutil.ReadAll(rp.Body)
  var ff brazer
  json.Unmarshal(stuff, &ff)
	return ff.TxID
}
