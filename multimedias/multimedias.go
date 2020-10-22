package multimedias

import "fmt"

//Estructura ContenidoWeb
type ContenidoWeb struct {
	Multimedias []Multimedia
}

func (fm *ContenidoWeb) Mostrar() {
	for _, m := range fm.Multimedias {
		m.Mostrar()
	}
}

//Interface Multimedia
type Multimedia interface {
	Mostrar()
}

//Estructura Imagen
type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

//Estructura Audio
type Audio struct {
	Titulo   string
	Formato  string
	Duracion string
}

//Estructura Video
type Video struct {
	Titulo  string
	Formato string
	Frames  string
}

//Metodo mostrar()
func (i *Imagen) Mostrar() {
	fmt.Println(i.Titulo)
	fmt.Println(i.Formato)
	fmt.Println(i.Canales)
}

//Metodo mostrar()
func (a *Audio) Mostrar() {
	fmt.Println(a.Titulo)
	fmt.Println(a.Formato)
	fmt.Println(a.Duracion)
}

//Metodo mostrar()
func (v *Video) Mostrar() {
	fmt.Println(v.Titulo)
	fmt.Println(v.Formato)
	fmt.Println(v.Frames)
}
