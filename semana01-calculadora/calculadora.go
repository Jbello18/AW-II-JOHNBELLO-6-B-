package main

import "fmt"

func main() {
	var historial string
	var contador int

	for {
		fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

		var n1, n2 float64
		var operacion string

		fmt.Print("Ingrese el primer número: ")
		fmt.Scanln(&n1)
		fmt.Print("Ingrese el segundo número: ")
		fmt.Scanln(&n2)

		fmt.Print("Ingrese la operación (+, -, *, /, ^, !): ")
		fmt.Scanln(&operacion)

		var resultadoActual string

		switch operacion {
		case "+":
			resultadoActual = fmt.Sprintf("%.2f + %.2f = %.2f", n1, n2, n1+n2)
		case "-":
			resultadoActual = fmt.Sprintf("%.2f - %.2f = %.2f", n1, n2, n1-n2)
		case "*":
			resultadoActual = fmt.Sprintf("%.2f * %.2f = %.2f", n1, n2, n1*n2)
		case "/":
			if n2 == 0 {
				fmt.Println("Error: no se puede dividir entre cero")
				continue
			}
			resultadoActual = fmt.Sprintf("%.2f / %.2f = %.2f", n1, n2, n1/n2)
		case "^":
			resPot := 1.0
			for i := 0; i < int(n2); i++ {
				resPot *= n1
			}
			resultadoActual = fmt.Sprintf("%.2f ^ %.0f = %.2f", n1, n2, resPot)
		case "!":
			if n1 < 0 {
				fmt.Println("Error: No existe factorial de números negativos.")
				continue
			}
			resFact := 1.0
			for i := 1; i <= int(n1); i++ {
				resFact *= float64(i)
			}
			resultadoActual = fmt.Sprintf("%.0f! = %.2f", n1, resFact)
		default:
			fmt.Println("Error: operación no reconocida")
			continue
		}

		fmt.Println("Resultado:", resultadoActual)

		// Guardar el historial
		contador++
		historial += fmt.Sprintf("%d) %s\n", contador, resultadoActual)

		var respuesta string
		fmt.Print("¿Desea realizar otra operación? (s/n): ")
		fmt.Scanln(&respuesta)

		if respuesta == "n" {
			fmt.Println("\n==== HISTORIAL DE OPERACIONES ====")
			fmt.Print(historial)
			fmt.Printf("Total de operaciones realizadas: %d\n", contador)
			fmt.Println("¡Hasta luego!")
			break
		}
	}
}
