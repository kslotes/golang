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
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ',' // Use this if the delimiter is a comma (default)

	// Crear un slice para almacenar las órdenes
	var orders []Order

	// Leer registros
	for {
		record, err := reader.Read()
        if err == io.EOF {
            break
        }
		if err != nil {
			log.Fatal(err)
		}

		// Crear una instancia de Order y llenarla con los datos del registro
		order := Order{
			OrderNumber:   record[0],
			OrderType:     record[1],
			AssetType:     record[2],
			FiatType:      record[3],
			TotalPrice:    record[4],
			Price:         record[5],
			Quantity:      record[6],
			ExchangeRate:  record[7],
			Counterparty:  record[8],
			Status:        record[9],
			CreatedTime:   record[10],
		}

		orders = append(orders, order)
	}

	// Imprimir las órdenes
}