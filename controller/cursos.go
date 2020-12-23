package controller

import "fmt"

type Curso struct {
	name    string
	alumnos []Alumno
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

func AgregarAlumnoACurso(alumno Alumno, nombreCurso string) {
	ifExist := ifExistCurso(nombreCurso)
	ifExistDni := false

	if ifExist == false {
		fmt.Println("El curso al cual desea agregar un alumno, no existe")
	} else {

		isExistAlumno := IfExistByDni(alumno.dni)

		if isExistAlumno == true {
			for i, _ := range CursoArray {
				if CursoArray[i].name == nombreCurso {
					for j, _ := range CursoArray[i].alumnos {
						if CursoArray[i].alumnos[j].dni == alumno.dni {
							fmt.Println("El alumno", CursoArray[i].alumnos[j].apellido, "ya se encuentra en el curso de :", CursoArray[i].name)
							ifExistDni = true
						}
					}
					if ifExistDni != true {
						CursoArray[i].alumnos = append(CursoArray[i].alumnos, Alumno{
							nombre:   alumno.nombre,
							apellido: alumno.apellido,
							dni:      alumno.dni,
						})
					}
				}
			}
		} else {
			fmt.Println("El Alumno que desea agregar al Curso no se encuentra..")
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
