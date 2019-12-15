package main

import (
	"net"
	"strings"
)
import "fmt"
import "bufio"
import "os"

func main() {

	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(text)) == "exit" {
			fmt.Println("client usnul")
			return
		} else {
			fmt.Fprintf(conn, text+"\n")
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Print("Message from server: " + message)
		}
	}
}
