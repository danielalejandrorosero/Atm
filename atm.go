package main

import (
	"fmt"

	"github.com/danielalejandrorosero/atm/clear"
)

var balance float32 = 0
var otraTranscacion int

func transacion() {
	var eleccion int

	fmt.Printf("\n Inicia la opcion a escoger!\n\n")
	fmt.Printf("1. Retirar\n")
	fmt.Printf("2. Depositar\n")
	fmt.Printf("3. Balance\n\n")
	fmt.Scan(&eleccion)

	var montoRetirar float32
	var montoDepositar float32

	switch eleccion {
	case 1:
		fmt.Printf("\n porfavor indique el monto a retirar: ")
		fmt.Scan(&montoRetirar)

		if montoRetirar > balance {
			fmt.Printf("no tienes suficientes fondos en tu cuenta")
			fmt.Printf("Quieres hacer otra transacion?\nPresiona 1 para continuar y 2 para salir")
			fmt.Scan(&otraTranscacion)

			switch otraTranscacion {
			case 1:
				clear.CallClear()
				transacion()
			default:
				fmt.Println("\n gracias por usar nuestro servicio!!! \n que tenga un buen dia")
			}
		} else {
			balance -= montoRetirar
			fmt.Printf("tu has retirado $%.2f y tu nuevo balance es $%.2f", montoRetirar, balance)
			fmt.Println("\n\nQuires hacer otra transacion?\nPresiona 1 si desea continuar y presiona 2 si quieres salir ")
			fmt.Scan(&otraTranscacion)

			switch otraTranscacion {
			case 1:
				clear.CallClear()
				transacion()
			default:
				fmt.Println("\n gracias por usar nuestro servicio!!! \n que tenga un lindo dia")
			}
		}

	case 2:
		fmt.Printf("\n porfavor ingresa el monto  a depositar")
		fmt.Scan(&montoDepositar)

		balance += montoDepositar

		fmt.Printf("gracias por tu deposito, tu nuevo balance es: $%.2f", balance)
		fmt.Printf("\n\nQuieres hacer otra transacion?\nPresiona 1 si deseas continuar presiona 2 si quieres salir")
		fmt.Scan(&otraTranscacion)

		switch otraTranscacion {
		case 1:
			clear.CallClear()
			transacion()
		default:
			fmt.Println("\n gracias por usar nuestro servicio!!! \n que tenga un buen dia")
		}

	case 3:
		fmt.Printf("\n tu balance en el banco es: $%.2f", balance)
		fmt.Printf("\n\nQuires hacer otra transacion?\nPresiona 1 si deseas continuar presiona 2 si quieres salir")
		fmt.Scan(&otraTranscacion)

		switch otraTranscacion {
		case 1:
			clear.CallClear()
			transacion()

		default:
			fmt.Println("\n gracias por usar nuestro servicio!!! \n que tenga un buen dia")
		}
	}
}

func main() {
	fmt.Println("\n bienvenido a mi atm simulator\n")
	transacion()
}
