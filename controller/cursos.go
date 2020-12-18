package controller

import "fmt"

type Curso struct {
	name    string
	alumnos []Alumno
}

type TestInterface interface {
	AgregarAlumnoACurso(Alumno, string)
}

var CursoArray []Curso

func AgregarCurso(nombreCurso string) {
	ifExist := ifExistCurso(nombreCurso)

	if ifExist == true {
		fmt.Println("Ese curso ya existe en el sistema")
	} else {
		CursoArray = append(CursoArray, Curso{
			name: nombreCurso,
		})
	}

}

func (c *Curso) AgregarAlumnoACurso(alumno Alumno, nombreCurso string) {
	ifExist := ifExistCurso(nombreCurso)
	if ifExist == false {
		fmt.Println("El curso al cual desea agregar un alumno, no existe")
	} else {
		for i, _ := range CursoArray {
			if CursoArray[i].name == nombreCurso {

				c.alumnos = append(c.alumnos, Alumno{
					nombre:   alumno.nombre,
					apellido: alumno.apellido,
					dni:      alumno.dni,
				})

				CursoArray[i].alumnos = c.alumnos
			}
		}
	}
}

func ifExistCurso(nombreCurso string) bool {
	ifExist := false

	for i, _ := range CursoArray {
		if CursoArray[i].name == nombreCurso {
			ifExist = true
		}
	}

	return ifExist
}
