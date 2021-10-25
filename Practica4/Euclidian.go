package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CleanScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func discanciaEuclidiana(vector1, vector2 []int64) float64 {
	temp := 0.0
	for i := 0; i < int(len(vector1)); i++ {
		aux := float64(vector1[i] - vector2[i])
		temp += math.Pow(aux, 2)
	}

	return math.Sqrt(temp)
}

func main() {
	var vector1 []int64
	var vector2 []int64
	var linea string
	nSize := 0
	fmt.Println("Calculo de Distancia Euclidiana entre 2 vectores")
	fmt.Println("<-------------------------->")
	for {
		fmt.Print("[+] Ingresa tamaño de los vectores:\n")
		fmt.Scanf("%d", &nSize)
		if nSize > 0 {
			break
		}
		CleanScreen()
	}
	fmt.Printf("[i] Tamaño Ingresado: %d\n", nSize)
	fmt.Println("[+] Ingresa el primer vector (Al finalizar pon un / y enter):\n[i] Ejemplo: 1,2,3,4/")

	in := bufio.NewReader(os.Stdin)
	fmt.Scanln(&linea)
	line, err := in.ReadString('/')
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(line[0:len(line)-1], ",")
	vector1 = make([]int64, len(strs))
	for i, str := range strs {
		if vector1[i], _ = strconv.ParseInt(str, 0, 32); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("[i] Vector1: ", vector1)
	fmt.Scanln(&linea)
	fmt.Println("[+] Ingrese Segundo vector")
	in2 := bufio.NewReader(os.Stdin)
	line2, err := in2.ReadString('/')
	if err != nil {
		log.Fatal(err)
	}
	strs2 := strings.Split(line2[0:len(line2)-1], ",")
	vector2 = make([]int64, len(strs2))
	for i, str := range strs2 {
		if vector2[i], _ = strconv.ParseInt(str, 0, 32); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("[i] Vector2: ", vector2)

	fmt.Println("[+] La Distancia euclidiana es: ", discanciaEuclidiana(vector1, vector2))

}
