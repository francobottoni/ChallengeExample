package main

import (
	"fmt"
	"skillfactory/controller"
)

func main() {
	alumno1 := controller.NewAlumno("Franco", "Bottoni", 39098710)
	alumno2 := controller.NewAlumno("Sofia", "Perez", 39043711)
	alumno3 := controller.NewAlumno("Marcos", "Aranda", 39093212)
	alumno4 := controller.NewAlumno("Maxi", "Gonzalez", 37128713)
	alumno5 := controller.NewAlumno("Carla", "Sanchez", 30293454)

	controller.AddAlumno(alumno1)
	controller.AddAlumno(alumno2)
	controller.AddAlumno(alumno3)
	controller.AddAlumno(alumno4)
	controller.AddAlumno(alumno5)

	fmt.Println(controller.AlumnoArray)

	controller.DeleteAlumno(39093212)

	fmt.Println(controller.AlumnoArray)

	controller.AgregarCurso("Matematica")
	controller.AgregarCurso("Literatura")
	controller.AgregarCurso("Ingles")

	fmt.Println(controller.CursoArray)

	controller.AgregarAlumnoACurso(alumno4, "Matematica")
	controller.AgregarAlumnoACurso(alumno2, "Matematica")
	controller.AgregarAlumnoACurso(alumno1, "Matematica")

	fmt.Println(controller.CursoArray)

	controller.AgregarAlumnoACurso(alumno5, "Literatura")
	controller.AgregarAlumnoACurso(alumno3, "Literatura")

	fmt.Println(controller.CursoArray)

	controller.AgregarAlumnoACurso(alumno2, "Ingles")
	controller.AgregarAlumnoACurso(alumno2, "Ingles")

	fmt.Println(controller.CursoArray)

}
