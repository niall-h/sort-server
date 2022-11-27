Sort Server (Socket Programming)

The idea of this application is to establish a client-server connection and perform a sort operation on the server. The client parses a binary file as input, attempts to connect to a server, and sends the data obtained from the file to the server using a TCP socket. The server receives this data, sorts it, and writes a response containing the sorted binary data back to the client. The client then writes this data back to a local output file.

__________________________________________________________________________________________________________________________________________

Specifications:
- The input file holds zero or more records. Each record is a 100 byte binary key-value pair, consisting of a 10-byte key and a 90-byte value.
- The sort program sorts the records in ascending order based on the keys.
- The files are guranteed to be in multiples of 100 bytes.

__________________________________________________________________________________________________________________________________________

Usage:

To start the server:
`go run server.go <host>:<port>`

To start the client:
`go run client.go <host>:<port> <inputfile> <outputfile>`

Useful utilities (adapted from George Porter's CSE124 class project):

Gensort

Gensort generates random input. If the 'randseed' parameter is provided, the given seed is used to ensure deterministic output.
'size' can be provided as a non-negative integer to generate that many bytes of output. However human-readable strings can be used as well, such as "10 mb" for 10 megabytes, "1 gb" for one gigabyte", "256 kb" for 256 kilobytes, etc. You can specify the input size in any format supported by the https://github.com/c2h5oh/datasize package.
If the specified size is not a multiple of 100 bytes, the requested size will be rounded up to the next multiple of 100.

Usage:
`./utils/<your-architecture>/bin/gensort <outputfile> <size>`

Showsort

Showsort shows the records in the provided file in a human-readable format, with the key followed by a space followed by an abbreviated version of the value.

Usage: 
`./utils/<your-architecture>/bin/showsort <inputfile>`

Valsort
Valsort scans the provided input file to check if it is sorted. Note that valsort does not verify that the set of key/value pairs in the provided file matches the set from gensort–it only checks that the key-value pairs in the provided file are sorted.

Usage: 
`./utils/<your-architecture>/bin/valsort <inputfile>`
