package services

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

//Wallet is struct for generate wallet
type Wallet struct {
	Address    string `json:"address"`
	PassPhrase string `json:"pass_phrase"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	KeyStore   string `json:"keystore_base64"`
}

//NewWallet return instance of Wallet
func NewWallet(passPhrase string) *Wallet {
	return &Wallet{PassPhrase: passPhrase}
}

//Generate return wallet generated
func (w *Wallet) Generate() error {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := key.NewAccount(w.PassPhrase)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir("./wallet")
	if err != nil {
		return err

	}
	if len(files) > 0 {

		b, err := ioutil.ReadFile(fmt.Sprintf("./wallet/%s", files[0].Name()))
		if err != nil {
			return err
		}
		keyStore, err := keystore.DecryptKey(b, w.PassPhrase)
		if err != nil {
			return err
		}
		keyStoreBase64 := base64.StdEncoding.EncodeToString(b)
		w.KeyStore = keyStoreBase64
		pvkBytes := crypto.FromECDSA(keyStore.PrivateKey)
		privateKey := hexutil.Encode(pvkBytes)
		w.PrivateKey = privateKey
		pbkBytes := crypto.FromECDSAPub(&keyStore.PrivateKey.PublicKey)
		publicKey := hexutil.Encode(pbkBytes)
		w.PublicKey = publicKey
		address := crypto.PubkeyToAddress(keyStore.PrivateKey.PublicKey).Hex()
		w.Address = address
	}
	if err := os.RemoveAll("./wallet"); err != nil {
		return err
	}

	return nil
}
