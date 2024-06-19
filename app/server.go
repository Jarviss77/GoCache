package main

import (
	"flag"         // Package flag implements command-line flag parsing.
	"fmt"          
	"io"           
	"log"          
	"net"          // Package net provides a portable interface for network I/O.
	"os"           // Package os provides a platform-independent interface to operating system functionality.
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

	for {
		c, err := l.Accept()  
		if err != nil {
			return errors.Wrap(err, "accept") 
		}
		go handleConnections(c)
	}
}

func handleConnections(c net.Conn) {

	for{
		buf := make([]byte, 128)  
		_, err := c.Read(buf)  
		if err != nil {
			log.Printf("error: %v", errors.Wrap(err, "read command"))  
			return
		}
		
		log.Printf("read command:\n%s", buf)  
		_, err = c.Write([]byte("+PONG\r\n"))  
		if err != nil {
			log.Printf("error: %v", errors.Wrap(err, "write response"))  
			return
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
