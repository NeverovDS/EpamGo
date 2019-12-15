package main

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)
import "fmt"
import "bufio"

func main() {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Message Received:", string(message))
		// Процесс выборки для полученной строки
		message = strings.TrimSpace(string(message))
		r, err1 := regexp.MatchString(`^[0-9]+$`, message)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		if r == true {
			newMessage, _ := strconv.Atoi(message)
			newMessage *= 2
			conn.Write([]byte(strconv.Itoa(newMessage) + "\n"))
		}
		if r == false {
			message = strings.ToUpper(message)
			conn.Write([]byte(string(message) + "\n"))
		}

	}
}
