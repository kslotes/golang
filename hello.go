package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Order struct {
	OrderNumber   string
	OrderType     string
	AssetType     string
	FiatType      string
	TotalPrice    string
	Price         string
	Quantity      string
	ExchangeRate  string
	Counterparty  string
	Status        string
	CreatedTime   string
}

func main() {
	file, err := os.Open("ejemplo.csv")
	if err != nil {
        log.Fatal(err);		
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)


	// Leer registros
	for {
		record, err := reader.Read()
        if err == io.EOF {
            break
        }
		if err != nil {
			log.Fatal(err)
		}

        for value := range record {
            fmt.Printf("%s\n", record[value])
        }
		// Crear una instancia de Order y llenarla con los datos del registro

	}

	// Imprimir las órdenes
}
