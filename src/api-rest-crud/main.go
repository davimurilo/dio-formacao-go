package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	listaBolos = []Bolo{
		Bolo{1, "Morango", 50.00},
		Bolo{2, "Chocolate", 40.00},
		Bolo{3, "Abacaxi", 30.00},
	}

	mutex sync.Mutex
)

type RetornoAPI struct {
	Sucesso  bool   `json:"sucesso"`
	Mensagem string `json:"mensagem"`
	Objeto   any    `json:"objeto"`
}

type Bolo struct {
	Codigo int32   `json:"codigo"`
	Sabor  string  `json:"sabor"`
	Preco  float32 `json:"preco"`
}

func main() {

	servidorAPI := http.NewServeMux()

	servidorAPI.HandleFunc("/listar", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RetornoAPI{Sucesso: false, Mensagem: "NOK"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(RetornoAPI{Sucesso: true, Mensagem: "OK", Objeto: listaBolos})

	})

	servidorAPI.HandleFunc("/cadastrar", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		defer r.Body.Close()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RetornoAPI{Sucesso: false, Mensagem: "Metodo inválido"})
			return
		}

		var bolo Bolo

		erro := json.NewDecoder(r.Body).Decode(&bolo)

		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RetornoAPI{Sucesso: false, Mensagem: "NOK"})
		}

		mutex.Lock()
		bolo.Codigo = listaBolos[len(listaBolos)-1].Codigo + 1
		listaBolos = append(listaBolos, bolo)
		mutex.Unlock()

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(RetornoAPI{Sucesso: true, Mensagem: "OK", Objeto: listaBolos})
	})

	servidorAPI.HandleFunc("/remover/{id}", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RetornoAPI{Sucesso: false, Mensagem: "Metodo inválido"})
			return
		}

		codigo, _ := strconv.Atoi(r.PathValue("id"))
		w.WriteHeader(http.StatusOK)
		var novaLista []Bolo

		for _, bolo := range listaBolos {
			if bolo.Codigo != int32(codigo) {
				novaLista = append(novaLista, bolo)
			}
		}

		listaBolos = novaLista

		json.NewEncoder(w).Encode(RetornoAPI{Sucesso: true, Mensagem: "OK", Objeto: listaBolos})
	})

	servidorAPI.HandleFunc("/atualizar/{id}", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RetornoAPI{Sucesso: false, Mensagem: "Metodo inválido"})
			return
		}

		codigo, _ := strconv.Atoi(r.PathValue("id"))
		w.WriteHeader(http.StatusOK)

		var boloAtualizar Bolo

		erro := json.NewDecoder(r.Body).Decode(&boloAtualizar)

		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RetornoAPI{Sucesso: false, Mensagem: "NOK"})
		}

		for indice, bolo := range listaBolos {
			if bolo.Codigo == int32(codigo) {
				bolo.Sabor = boloAtualizar.Sabor
				bolo.Preco = boloAtualizar.Preco
				listaBolos[indice] = bolo
			}
		}

		json.NewEncoder(w).Encode(RetornoAPI{Sucesso: true, Mensagem: "OK", Objeto: listaBolos})
	})

	erro := http.ListenAndServe(":8089", servidorAPI)

	if erro != nil {
		log.Fatal("Erro ao subir o servidor : " + erro.Error())
	}
}
