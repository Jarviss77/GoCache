package main

import (
	"flag"         // Package flag implements command-line flag parsing.
	"fmt"          
	"io"           
	"log"          
	"net"          // Package net provides a portable interface for network I/O.
	"os"           // Package os provides a platform-independent interface to operating system functionality.
	"strings"      
	"github.com/pkg/errors"  // Package errors implements functions to manipulate errors.
)

var (
	listen = flag.String("listen", ":6379", "address to listen to")  // Command-line flag to specify the address to listen on.
)

func main() {
	flag.Parse()  // Parse the command-line flags.
	err := run() 
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)  
		os.Exit(1) 
	}
}

func run() (err error) {
	l, err := net.Listen("tcp", *listen)  // Listen for incoming TCP connections on the address specified by the command-line flag.
	if err != nil { 
		return errors.Wrap(err, "listen")  
	}
	// Defer closing the listener and handle any error during closure
	defer closeIt(l, &err, "close listener")  

	log.Printf("listening %v", l.Addr())  

	// Accept multiple client connections
	for {
		c, err := l.Accept()  
		if err != nil {
			return errors.Wrap(err, "accept") 
		}
		go handleConnections(c)
	}
}

// This handleConnections function reads a command from the client and writes a response.
func handleConnections(c net.Conn) {
	var commands = map[string]func([]Value) ReturnValue{
		"PING": ping,
		"ECHO": echo,
		"SET": set,
		"GET": get,
	}

	for{
		buf := make([]byte, 1024) 
		_, err := c.Read(buf)
		if err != nil {
			log.Printf("error: %v", errors.Wrap(err, "read command"))  
			return
		}

		input, err := parseCommand(buf)
		log.Printf("input: %v", input)

		if err != nil {
			log.Printf("error: %v", errors.Wrap(err, "parse command"))
			return
		}

		trimmedCommand := strings.TrimSpace(input.String())

		args := strings.Split(trimmedCommand, " ")

		cmd := strings.ToUpper(args[0])
		
		if handler, exists := commands[cmd]; exists {
			var cmdArgs []Value

			cmdArgs = append(cmdArgs, Value{str: args})

			response := handler(cmdArgs)
			_, err = c.Write([]byte(response.str + "\r\n"))
			if err != nil {
				log.Printf("error: %v", errors.Wrap(err, "write response"))  
				return
			}
		} else {
			_, err = c.Write([]byte("-ERR unknown command '" + trimmedCommand + "'\r\n"))
			if err != nil {
				log.Printf("error: %v", errors.Wrap(err, "write response"))  
				return
			}
		}
		
	}

}

func closeIt(c io.Closer, errp *error, msg string) {
	err := c.Close() 
	if *errp == nil {
		*errp = errors.Wrap(err, "%v") 
		log.Printf("%s", msg) 
	}
}
