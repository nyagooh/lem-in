package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//readfile from the command line
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	var rooms []string
	var tunnels []string
	var bfs []string
	var startant, starttunnel string
	var end, start int
	for scanner.Scan() {
		line := scanner.Text()
		bfs = append(bfs, line)
	}
	for i, ch := range bfs {
		if ch == "##start" {
			startant = bfs[i+1]
			start = i+2
		} else if ch == "##end" {
			starttunnel = bfs[i+1]
			end = i
		}

	}
	rooms = append(rooms, bfs[start:end]...)
	tunnels = append(tunnels, bfs[end+2:]...)
	
}
