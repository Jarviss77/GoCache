// +build server

package main

import (
	"flag"         // Package flag implements command-line flag parsing.
	"fmt"          // Package fmt implements formatted I/O.
	"io"           // Package io provides basic interfaces to I/O primitives.
	"log"          // Package log implements a simple logging package.
	"net"          // Package net provides a portable interface for network I/O.
	"os"           // Package os provides a platform-independent interface to operating system functionality.
	"github.com/pkg/errors"  // Package errors implements functions to manipulate errors.
)

var (
	listen = flag.String("listen", ":6379", "address to listen to")  // Command-line flag to specify the address to listen on.
)

func main() {
	flag.Parse()  // Parse the command-line flags.
	err := run()  // Call the run function and capture any returned error.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)  // Print error message to standard error.
		os.Exit(1)  // Exit the program with a non-zero status code indicating failure.
	}
}

func run() (err error) {
	l, err := net.Listen("tcp", *listen)  // Listen for incoming TCP connections on the address specified by the command-line flag.
	if err != nil {
		return errors.Wrap(err, "listen")  // Wrap and return the error with context if Listen fails.
	}
	defer closeIt(l, &err, "close listener")  // Defer closing the listener and handle any error during closure.

	log.Printf("listening %v", l.Addr())  // Log the address the server is listening on.

	_, err = l.Accept()  // Accept waits for and returns the next connection to the listener.
	c, err := l.Accept()  // Accept another connection from the listener.
	if err != nil {
		return errors.Wrap(err, "accept")  // Wrap and return the error with context if Accept fails.
	}
	defer closeIt(c, &err, "close connection")  // Defer closing the connection and handle any error during closure.

	buf := make([]byte, 128)  // Create a buffer with a capacity of 128 bytes.
	_, err = c.Read(buf)  // Read command data into the buffer from the connection.
	if err != nil {
		return errors.Wrap(err, "read command")  // Wrap and return the error with context if Read fails.
	}
	log.Printf("read command:\n%s", buf)  // Log the command read from the connection.

	_, err = c.Write([]byte("+PONG\r\n"))  // Write the response "+PONG\r\n" to the connection.
	if err != nil {
		return errors.Wrap(err, "write response")  // Wrap and return the error with context if Write fails.
	}

	return nil  // Return nil indicating success.
}

func closeIt(c io.Closer, errp *error, msg string) {
	err := c.Close()  // Close the closer (listener or connection).
	if *errp == nil {
		*errp = errors.Wrap(err, "%v")  // Wrap and assign the error if errp is not yet set.
	}
}
