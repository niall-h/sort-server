package main

import (
	"bytes"
	"encoding/hex"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: go run client.go <host>:<port> <inputfile> <outputfile>")
	}

	// get host:port
	service := os.Args[1]
	inputFile := os.Args[2]

	// input file
	file, err := os.Open(inputFile)
	checkError(err)
	defer file.Close()

	// attempt connection to service
	conn, err := net.Dial("tcp", service)
	checkError(err)
	defer conn.Close()

	// send data in 100 byte chunks
	for {
		buf := make([]byte, 100)
		_, err := file.Read(buf)
		log.Printf("Reading 100 bytes from file: %s", string(buf))
		if err != nil {
			if err == io.EOF {
				// send ending bit
				log.Printf("Signalling end of file: writing byte 0\n")
				conn.Write([]byte{0})
				break
			}
			log.Fatalf("Fatal error: %s", err.Error())
			return
		}

		// send data across tcp connection
		log.Printf("Writing to tcp connection")
		_, err = conn.Write(buf)
		checkError(err)
	}

	// read response from server
	result, err := readFully(conn)
	checkError(err)

	writeToFile(result)
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	for {
		buf := make([]byte, 100)
		n, err := conn.Read(buf[0:])
		log.Printf("Reading response, %d bytes, %s\n", n, hex.EncodeToString(buf[0:n]))
		result.Write(buf[0:n])

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func writeToFile(data []byte) {
	// create output file
	file, err := os.Create(os.Args[3])
	checkError(err)
	file.Write(data)
	file.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
