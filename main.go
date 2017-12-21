package main

import "fmt"

type Nodo struct {
	Valor int
	Izq   *Nodo
	Der   *Nodo
}

var raiz *Nodo

func main() {
	var nodo Nodo

	for true {
		fmt.Println("--------------------------")
		fmt.Println("Iniciando...")
		fmt.Print("Valor: ")
		fmt.Scanf("%d", &nodo.Valor)
		fmt.Printf("%s\n", insertar(&nodo))
		fmt.Println("--------------------------")
	}
}

/*func inicializarRaiz(v int) (valorRaiz int) {
	raiz = Nodo{Valor: v}
	valorRaiz = raiz.Valor
	return
}*/

func insertar(nodo *Nodo) (mensaje string) {
	if raiz == nil {
		mensaje = "Agregado uno nuevo en la raíz..."
		raiz = nodo
	} else {
		if nodo.Valor == raiz.Valor {
			fmt.Println("Nodo: ", nodo.Valor)
			fmt.Println("Raíz: ", raiz.Valor)
			mensaje = "No se pueden repetir números"
		} else if nodo.Valor < raiz.Valor {
			mensaje = "Agregado un número a la izquierda"
			insertar(nodo.Izq)
		} else {
			mensaje = "Agregado un número a la derecha"
			insertar(nodo.Der)
		}
	}
	return
}
