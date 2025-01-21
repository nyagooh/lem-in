package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var (
		rooms []string
		// ants int
		tunnels                []string
		bfs                    []string
		startRoom, starttunnel int
		end, start             int
	)

	for scanner.Scan() {
		line := scanner.Text()
		bfs = append(bfs, line)
	}
	// ants, _ = strconv.Atoi(strings.TrimSpace(bfs[0]))

	for i, ch := range bfs {
		if ch == "##start" {
			startant1 := bfs[i+1]
			startRoom, _ = strconv.Atoi(string(startant1[0]))
			start = i + 2
		} else if ch == "##end" {
			starttunnel1 := bfs[i+1]
			starttunnel, _ = strconv.Atoi(string(starttunnel1[0]))
			end = i
		}

	}
	rooms = append(rooms, bfs[start:end]...)
	tunnels = append(tunnels, bfs[end+2:]...)
	// fbfs(starttunnel, endtunnel,)

	graph := createGraph(rooms, tunnels)
	path := BFS(graph, startRoom, starttunnel)
	fmt.Println(path[0])

}

//this graph takes in the rooms and the connections to the room
type Graph map[int][]int

func createGraph(rooms []string, tunnels []string) Graph {
	//reference the type of the graph.so the graph created is a map that has keys of type int and values of type []int
	graph := make(Graph)

	// fromm data give each room is rep by three values that is wahy we increment by 3
	for i := 0; i < len(rooms); i += 3 {
		node, _ := strconv.Atoi(rooms[i])
		graph[node] = []int{}
	}

	// each tunnel is a string like 0-1.meaning room 0 connect to room 1.We using same graph because the tunnels represent connection between rooms
	for _, tunnel := range tunnels {
		parts := strings.Split(tunnel, "-")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	return graph
}
func BFS(graph Graph, start, end int) [][]int {
	var paths [][]int
	queue := [][]int{{start}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]
		if node == end {
			paths = append(paths, path)
			continue
		}

		for _, neighbor := range graph[node] {
			if !contain(path, neighbor) {
				newPath := append([]int{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}
func contain(path []int, node int) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}