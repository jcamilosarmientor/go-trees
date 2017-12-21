package structs

type Nodo struct {
	Valor *int
	Izq   *Nodo
	Der   *Nodo
}
