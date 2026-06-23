package main

import (
	"fmt"
)

type Docente struct {
	ID          int
	Nombre      string
	Competencia string
	Nivel       int
}

type Asignacion struct {
	Docente Docente
	Horas   int
}

// FILTER
func Filter[T any](collection []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// MAP
func Map[T any, U any](collection []T, transform func(T) U) []U {
	result := make([]U, len(collection))
	for i, item := range collection {
		result[i] = transform(item)
	}
	return result
}

// REDUCE
func Reduce[T any, U any](collection []T, initial U, accumulator func(U, T) U) U {
	result := initial
	for _, item := range collection {
		result = accumulator(result, item)
	}
	return result
}

// Filtrar docentes con nivel mínimo
func FiltrarPorNivel(minNivel int) func(Docente) bool {
	return func(d Docente) bool {
		return d.Nivel >= minNivel
	}
}

// Incrementar nivel de competencia
func IncrementarNivel(incremento int) func(Docente) Docente {
	return func(d Docente) Docente {
		return Docente{
			ID:          d.ID,
			Nombre:      d.Nombre,
			Competencia: d.Competencia,
			Nivel:       d.Nivel + incremento,
		}
	}
}

func main() {

	docentes := []Docente{
		{ID: 1, Nombre: "Juan Perez", Competencia: "Investigacion", Nivel: 4},
		{ID: 2, Nombre: "Maria Lopez", Competencia: "Liderazgo", Nivel: 2},
		{ID: 3, Nombre: "Carlos Sanchez", Competencia: "Innovacion Educativa", Nivel: 5},
		{ID: 4, Nombre: "Ana Torres", Competencia: "Investigacion", Nivel: 3},
	}

	fmt.Println("=== 1. LISTA DE DOCENTES ===")
	for _, d := range docentes {
		fmt.Printf("- %s | Competencia: %s | Nivel: %d\n",
			d.Nombre,
			d.Competencia,
			d.Nivel)
	}

	// FILTER
	docentesDestacados := Filter(docentes, FiltrarPorNivel(4))

	fmt.Println("\n=== 2. FILTER (NIVEL MAYOR O IGUAL A 4) ===")
	for _, d := range docentesDestacados {
		fmt.Printf("- %s | Nivel: %d\n",
			d.Nombre,
			d.Nivel)
	}

	// MAP
	docentesActualizados := Map(docentesDestacados, IncrementarNivel(1))

	fmt.Println("\n=== 3. MAP (INCREMENTO DE NIVEL) ===")
	for _, d := range docentesActualizados {
		fmt.Printf("- %s | Nuevo Nivel: %d\n",
			d.Nombre,
			d.Nivel)
	}

	// Datos para REDUCE
	asignaciones := []Asignacion{
		{Docente: docentesActualizados[0], Horas: 20},
		{Docente: docentesActualizados[1], Horas: 15},
	}

	totalHoras := Reduce(asignaciones, 0, func(acumulador int, item Asignacion) int {
		return acumulador + item.Horas
	})

	fmt.Println("\n=== 4. REDUCE (TOTAL DE HORAS ACADÉMICAS) ===")
	fmt.Printf("Total de horas asignadas: %d\n", totalHoras)

	// Inmutabilidad
	fmt.Println("\n=== 5. COMPROBACIÓN DE INMUTABILIDAD ===")
	fmt.Printf("Nivel original de %s: %d (permanece intacto)\n",
		docentes[0].Nombre,
		docentes[0].Nivel)
}