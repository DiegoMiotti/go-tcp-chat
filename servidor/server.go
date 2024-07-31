package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var clients = make(map[net.Conn]string)

func main() {
	fmt.Println("Iniciando servidor...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al intentar aceptar conexión:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Obtener nickname del cliente
	conn.Write([]byte("Ingrese su nickname: "))
	nickname, _ := reader.ReadString('\n')
	nickname = strings.TrimSpace(nickname)
	fmt.Println("Nickname recibido:", nickname) // Debug
	clients[conn] = nickname

	for {
		// Leer mensaje del cliente
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Usuario desconectado:", nickname)
			removeClient(conn)
			return
		}

		fmt.Printf("[%s] %s", nickname, message)

		// Enviar mensaje a todos los clientes conectados
		broadcastMessage(fmt.Sprintf("[%s] %s", nickname, message), conn)
	}
}

func broadcastMessage(message string, sender net.Conn) {
	for client := range clients {
		if client != sender {
			_, err := client.Write([]byte(message))
			if err != nil {
				fmt.Println("Error al enviar el mensaje:", err)
			}
		}
	}
}

func removeClient(conn net.Conn) {
	delete(clients, conn)
}
