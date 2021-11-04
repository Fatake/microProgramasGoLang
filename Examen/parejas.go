package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"

	util "github.com/Fatake/fatakeGo"
)

type Persona struct {
	nombre string
	edad   int
	sexo   string
}

func (p Persona) ToString() string {
	aux := "Nombre:" + p.nombre + " Edad:" + strconv.Itoa(p.edad) + " Sexo:" + p.sexo
	return aux
}

/**
Se dispone de una relacion de n-personas de sexo masculono y femenino..
Imprementa un programa en que pida los datos de estas n-personas (Nombre, edad, sexo)
y sean amalmacenados en estructuras,
y a continuaicon muestre todas las combinaciones posibles de parejas
que pueden contraer matrimonio, siempre y cuando ambas sean de distinto sexo y
diferencia de edad no supere los 10 años]
*/
func listarPersonas(lista []Persona) {
	fmt.Println("Listado de Personas")
	for i, p := range lista {
		fmt.Printf("[%d] %s\n", (i + 1), p.ToString())
	}
	fmt.Println("---------------------------")
}

func agregaPersona(lista *[]Persona) {
	var numero int
	var line string
	fmt.Println("Agregar Personas")
	fmt.Printf("[+] Ingresa el numero de personas a agregar\n-> ")
	fmt.Scanf("%d", &numero)
	fmt.Scanf("%d", &numero)

	for i := 0; i < numero; i++ {
		fmt.Printf("[i] Ingrese la Persona %d de la siguiente forma: nombre,edad,M/F \nP/E: Pedro,21,M\n-> ", (i + 1))
		fmt.Scanf("%s", &line)
		fmt.Scanf("%s", &line)
		strs := strings.Split(line, ",")
		n, _ := strconv.Atoi(strs[1])
		lista = append(lista[0:], Persona{strs[0], n, strs[2]})
	}
}

func menu() int {
	op := 9
	for {
		fmt.Println("\tPrograma de N-Parejas")
		fmt.Println("Seleccione una opcion:")
		fmt.Println("[1] Ver Lista de Personas")
		fmt.Println("[2] Agregar Persona(s)")
		fmt.Println("[3] Eliminar a una Persona")
		fmt.Println("[4] Generar Parejas")
		fmt.Println("[0] Salir !!")
		fmt.Printf("------------------------\n-> ")
		fmt.Scanf("%d", &op)
		fmt.Scanf("%d", &op)
		util.CleanScreen()
		if op > 0 && op < 5 {
			break
		}
	}
	return op
}

func main() {
	var nPersonas int
	var line string
	lista := make([]Persona, 1, 1000)

	fmt.Println("\tPrograma de N-Parejas")
	fmt.Printf("Ingresa el numero de personas\n-> ")
	fmt.Scanf("%d", &nPersonas)

	util.CleanScreen()
	fmt.Printf("[i] Numero de personas Ingresadas: %d\n", nPersonas)

	for i := 0; i < nPersonas; i++ {
		fmt.Printf("[i] Ingrese la Persona %d de la siguiente forma: nombre,edad,M/F \nP/E: Pedro,21,M\n-> ", (i + 1))
		fmt.Scanf("%s", &line)
		fmt.Scanf("%s", &line)
		strs := strings.Split(line, ",")
		n, _ := strconv.Atoi(strs[1])
		lista = append(lista, Persona{strs[0], n, strs[2]})
	}

	util.CleanScreen()
	lista = lista[1:]
	fmt.Println(lista)
	for {
		opcion := menu()
		switch opcion {
		case 1: // Listar
			util.CleanScreen()
			listarPersonas(lista)

		case 2:
			util.CleanScreen()
			agregaPersona(&lista)

		case 3:
			util.CleanScreen()
			listarPersonas(lista)
		case 4:
			util.CleanScreen()
			listarPersonas(lista)
		case 0:
			fmt.Println("Adios u.u")
			syscall.Exit(0)
		}
	}
}
