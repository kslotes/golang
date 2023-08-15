package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


type Order struct {
	orderNumber   string
	orderType     string
	assetType     string
	fiatType      string
	totalPrice    string
	price         string
	quantity      string
	exchangeRate  string
	counterparty  string
	status        string
	createdTime   string
}

type Summary struct {
    orders []Order
    fiatType string
    totalOrders int
    totalFiatQuantity float64
    totalFiatBought float64
    totalFiatSold float64
    totalRounds float64
}

type Round struct {
    orders []Order
    fiatType string
    totalOrders int
    totalFiatBought float64
    totalFiatSold float64
    avgBuyPrice float64
    avgSellPrice float64
}
func main() {
	file, err := os.Open("data.csv")
	if err != nil {
        log.Fatal(err);		
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)


	// Crear un slice para almacenar las órdenes
	var ordersARS []Order
	var ordersUSD []Order
    var orderQuantity int;
    var orderQuantityUSD int;
    var orderQuantityARS int;
    var fiatQuantityUSD float64;
    var fiatQuantityARS float64;
    const usdtPerRound = 2000;
    var arsRoundCount = 1;
    var usdAccumulated float64;
    var firstDate string = "9999999";
    var lastDate string = "-1";
    var rowNumber int;
    var totalArsSold float64;
    var totalArsBought float64;

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
        
        rowNumber++
		// Crear una instancia de Order y llenarla con los datos del registro
		order := Order{
			orderNumber:   record[0],
			orderType:     record[1],
			assetType:     record[2],
			fiatType:      record[3],
			totalPrice:    record[4],
			price:         record[5],
			quantity:      record[6],
			exchangeRate:  record[7],
			counterparty:  record[8],
			status:        record[9],
			createdTime:   record[10],
		}
        if order.createdTime > lastDate && rowNumber != 1{
            lastDate = order.createdTime;
        }

        if order.createdTime < firstDate && rowNumber != 1{
            firstDate = order.createdTime;
        }

        orderQuantity++
        if order.fiatType == "ARS" {
            ordersARS = append(ordersARS, order)
        }
        if order.fiatType == "USD" {
            ordersUSD = append(ordersUSD, order)
        }
	}

	// Imprimir las órdenes
    for _, order := range ordersARS {

        orderQuantityARS++
        totalPrice, err := strconv.ParseFloat(order.totalPrice, 64);
        quantity, err := strconv.ParseFloat(order.quantity, 64);

        if err != nil {
            log.Fatal(err)
        }

        if strings.ToUpper(order.orderType) == "SELL" {
            totalArsSold += totalPrice;
        }

        if strings.ToUpper(order.orderType) == "BUY" {
            totalArsBought += totalPrice;
        }

        usdAccumulated += quantity;
        if usdAccumulated >= usdtPerRound {
            usdAccumulated = 0;
            arsRoundCount++;
        }
        
        fiatQuantityARS += totalPrice;
    }

    for _, order := range ordersUSD {
        orderQuantityUSD++
        totalPrice, err := strconv.ParseFloat(order.totalPrice, 64);
        if err != nil {
            log.Fatal(err)
        }
        fiatQuantityUSD += totalPrice;
    }

    fmt.Println()
    fmt.Println("--- SUMMARY ---")
    fmt.Println("Total Orders: ", orderQuantity)
    fmt.Println("USD Orders: ", orderQuantityUSD)
    fmt.Println("ARS Orders: ", orderQuantityARS)
    fmt.Println()
    fmt.Println("firstDate registered: ", firstDate)
    fmt.Println("lastDate registered: ", lastDate)
    fmt.Println()
    fmt.Println("ARS Rounds: ", arsRoundCount)
    fmt.Printf("Total ARS Bought: $%.2f\n", totalArsBought)
    fmt.Printf("Total ARS Sold: $%.2f\n", totalArsSold)
    fmt.Printf("Outcome: $%.2f\n", totalArsSold - totalArsBought)
    fmt.Println()
    fmt.Printf("Total USD: %.2f\n", fiatQuantityUSD)
    fmt.Printf("Total ARS: %.2f\n", fiatQuantityARS)
    fmt.Println("--- END OF SUMMARY ---")
    fmt.Println()
	// Imprimir las órdenes
}
