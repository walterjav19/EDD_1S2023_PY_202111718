package lectura

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Leer(path string) [][]string {
	// Abrir el archivo
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
	}
	defer file.Close()

	// Crear un lector CSV
	reader := csv.NewReader(file)

	// Leer todas las filas
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
	}
	
	if len(rows) > 1 {
		rows = rows[1:]
	} else {
		rows = nil
	}
	// Imprimir las filas
	return rows
}
