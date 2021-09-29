package main

import "fmt" //Area donde se importan las librerias encesrias para ejecutar el programa

func privada() {
	fmt.Println("Ejecutar alguna lógica que no requiere ser exportada")
}

func Publica() {
	fmt.Println("Ejecutar alguna lógica que pretenfo exportar")
}

func saludar() {
	defer fmt.Println("Buenas tardes") //Esta porción de código debe ser ejecutado hasta que la función haya terminado.
	fmt.Println("¿Cómo se encuentra?")
}

func main() {

	var message string = "Que se dice parcero"
	easyMessage := "Que se dice parcero usando :="
	fmt.Println(message)
	fmt.Println(easyMessage)

	//Comentarios !!!
	a := 10.
	var b float64 = 3
	fmt.Println((a / b))

	var c int = 10
	d := 3
	fmt.Println(c / d)
	x := true
	y := false
	fmt.Println((x || y))
	fmt.Println(y && x)
	fmt.Println(!x)
	privada()
	Publica()
	saludar()

}
