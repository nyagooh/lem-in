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
		tunnels            []string
		bfs                []string
		startRoom            string
		endRoom               string
	)

	for scanner.Scan() {
		line := scanner.Text()
		bfs = append(bfs, line)
	}
	// ants, _ = strconv.Atoi(strings.TrimSpace(bfs[0]))
	// ant, _ = strconv.Atoi(bfs[0])
	i := 1
	for _, ch := range bfs {
		if ch == "##start" {
			startRoom = bfs[i+1]
			continue
		}
		if ch == "##end" {
			endRoom = bfs[i+1]
			continue
		}
		if strings.Contains(ch, "-") {
			slice := strings.Split(ch, "-")
			tunnels = append(tunnels, slice[0])
			tunnels = append(tunnels, slice[0])
		}else {
			rooms = append(rooms, ch)
		}
		i++
	}

	graph := createGraph(rooms, tunnels)
	// fmt.Println(graph)
	path := BFS(graph, startRoom, starttunnel)
	// fmt.Println(path)

	fmt.Println(CollidingPaths(path))
	fmt.Println(ant)

}

// this graph takes in the rooms and the connections to the room
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

func CollidingPaths(paths [][]int) [][]int {
	i := 0
	for i < len(paths) {
		pathsToRemove := make(map[int]bool)
		for j := i + 1; j < len(paths); j++ {
			if FindCollisions(paths[i], paths[j]) {
				pathsToRemove[j] = true
			}
		}
		paths = RemovePaths(paths, pathsToRemove)
		i++
	}
	return paths
}
func FindCollisions(path1, path2 []int) bool {
	for k := 1; k < len(path1)-1; k++ {
		for l := 1; l < len(path2)-1; l++ {
			if path1[k] == path2[l] {
				return true
			}
		}
	}
	return false
}
func RemovePaths(paths [][]int, pathsToRemove map[int]bool) [][]int {
	updatedPaths := [][]int{}
	for idx, path := range paths {
		if !pathsToRemove[idx] {
			updatedPaths = append(updatedPaths, path)
		}
	}
	return updatedPaths
}
