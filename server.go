package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
)

// struct for storing each record
type entry struct {
	key   []byte
	value []byte
}

func main() {
	listener, err := net.Listen("tcp", os.Args[1])
	checkError(err)

	defer listener.Close()
	log.Printf("Ready to listen on %s", os.Args[1])

	for {
		conn, err := listener.Accept()
		log.Printf("Accepted a new connection")
		checkError(err)

		var data = make([]entry, 0)
		c := make(chan entry)

		go processClient(conn, c)

		fmt.Printf("Channel content: ")
		for entry := range c {
			data = append(data, entry)
			fmt.Printf("%s,", hex.EncodeToString(entry.key))
		}
		logData(data)
		writeResponse(conn, data)
	}
}

func processClient(conn net.Conn, c chan entry) {
	// separating key and value pairs
	k := make([]byte, 10)
	v := make([]byte, 90)

	for {
		n, err := conn.Read(k[0:])
		checkError(err)
		log.Printf("Read key, size: %d, %s", n, hex.EncodeToString(k[0:n]))
		// check if number of bytes read is 10, otherwise it means EOF
		if n != 10 {
			close(c)
			break
		}
		n, err = conn.Read(v)
		checkError(err)
		log.Printf("Read value, size: %d", n)

		// append to data
		key := k
		record := entry{key, v}
		c <- record
	}
}

func sortData(data []entry) []entry {
	// sort the entries
	sort.SliceStable(data, func(i, j int) bool {
		res := bytes.Compare(data[i].key, data[j].key)
		return res == -1
	})
	return data
}

func writeResponse(conn net.Conn, data []entry) {
	for _, entry := range data {
		_, err := conn.Write(entry.key)
		checkError(err)
		_, err = conn.Write(entry.value)
		checkError(err)
	}
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}

func logData(data []entry) {
	fmt.Printf("Size: %d ---- [", len(data))
	for _, entry := range data {
		fmt.Printf("%s ", hex.EncodeToString(entry.key))
	}
	fmt.Printf("]")
}
