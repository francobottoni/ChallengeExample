package controller

import (
	"fmt"
)

type Alumno struct {
	nombre   string
	apellido string
	dni      int32
}

var AlumnoArray []Alumno

func NewAlumno(nombre string, apellido string, dni int32) Alumno {
	return Alumno{
		nombre:   nombre,
		apellido: apellido,
		dni:      dni,
	}
}

func AddAlumno(a Alumno) {
	ifExist := IfExistByDni(a.dni)

	if ifExist == true {
		fmt.Println("Ya hay un alumno cargado con ese DNI")
	} else {
		AlumnoArray = append(AlumnoArray, Alumno{
			nombre:   a.nombre,
			apellido: a.apellido,
			dni:      a.dni,
		})
	}

}

func DeleteAlumno(dni int32) {
	ifExist := IfExistByDni(dni)
	if ifExist == true {
		for i := 0; i < len(AlumnoArray); i++ {
			if AlumnoArray[i].dni == dni {
				fmt.Println(i)
				AlumnoArray = append(AlumnoArray[:i], AlumnoArray[i+1:]...)
			}
		}
	} else {
		fmt.Println("No se encontro un alumno con ese DNI")

	}

}

func IfExistByDni(dni int32) bool {
	ifExist := false
	for i, _ := range AlumnoArray {
		if AlumnoArray[i].dni == dni {
			ifExist = true
		}
	}

	return ifExist
}
