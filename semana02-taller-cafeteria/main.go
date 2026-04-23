package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}

func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1
}

func ListarClientes(clientes []Cliente) {
	fmt.Println("\n==== LISTADO DE CLIENTES ====")
	if len(clientes) == 0 {
		fmt.Println("(no hay clientes registrados)")
		return
	}

	fmt.Printf("%-5s | %-20s | %-12s | %-10s\n", "ID", "Nombre", "Carrera", "Saldo")
	fmt.Println("------------------------------------------------------------")
	for _, c := range clientes {
		fmt.Printf("%-5d | %-20s | %-12s | $%.2f\n", c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarClientePorID(clientes, id)
	if idx == -1 {
		fmt.Println("Error: Cliente no encontrado")
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, p := range productos {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func ListarProductos(productos []Producto) {
	fmt.Println("\n==== PRODUCTOS DISPONIBLES ====")
	if len(productos) == 0 {
		fmt.Println("(no hay productos)")
		return
	}
	fmt.Printf("%-5s | %-15s | %-10s |   %-8s | %-10s\n", "ID", "Nombre", "Precio", "Stock", "Categoría")
	fmt.Println("----------------------------------------------------------------------")
	for _, p := range productos {
		fmt.Printf("%-5d | %-15s | $%-9.2f | %-8d | %-10s\n", p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}
}

func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
	if idx == -1 {
		fmt.Println("Error: Producto no encontrado")
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

func DescontarSaldo(cliente *Cliente, monto float64) error {
	if monto < 0 {
		return fmt.Errorf("el monto no puede ser negativo")
	}
	if cliente.Saldo < monto {
		return fmt.Errorf("saldo insuficiente: tiene $%.2f, necesita $%.2f", cliente.Saldo, monto)
	}
	cliente.Saldo -= monto
	return nil
}

func DescontarStock(producto *Producto, cantidad int) error {
	if cantidad <= 0 {
		return fmt.Errorf("la cantidad debe ser mayor a cero")
	}
	if producto.Stock < cantidad {
		return fmt.Errorf("stock insuficiente: solo hay %d unidades", producto.Stock)
	}
	producto.Stock -= cantidad
	return nil
}

func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
	fecha string,
) ([]Pedido, error) {
	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		return pedidos, fmt.Errorf("cliente no encontrado")
	}

	idxP := BuscarProductoPorID(productos, productoID)
	if idxP == -1 {
		return pedidos, fmt.Errorf("producto no encontrado")
	}

	total := productos[idxP].Precio * float64(cantidad)

	err := DescontarStock(&productos[idxP], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarSaldo(&clientes[idxC], total)
	if err != nil {
		productos[idxP].Stock += cantidad
		return pedidos, err
	}

	nuevoPedido := Pedido{
		ID:         len(pedidos) + 1,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}

	pedidos = append(pedidos, nuevoPedido)
	return pedidos, nil
}

func leerLinea(lector *bufio.Reader) string {
	linea, _ := lector.ReadString('\n')
	return strings.TrimSpace(linea)
}

func leerEntero(lector *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	n, err := strconv.Atoi(texto)
	if err != nil {
		return -1
	}
	return n
}

func leerFloat(lector *bufio.Reader, prompt string) float64 {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	f, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		return -1
	}
	return f
}

func PedidosDeCliente(pedidos []Pedido, clientes []Cliente, productos []Producto, clienteID int) {
	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		fmt.Println("Error: Cliente no existe.")
		return
	}

	fmt.Printf("\nREPORTE DE PEDIDOS: %s\n", clientes[idxC].Nombre)
	fmt.Printf("%-5s | %-15s | %-8s | %-8s | %-10s\n", "ID", "Producto", "Cant.", "Total", "Fecha")
	fmt.Println("------------------------------------------------------------------")

	totalGastado := 0.0
	encontrado := false

	for _, p := range pedidos {
		if p.ClienteID == clienteID {
			encontrado = true
			idxP := BuscarProductoPorID(productos, p.ProductoID)
			nombreProd := "Desconocido"
			if idxP != -1 {
				nombreProd = productos[idxP].Nombre
			}
			fmt.Printf("%-5d | %-15s | %-8d | $%-7.2f | %-10s\n", p.ID, nombreProd, p.Cantidad, p.Total, p.Fecha)
			totalGastado += p.Total
		}
	}

	if !encontrado {
		fmt.Println("(El cliente no tiene pedidos registrados)")
	} else {
		fmt.Printf("\nTOTAL ACUMULADO: $%.2f\n", totalGastado)
	}
}

func main() {
	// --- MOVIDO AL PRINCIPIO PARA EVITAR ERROR 'UNDEFINED' ---
	pedidos := []Pedido{}
	lector := bufio.NewReader(os.Stdin)

	clientes := []Cliente{
		{ID: 1, Nombre: "John Bello", Carrera: "TI", Saldo: 155.06},
		{ID: 2, Nombre: "Juan Lopez", Carrera: "Ing. Civil", Saldo: 160.55},
		{ID: 3, Nombre: "Lady Lucas", Carrera: "Medicina", Saldo: 180.50},
	}

	productos := []Producto{
		{ID: 1, Nombre: "Cafe", Precio: 1.50, Stock: 15, Categoria: "Bebida"},
		{ID: 2, Nombre: "Sanduche", Precio: 2.25, Stock: 30, Categoria: "Snack"},
		{ID: 3, Nombre: "Torta", Precio: 3.06, Stock: 10, Categoria: "Postre"},
		{ID: 4, Nombre: "Agua", Precio: 0.50, Stock: 50, Categoria: "Bebida"},
	}

	fmt.Println("==== DEMOSTRACIÓN DE REGISTRO DE PEDIDOS ====")

	fmt.Println("\n1. Intentando pedido exitoso: John (ID 1) compra 2 Cafés (ID 1)...")
	var err error
	pedidos, err = RegistrarPedido(clientes, productos, pedidos, 1, 1, 2, "2026-04-16")
	if err != nil {
		fmt.Println("Error inesperado:", err)
	} else {
		fmt.Println("¡Pedido registrado con éxito!")
	}

	fmt.Println("\n2. Intentando comprar 100 Sanduches (ID 2) cuando solo hay 30...")
	pedidos, err = RegistrarPedido(clientes, productos, pedidos, 1, 2, 100, "2026-04-16")
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	ListarClientes(clientes)
	ListarProductos(productos)

	fmt.Printf("\nTotal de pedidos en memoria: %d\n", len(pedidos))

	for {
		fmt.Println("\n--- MINI-CAFETERÍA UNIVERSITARIA ---")
		fmt.Println("1. Listar clientes")
		fmt.Println("2. Listar productos")
		fmt.Println("3. Agregar cliente")
		fmt.Println("4. Agregar producto")
		fmt.Println("5. Registrar pedido")
		fmt.Println("6. Ver pedidos de un cliente")
		fmt.Println("0. Salir")

		opcion := leerEntero(lector, "Seleccione una opción: ")

		switch opcion {
		case 1:
			ListarClientes(clientes)
		case 2:
			ListarProductos(productos)
		case 3:
			fmt.Println("\n--- NUEVO CLIENTE ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Carrera: ")
			carrera := leerLinea(lector)
			saldo := leerFloat(lector, "Saldo inicial: ")
			clientes = AgregarCliente(clientes, Cliente{id, nombre, carrera, saldo})
		case 4:
			fmt.Println("\n--- NUEVO PRODUCTO ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			precio := leerFloat(lector, "Precio: ")
			stock := leerEntero(lector, "Stock: ")
			fmt.Print("Categoría: ")
			cat := leerLinea(lector)
			productos = AgregarProducto(productos, Producto{id, nombre, precio, stock, cat})
		case 5:
			cID := leerEntero(lector, "ID Cliente: ")
			pID := leerEntero(lector, "ID Producto: ")
			cant := leerEntero(lector, "Cantidad: ")
			fmt.Print("Fecha (YYYY-MM-DD): ")
			fecha := leerLinea(lector)

			var err error
			pedidos, err = RegistrarPedido(clientes, productos, pedidos, cID, pID, cant, fecha)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("¡Pedido exitoso!")
			}
		case 6:
			cID := leerEntero(lector, "ID del cliente para el reporte: ")
			PedidosDeCliente(pedidos, clientes, productos, cID)
		case 0:
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
