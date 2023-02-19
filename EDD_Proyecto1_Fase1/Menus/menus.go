package menus

import "fmt"

func Desicion() {
	fmt.Println("*                1. Aceptar Estudiante        *\n" +
		"*                2. Rechazar Estudiante       *\n" +
		"*                3. Volver Menu               *")
}

func Vacia() {
	fmt.Println(`      ___           ___           ___                       ___     
     /\__\         /\  \         /\  \          ___        /\  \    
    /:/  /        /::\  \       /::\  \        /\  \      /::\  \   
   /:/  /        /:/\:\  \     /:/\:\  \       \:\  \    /:/\:\  \  
  /:/__/  ___   /::\~\:\  \   /:/  \:\  \      /::\__\  /::\~\:\  \ 
  |:|  | /\__\ /:/\:\ \:\__\ /:/__/ \:\__\  __/:/\/__/ /:/\:\ \:\__\
  |:|  |/:/  / \/__\:\/:/  / \:\  \  \/__/ /\/:/  /    \/__\:\/:/  /
  |:|__/:/  /       \::/  /   \:\  \       \::/__/          \::/  / 
   \::::/__/        /:/  /     \:\  \       \:\__\          /:/  /  
    ~~~~           /:/  /       \:\__\       \/__/         /:/  /   
                   \/__/         \/__/                     \/__/`)
}

func Bienvenido() {
	banner := []string{
		".______    __   _______ .__   __. ____    ____  _______ .__   __.  __   _______   ______   ",
		"|   _  \\  |  | |   ____||  \\ |  | \\   \\  /   / |   ____||  \\ |  | |  | |       \\ /  __  \\  ",
		"|  |_)  | |  | |  |__   |   \\|  |  \\   \\/   /  |  |__   |   \\|  | |  | |  .--.  |  |  |  | ",
		"|   _  <  |  | |   __|  |  . `  |   \\      /   |   __|  |  . `  | |  | |  |  |  |  |  |  | ",
		"|  |_)  | |  | |  |____ |  |\\   |    \\    /    |  |____ |  |\\   | |  | |  '--'  |  `--'  | ",
		"|______/  |__| |_______||__| \\__|     \\__/     |_______||__| \\__| |__| |_______/ \\______/  ",
	}

	for _, line := range banner {
		fmt.Println(line)
	}
}

func Adios() {
	banner := []string{
		" ______      __                        ",
		"/\\  _  \\    /\\ \\  __                   ",
		"\\ \\ \\L\\ \\   \\_\\ \\/\\_\\    ___     ____  ",
		" \\ \\  __ \\  /'_` \\/\\ \\  / __`\\  /',__\\ ",
		"  \\ \\ \\/\\ \\/\\ \\L\\ \\ \\ \\/\\ \\L\\ \\/\\__, `\\",
		"   \\ \\_\\ \\_\\ \\___,_\\ \\_\\ \\____/\\/\\____/",
		"    \\/_/\\/_/\\/__,_ /\\/_/\\/___/  \\/___/  ",
	}
	for _, line := range banner {
		fmt.Println(line)
	}
}

func MenuAdmin() {
	fmt.Println("\n*** Dashboard Administrador - EDD GoDrive ***\n" +
		"*       1. Ver Estudiantes Pendientes       *\n" +
		"*       2. Ver Estudiantes del Sistema      *\n" +
		"*       3. Registrar Nuevos Estudiantes     *\n" +
		"*       4. Carga Masiva de Estudiantes      *\n" +
		"*       5. Cerrar Sesion                    *\n" +
		"*********************************************")
}


func MenuUsuario(nombre,apellido,carnet string) {
	fmt.Println("\n***    Dashboard Usuario - EDD GoDrive    ***\n" +
				  "*       Usuario: "+nombre+" "+apellido+"    \n" +
		          "*       Carnet: "+carnet+"                  \n"+
				  "*********************************************")
}


func MenuInicio() {
	fmt.Println("\n**************** EDD GoDrive ****************\n" +
		"*             1. Iniciar Sesion             *\n" +
		"*            2. Salir del Sistema           *\n" +
		"*********************************************")
}
