package tickets

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Ticket struct
type Ticket struct {
	id             uint16
	passengerName  string
	passengerEmail string
	destination    string
	time           string
	price          uint16
}

func LoadTickets(filename string) ([]Ticket, error) {
	// Leer el archivo csv
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al abrir el archivo CSV:", err)
		return nil, err
	}
	defer file.Close()

	//Crear el lector de archivos CSV
	reader := csv.NewReader(file)
	reader.Comma = ','

	// Crear el slice de tickets
	tickets := make([]Ticket, 0)

	// Leer el archivo CSV y crear los tickets
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// Si está al final del archivo, salir del bucle
				break
			}
			fmt.Println("Error al leer una fila:", err)
			return nil, err
		}

		// Se convierte el id de string a uint16
		num, _ := strconv.ParseUint(record[0], 10, 16)
		uint16num := uint16(num)

		// Se convierte el precio de string a uint16
		priceNum, _ := strconv.ParseUint(record[5], 10, 16)
		uint16PriceNum := uint16(priceNum)

		// Crear un nuevo ticket y asignar los valores
		ticket := Ticket{
			id:             uint16num,
			passengerName:  record[1],
			passengerEmail: record[2],
			destination:    record[3],
			time:           record[4],
			price:          uint16PriceNum,
		}

		// Agregar el ticket al slice
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

// Requerimiento 1
func GetTotalTickets(tickets []Ticket, destination string) (int, error) {
	total := 0
	for _, ticket := range tickets {
		if ticket.destination == destination {
			total++
		}
	}
	return total, nil
}

// Requerimiento 2
func GetCountByPeriod(tickets []Ticket) map[string]int {
	passengersByTime := map[string]int{
		"madrugada": 0,
		"manana":    0,
		"tarde":     0,
		"noche":     0,
	}

	total := 0
	for _, ticket := range tickets {
		hourStr := strings.Split(ticket.time, ":")[0]
		hour, err := strconv.Atoi(hourStr)
		if err != nil {
			fmt.Println("Error al convertir la hora a entero:", err)
		}

		switch {
		case hour >= 0 && hour < 6:
			passengersByTime["madrugada"]++
		case hour >= 7 && hour < 12:
			passengersByTime["manana"]++
		case hour >= 13 && hour < 19:
			passengersByTime["tarde"]++
		case hour >= 20 && hour <= 24:
			passengersByTime["noche"]++
		}
		total++
	}
	return passengersByTime
}

// Requerimiento 3
func AverageDestination(tickets []Ticket, destination string) (float64, error) {
	totalPassengers := len(tickets)
	totalPassengersByDestination, _ := GetTotalTickets(tickets, destination)

	if totalPassengers == 0 {
		return 0.0, fmt.Errorf("No se encontraron tickets para hoy")
	}

	average := float64(totalPassengersByDestination) / float64(totalPassengers)
	return average, nil
}
