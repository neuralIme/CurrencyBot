package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("tokens.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	rd := bufio.NewReader(file)

	tbToken, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	startBot(strings.Trim(tbToken, "\n"))
}
