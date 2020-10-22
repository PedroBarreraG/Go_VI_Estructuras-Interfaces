package main

import (
	"fmt"

	"./multimedias"
)

func mostrar(multi ...multimedias.Multimedia) {
	for _, m := range multi {
		m.Mostrar()
	}
}

func main() {
	op := 1
	var Multimedias []multimedias.Multimedia
	for op != 0 {
		fmt.Print("1.Capturar Imagen\n2.Capturar Audio\n3.Capturar Video\n4.Mostrar Slice\n0.Salir\nOpcion: ")
		fmt.Scan(&op)
		switch op {
		case 1:
			var titulo string
			fmt.Print("Titulo: ")
			fmt.Scan(&titulo)
			var formato string
			fmt.Print("Formato: ")
			fmt.Scan(&formato)
			var canales string
			fmt.Print("Canales: ")
			fmt.Scan(&canales)
			mulIma := multimedias.Imagen{Titulo: titulo, Formato: formato, Canales: canales}
			Multimedias = append(Multimedias, &mulIma)
		case 2:
			var titulo string
			fmt.Print("Titulo: ")
			fmt.Scan(&titulo)
			var formato string
			fmt.Print("Formato: ")
			fmt.Scan(&formato)
			var duracion string
			fmt.Print("Duracion: ")
			fmt.Scan(&duracion)
			mulAud := multimedias.Audio{Titulo: titulo, Formato: formato, Duracion: duracion}
			Multimedias = append(Multimedias, &mulAud)
		case 3:
			var titulo string
			fmt.Print("Titulo: ")
			fmt.Scan(&titulo)
			var formato string
			fmt.Print("Formato: ")
			fmt.Scan(&formato)
			var frames string
			fmt.Print("Frames: ")
			fmt.Scan(&frames)
			mulAud := multimedias.Video{Titulo: titulo, Formato: formato, Frames: frames}
			Multimedias = append(Multimedias, &mulAud)
		case 4:
			fm := multimedias.ContenidoWeb{
				Multimedias,
			}
			fm.Mostrar()
		}
	}
}
