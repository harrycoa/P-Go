package main

import "fmt"

// interface.. tipo de datos estructura, vacio
// solo va definicion , con metodos que no estan implementados
type Usuarios interface {
	Permisos() int
	nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.nombre
}

func auth(usuario Usuarios) string {
	permisos := usuario.Permisos()
	if usuario.Permisos() > 4 {
		return usuario.Nombre() + "tiene permisos admin"
	} else if permisos == 3 {
		return usuario.Nombre() + "tiene permisos editos"
	}
	return ""
}

func main()  {

	admin := Admin{"harri"}
	fmt.Println(auth(admin))

	usuarios := []Usuarios{Admin{"A"}}
	for _,us := range usuarios{
		fmt.Println(auth(us))
	}

}