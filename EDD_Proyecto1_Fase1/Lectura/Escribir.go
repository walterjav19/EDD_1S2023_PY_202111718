package lectura

import (
	"fmt"
	"os"
)


func Escribir(texto string) {
    // Crear un archivo
    file, err := os.Create("C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Grafics/cola.dot")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Escribir en el archivo
    _, err = file.WriteString(texto)
    if err != nil {
        fmt.Println(err)
        return
    }

    
}