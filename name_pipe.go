package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fileBasedPipe()
	inMemorySyncPipe()
}

CI_ENABLE_SENTINEL=true
CI_SENTINEL_REDIS_PASS=dcos-redis
CI_SENTINEL_GROUP=redis-master-8000
CI_SENTINEL_NODES=172.24.4.236:18001,172.24.4.236:18002,172.24.4.237:18000

REDIS_SENTINEL_ADDR=172.24.4.236:18001,172.24.4.236:18002,172.24.4.237:18000
REDIS_PORT=6379
REDIS_MODE=sentinel
REDIS_PASSWORD=dcos-redis
REDIS_HOST=172.24.4.111
REDIS_SENTINEL_GROUP=redis-master-8000

func fileBasedPipe() {
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Printf("Error: Couldn't create the named pipe: %s\n", err)
	}
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}

func inMemorySyncPipe() {
	reader, writer := io.Pipe()
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [in-memory pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}
