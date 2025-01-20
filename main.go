package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)
var debugOut = bufio.NewWriter(os.Stderr)

const INF = int(1e9 + 7)

type Point struct {
	X, Y int
}

func main() {
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)
	grid := make([][]int, N)
	dist := make([][]int, N)
	prev := make([][]Point, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, M)
		dist[i] = make([]int, M)
		prev[i] = make([]Point, M)
		for j := 0; j < M; j++ {
			fmt.Fscan(in, &grid[i][j])
			dist[i][j] = INF
			prev[i][j] = Point{-1, -1}
			if grid[i][j] < 0 || grid[i][j] > 9 {
				log.Fatalln("invalid grid: values should be in [0, 9]")
				os.Exit(1)
			}
		}
	}

	inside := func(x, y int) bool {
		return 0 <= x && x < N && 0 <= y && y < M
	}

	var xStart, yStart, xFinish, yFinish int
	fmt.Fscan(in, &xStart, &yStart, &xFinish, &yFinish)
	if !inside(xStart, yStart) || !inside(xFinish, yFinish) {
		log.Fatalln("invalid start or finish coord")
		os.Exit(1)
	}
	if grid[xStart][yStart] == 0 || grid[xFinish][yFinish] == 0 {
		log.Fatalln("impassible path")
		os.Exit(1)
	}

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	dist[xStart][yStart] = 0
	q := make([]Point, 0)
	q = append(q, Point{xStart, yStart})
	for len(q) > 0 {
		x, y := q[0].X, q[0].Y
		q = q[1:]
		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			// клетка находится вне поля или является непроходимой
			if !inside(nx, ny) || grid[nx][ny] == 0 {
				continue
			}
			if dist[nx][ny] > dist[x][y]+1 {
				dist[nx][ny] = dist[x][y] + 1
				prev[nx][ny] = Point{x, y}
				q = append(q, Point{nx, ny})
			}
		}
	}
	if dist[xFinish][yFinish] == INF {
		log.Fatalln("no path")
		os.Exit(1)
	}
	path := make([]Point, 0)
	for x, y := xFinish, yFinish; x != -1 && y != -1; x, y = prev[x][y].X, prev[x][y].Y {
		path = append(path, Point{x, y})
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Fprintln(out, path[i].X, path[i].Y)
	}
	fmt.Fprintln(out, ".")
}

//lint:ignore U1000 aboba
func debug(format string, args ...any) {
	defer debugOut.Flush()
	fmt.Fprintf(debugOut, format, args...)
}
