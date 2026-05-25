package client

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func processInput(reader *bufio.Reader) {
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "exit" {
		log.Println("Exit command received. Stopping client...")
		os.Exit(0)
	}
	inCh <- input
}