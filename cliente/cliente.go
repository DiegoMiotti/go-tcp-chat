package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error al intentar conectar al servidor:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	// Obtener nickname del usuario
	fmt.Print("Ingrese su nickname: ")
	nickname, _ := reader.ReadString('\n')
	nickname = strings.TrimSpace(nickname)
	conn.Write([]byte(nickname + "\n"))
	fmt.Println("Nickname enviado:", nickname) // Debug

	go readMessages(conn)

	for {
		fmt.Print("Ingrese mensaje: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error al enviar el mensaje:", err)
			return
		}
	}
}

func readMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Se ha desconectado")
			return
		}
		fmt.Print("Nuevo mensaje del servidor: " + message)
	}
}
