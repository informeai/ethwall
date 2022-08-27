package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/informeai/ethwall/dto"
	"github.com/informeai/ethwall/services"
)

func generateWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("{\"status\":\"error\",\"message\":\"method not allowed\"}"))
		return
	}
	defer r.Body.Close()
	passPhrase := dto.Payload{}
	if err := json.NewDecoder(r.Body).Decode(&passPhrase); err != nil {
		w.Write([]byte("{\"status\":\"error\",\"message\":\"decode error\"}"))
		return
	}
	wallet := services.NewWallet(passPhrase.PassPhrase)
	err := wallet.Generate()
	if err != nil {
		w.Write([]byte("{\"status\":\"error\",\"message\":\"error in generate\"}"))
		return
	}
	walletBytes, err := json.Marshal(wallet)
	if err != nil {
		w.Write([]byte("{\"status\":\"error\",\"message\":\"error in marshal wallet\"}"))
		return
	}
	w.Write(walletBytes)
}
func main() {
	f := http.FileServer(http.Dir("./static"))
	http.Handle("/", f)
	http.HandleFunc("/wallet", generateWallet)
	fmt.Printf("listen in port 4000\n")
	http.ListenAndServe(":4000", nil)
}
