package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// Cargar tickets
	ticketsSlice, errLoadTickets := tickets.LoadTickets("tickets.csv")
	if errLoadTickets != nil {
		fmt.Println("Error cargando los tickets", errLoadTickets)
		return
	}

	// Llamado a Requerimiento 1
	total, _ := tickets.GetTotalTickets(ticketsSlice, "Finland")
	fmt.Println("Requerimiento 1:", total)

	// Llamado a Requerimiento 2
	mapByPeriod := tickets.GetCountByPeriod(ticketsSlice)
	fmt.Println("Requerimiento 2:", mapByPeriod)

	//Llamado a Requerimiento 3
	averageByDestination, _ := tickets.AverageDestination(ticketsSlice, "Finland")
	fmt.Println("Requerimiento 3:", averageByDestination)
}
