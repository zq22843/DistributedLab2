package main

import (
	"bufio"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"net/rpc"
	"flag"
	"net/rpc"
	"os"
	//	"bufio"
	//	"os"
	//	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	"fmt"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	file, _ := os.Open("wordlist")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//Print the line
		request := stubs.Request{Message: scanner.Text()}
		response := new(stubs.Response)
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responded: " + response.Message)
	}

	//TODO: connect to the RPC server and send the request(s)
}
