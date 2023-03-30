package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	tickets, err := LoadTickets("../../tickets.csv")
	assert.NoError(t, err, "Error al cargar los tickets para hacer tests")
	t.Run("Prueba para destino con tiquetes", func(t *testing.T) {
		total, err := GetTotalTickets(tickets, "Finland")
		assert.NoError(t, err, "Error al calcular el total de tickets para Finland")
		assert.Equal(t, 8, total, "El total de tickets para Finland debe ser 2")
	})
	t.Run("Prueba para destino sin tiquetes", func(t *testing.T) {
		total, err := GetTotalTickets(tickets, "Qatar")
		assert.NoError(t, err, "Error al calcular el total de tickets para Qatar")
		assert.Equal(t, 0, total, "El total de tickets para Qatar debe ser 0")
	})
}

func TestGetCountByPeriod(t *testing.T) {
	tickets, err := LoadTickets("../../tickets.csv")
	assert.NoError(t, err, "Error al cargar los tickets para hacer tests")
	passengersByTime := GetCountByPeriod(tickets)
	t.Run("Prueba para cantidad de pasajeros en la madrugada", func(t *testing.T) {
		assert.Equal(t, 260, passengersByTime["madrugada"], "El total de pasajeros en madrugada debe ser 260")
	})
	t.Run("Prueba para cantidad de pasajeros en la noche", func(t *testing.T) {
		assert.Equal(t, 151, passengersByTime["noche"], "El total de pasajeros en noche debe ser 151")
	})
}

func TestAverageDestination(t *testing.T) {
	tickets, err := LoadTickets("../../tickets.csv")
	assert.NoError(t, err, "Error al cargar los tickets para hacer tests")
	t.Run("Prueba para promedio de pasajeros para Finland", func(t *testing.T) {
		average, err := AverageDestination(tickets, "Finland")
		assert.NoError(t, err, "Error al calcular el promedio de pasajeros para Finland")
		assert.Equal(t, 0.008, average, "El promedio de pasajeros para Finland debe ser 0.008")
	})
	t.Run("Prueba para promedio de pasajeros para Brazil", func(t *testing.T) {
		average, err := AverageDestination(tickets, "Brazil")
		assert.NoError(t, err, "Error al calcular el promedio de pasajeros para Brazil")
		assert.Equal(t, 0.045, average, "El promedio de pasajeros para Brazil debe ser 0.045")
	})
}
