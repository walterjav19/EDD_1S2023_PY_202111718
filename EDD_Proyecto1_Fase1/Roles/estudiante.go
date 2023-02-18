package roles

import "strconv"

type Student struct {
	Nombre   string
	Apellido string
	Carnet   int
	Password string
}

func (s *Student) ToString() string {
	return "Nombre: " + s.Nombre + " " + s.Apellido + ", Carnet: " + strconv.Itoa(s.Carnet) + "\n" +
		"***********************************************"
}
