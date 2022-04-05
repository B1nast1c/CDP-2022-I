package main

//ESTE CODIGO YA MIDE EL TIEMPO :)
import (
	"fmt"
	"math/rand"
	"time"
)

var SIZE int
var tiempoFinal int64
var tiempoInit time.Time

type Result struct {
	x1, y1, x2, y2            int
	matriz_1, matriz_2, final [][]int
}

func Init(a, b, c, d int) *Result {
	if b == c {
		temp_1 := make([][]int, a)
		for idx := range temp_1 {
			temp_1[idx] = rand.Perm(b)
		}

		temp_2 := make([][]int, c)
		for idx := range temp_2 {
			temp_2[idx] = rand.Perm(d)
		}

		temp_3 := make([][]int, a)
		for idx := range temp_3 {
			temp_3[idx] = make([]int, d)
		}

		return &Result{
			x1: a, y1: b, x2: c, y2: d,
			matriz_1: temp_1, matriz_2: temp_2, final: temp_3,
		}
	}
	return &Result{}
}

func generarMatrices(num int) (*Result, int) { //Funcion generadora de matrices
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(num) + 1
	SIZE = a
	generado := Init(a, a, a, a) //Genera una matriz cuadrada siempre, acomodado para otros casos
	return generado, a * a
}

func main() {
	var limite int
	fmt.Print("Ingrese limite maximo de dimensiones ")
	fmt.Scan(&limite)
	for idx := 1; idx <= limite; idx++ {
		productoMatricesClasico(idx)
	}
}

//Resolución de las matrices, a manera básica
func productoMatricesClasico(limit int) *Result {
	matriz, _ := generarMatrices(limit)
	for idxf, item := range matriz.final {
		for idxc := range item {
			tiempoInit = time.Now()
			productoIndividual(idxf, idxc, *matriz, tiempoInit) //Se crean hilos individuales para cada elemento de nuestra matriz
		}
	}
	defer fmt.Printf("%d %d \n", limit, tiempoFinal)
	return matriz
}

func productoIndividual(fila int, column int, obj Result, tiempo time.Time) {
	for idx := 0; idx < SIZE; idx++ {
		obj.final[fila][column] += obj.matriz_1[fila][idx] * obj.matriz_2[idx][column]
	}
	defer func() {
		recorrido := time.Since(tiempo).Milliseconds()
		tiempoFinal += recorrido
	}()
}
