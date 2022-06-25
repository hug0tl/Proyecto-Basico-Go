package main

import (
	// Imprimir datos en consola
	"fmt"

	//Importando paquetes para leer desde consola
	"bufio"
	"os"

	// para eliminar salto de linea
	"strings"

	// Convertir string a entero
	"strconv"

	// Paquete para ejecutar instrucciones en consola: Limpiar consola
	"os/exec"
)

// #### VARIABLES GLOBALES ####
// Variable global para leer en pantalla
var reader *bufio.Reader

// Mapa que contendra a los usuarios
var users map[int]Usuario
var id int

// Creacion de estructura Usuario
type Usuario struct {
	id       int
	userName string
	email    string
	edad     int
}

// Creacion de metodos CRUD para usuario
func crearUsuario() {
	limpiar()

	fmt.Print("Username: ")
	username := readLine()

	fmt.Print("Email: ")
	email := readLine()

	fmt.Print("Edad: ")
	if edad, error := strconv.Atoi(readLine()); error != nil {
		panic("Valor edad no valido")
	} else {
		id++
		user := Usuario{id, username, email, edad}
		users[id] = user
		fmt.Println("Usuario creado exitosamente!")
	}

}
func listarUsuario() {
	limpiar()

	if len(users) == 0 {
		fmt.Println("No hay usuarios que listar")
	} else {
		fmt.Println("USUARIOS:")
		for id, user := range users {
			fmt.Println(id, " - ", user.userName, " - ", user.email, " - ", user.edad)
		}
	}
}
func actualizarUsuario() {

}
func eliminarUsuario() {
	limpiar()

	if len(users) == 0 {
		fmt.Println("No hay usuarios para eliminar")
	} else {
		fmt.Print("Ingresa el id: ")
		if id, error := strconv.Atoi(readLine()); error != nil {
			panic("Id no valido")
		} else {
			if _, ok := users[id]; ok {
				delete(users, id)
				fmt.Println("Usuario eliminado exitosamente")
			} else {
				fmt.Println("Usuario no encontrado")
			}

		}
	}
}
func readLine() string {
	if opcion, error := reader.ReadString('\n'); error != nil {
		panic("No es posible obtener el valor")
	} else {
		return strings.TrimSuffix(opcion, "\n")
	}
}
func limpiar() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func saltos() {
	fmt.Print("\n\n\n")
}

func main() {
	limpiar()
	users = make(map[int]Usuario)
	var opcion string
	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Println("A) Agregar")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")
		fmt.Println("Q) Salir")
		fmt.Println("Elige una opcion: ")
		opcion = readLine()
		if opcion == "q" || opcion == "Q" {
			break
		}

		switch opcion {
		case "a", "A":
			crearUsuario()
		case "b", "B":
			listarUsuario()
		case "c", "C":
			actualizarUsuario()
		case "d", "D":
			eliminarUsuario()
		default:
			limpiar()
			fmt.Println("Opcion invalida")
			saltos()
		}
		saltos()
	}
	limpiar()
	fmt.Println("Gracias por usar la aplicacion")
}
