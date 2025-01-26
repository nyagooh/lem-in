package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var (
		rooms []string

		tunnels     []string
		bfs         []string
		startRoom   int
		endRoom     int
		ant         int
		room, room1 string
	)

	for scanner.Scan() {
		line := scanner.Text()
		bfs = append(bfs, line)
	}

	i := 1
	for _, ch := range bfs {
		if ch == bfs[0] {
			ant, _ = strconv.Atoi(ch)
			continue
		}
		if ch == "##start" {
			room1 = bfs[i+1]
			startRoom, _ = strconv.Atoi(string(room1[0]))
			continue
		}
		if ch == "##end" {
			room = bfs[i+2]
			endRoom, _ = strconv.Atoi(string(room[0]))
			continue
		}
		if strings.Contains(ch, "-") {
			tunnels = append(tunnels, ch)
		} else if ch != bfs[0] && ch != room1 && ch != room {
			rooms = append(rooms, ch)
		}

		i++
	}
	graph := createGraph(rooms, tunnels)
	path := BFS(graph, startRoom, endRoom)
	fmt.Println(CollidingPaths(path))
	fmt.Println(Antnames(ant))

}

type Graph map[int][]int

func createGraph(rooms []string, tunnels []string) Graph {
	graph := make(Graph)

	for i := 0; i < len(rooms); i += 3 {
		node, _ := strconv.Atoi(rooms[i])
		graph[node] = []int{}
	}

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

func Antnames(ant int) []string {
	antnames := []string{}
	for i := 1; i <= ant; i++ {
		name := fmt.Sprintf("l%d", i)
		antnames = append(antnames, name)
	}
	return antnames
}
