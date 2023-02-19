package main

import (
	"estructuras"
	"fmt"
	"lectura"
	"menus"
	"roles"
	"strconv"
	"strings"
)

func main() {

	ColaAlumnos := &estructuras.Queu{}
	ListAlumnos := &estructuras.LinkedList{}
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
						if ColaAlumnos.IsEmpty() {
							menus.Vacia()
						} else {
							var desicion int

							fmt.Println("\n********** Estudiantes Pendientes **********\n")
						enlistar:
							for desicion != 3 && !ColaAlumnos.IsEmpty() {
								fmt.Println("**********       Pendientes:" + strconv.Itoa(ColaAlumnos.Size()) + "     **********")
								fmt.Println("*    Estudiante Actual: " + ColaAlumnos.Cabeza().Nombre + " " + ColaAlumnos.Cabeza().Apellido + "         *")
								menus.Desicion()
								fmt.Print("\nElige una Opcion: ")
								fmt.Scanln(&desicion)
								switch desicion {
								case 1:
									fmt.Println("Estudiante Aceptado en el Sistema!!!!")
									a := ColaAlumnos.Dequeu()
									ListAlumnos.Append(a) //agrego el desencolado a la lista
									
								case 2:
									ColaAlumnos.Dequeu() //desencolar sin hacer nada mas ya que se les rechazo
									fmt.Print("\nAlumno Rechazado !!! \n")
								case 3:
									fmt.Println("\nSaliendo !!!!")
									break enlistar
								default:
									fmt.Println("elija entre el 1-3")

								}
							}

						}
					case 2:
						//ListAlumnos.Sort()    //sort me esta petando aaa ojo
						fmt.Println("**********Listado de Estudiantes**********")
						ListAlumnos.Recorrer()
						
					case 3:
						fmt.Println("***** Registro de Estudiantes - EDD GoDrive *****")
						var nombre, apellido, password string
						var carnet int
						fmt.Print("Ingresa Nombre: ")
						fmt.Scanln(&nombre)
						fmt.Print("Ingresa Apellido: ")
						fmt.Scanln(&apellido)
						fmt.Print("Ingresa tu carnet: ")
						fmt.Scanln(&carnet)
						fmt.Print("Ingresa tu contraseña: ")
						fmt.Scanln(&password)

						user := &roles.Student{nombre, apellido, carnet, password}
						ColaAlumnos.Enqueu(user)
						ColaAlumnos.GenerarDotCola()
					

						fmt.Println("\nEstudiante " + user.Nombre + " " + user.Apellido + " en cola de espera")

					case 4:
						var path string
						fmt.Print("Ingrese Ruta de Su Archivo CSV: ")
						fmt.Scanln(&path)
						MatrizDatos := lectura.Leer(path)
						for _, row := range MatrizDatos {
							NombreCompleto := strings.Split(row[1], " ")
							num, err := strconv.Atoi(row[0])
							if err != nil {
								fmt.Println("Trono")
							} else {
								Estudiante := &roles.Student{NombreCompleto[0], NombreCompleto[1], num, row[2]}
								ColaAlumnos.Enqueu(Estudiante)
								
							}

						}
						fmt.Println("\nAlumnos agregados a la Cola de Espera !!!")

						
					case 5:
						break dashboard
					}
				}

			} else {
				usuario,err:=strconv.Atoi(usuario)
				if err != nil {
					fmt.Println("Asegurese de Ingresar como usuario su carnet")
					
				} else {
					if !ListAlumnos.Empty(){
						if ListAlumnos.Validar(usuario,contrasenia){
							menus.Bienvenido()
							p:=ListAlumnos.GetStudent(usuario,contrasenia)
							menus.MenuUsuario(p.Nombre,p.Apellido,strconv.Itoa(p.Carnet))
						}else{
							if ListAlumnos.ObtenerCarnet(usuario)==false{
								fmt.Println("Usuario Inexistente")
							}else{
								if ListAlumnos.ObtenerPassword(contrasenia)==false{
									fmt.Println("Contraseña No coincide con El Usuario")
								}
							} 
							

					}

					
					

						
					}else{
						fmt.Println("Aun No hay Estudiantes en la Plataforma")
					}	
				}


			}
		case 2:
			menus.Adios()
			return
		default:
			fmt.Println("Opción inválida elija 1 o 2")
		}
	}

}
