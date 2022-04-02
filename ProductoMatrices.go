package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup //Espera que una colección de rutinas de Go se ejecuten - PARA LA SEGUNDA PARTE
var SIZE int

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
	//a, b, c := rand.Intn(num)+1, rand.Intn(num)+1, rand.Intn(num)+1
	//Dimensiones de la matriz resultante a - filas matriz1 b - columnas matriz1 y filas matriz2 - c columnas matriz
	a := rand.Intn(num) + 1
	SIZE = a
	generado := Init(a, a, a, a) //Genera una matriz cuadrada siempre, acomodado para otros casos
	return generado, a * a
}

func main() {
	//defer fmt.Println("El producto de las matrices se ha completado") //Al final
	var limite int
	fmt.Print("Ingrese limite maximo de dimensiones ")
	fmt.Scan(&limite)
	productoMatricesClasico(limite)
	wg.Wait()
}

//Resolución de las matrices, a manera básica
func productoMatricesClasico(limit int) *Result {
	matriz, numHilos := generarMatrices(limit)
	wg.Add(numHilos) //Cantidad de hilos independientes que va a esperar
	for idxf, item := range matriz.final {
		for idxc := range item {
			go productoIndividual(idxf, idxc, *matriz) //Se crean hilos individuales para cada elemento de nuestra matriz
		}
	}
	defer fmt.Print(matriz)
	return matriz
}

func productoIndividual(fila int, column int, obj Result) {
	defer fmt.Println("Acabado")
	for idx := 0; idx < SIZE; idx++ {
		obj.final[fila][column] += obj.matriz_1[fila][idx] * obj.matriz_2[idx][column]
	}
	wg.Done() //La rutina ha terminado
}

//Pendiente, administración de threads DIA LUNES