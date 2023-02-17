package main

import (
	"estructuras"
	"fmt"
	"menus"
	"roles"
)

func main() {

	pila := estructuras.Stack{}
	pila.Push(1)
	pila.Push(2)
	pila.Push(3)
	pila.Recorrer()
	fmt.Println(pila.Top())

	var option int
	// iniciamos el administrador
	admin := roles.Admin{"admin", "admin"}

	for {
		menus.MenuInicio()
		fmt.Print("Elija su opcion: ")
		fmt.Scanln(&option)

		switch option {
		case 1:

			var usuario, contrasenia string

			fmt.Print("\nIngresa tu Usuario: ")
			fmt.Scanln(&usuario)
			fmt.Print("Ingresa tu contraseña: ")
			fmt.Scanln(&contrasenia)

			if usuario == admin.User && contrasenia == admin.Password {
			dashboard:
				for {
					var op2 int
					menus.MenuAdmin()
					fmt.Print("Elija su opcion: ")
					fmt.Scanln(&op2)
					switch op2 {
					case 1:
						fmt.Println("ESTUDIANTES PENDIENTES")
					case 2:
						fmt.Println("ESTUDIANTES DEL SISTEMA")
					case 3:
						fmt.Println("Rgistrar Nuevo Estudiante")
					case 4:
						fmt.Println("Carga masiva de estudiantes")
					case 5:
						break dashboard
					}
				}

			} else {
				fmt.Println("recorrer lista")
			}
		case 2:
			menus.Adios()
			return
		default:
			fmt.Println("Opción inválida elija 1 o 2")
		}
	}

}
