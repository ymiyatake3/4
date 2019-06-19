package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

var nameToNum map[string]int = map[string]int{}
var numToName map[int]string = map[int]string{}

const pageNum int = 1483276

func setNamesMap() {

    var fp *os.File
    var err error

    // Open file
    fp, err = os.Open("wiki_pages.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)

    for scanner.Scan() {
        // Read a line
        line := scanner.Text()

        // Split A\tB to [A, B]
        array := strings.Split(line, "\t")

        // Get number and name data
        num, _ := strconv.Atoi(array[0])
        name := array[1]

        nameToNum[name] = num
        numToName[num] = name
    }
}


func makeAdjacencyList() map[int][]int {

    adjList := make(map[int][]int)

    var fp *os.File
    var err error

    // Open file
    fp, err = os.Open("wiki_links.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)

    for scanner.Scan() {
        // Read a line
        line := scanner.Text()

        // Split A\tB to [A, B]
        array := strings.Split(line, "\t")

        // Convert each number to int
        from, _ := strconv.Atoi(array[0])
        to, _ := strconv.Atoi(array[1])

        adjList[from] = append(adjList[from], to)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return adjList
}

func bfs(adjList map[int][]int, start int, goal int) {

    queue := make([]int, 0)
    var isConnected[pageNum] bool

    // Dummy node to count step
    cntPoint := -1

    // Step counter
    cnt := 1

    now := start
    target := goal

    for {
        // For debugging
        //fmt.Println(now)
        //fmt.Println(queue)

        if !isConnected[now] {

            isConnected[now] = true

            if now == target {
                fmt.Println("Found! step = " + strconv.Itoa(cnt))
                break
            } else {
                // If next root from 'now' exists
                _, exist := adjList[now]
                if exist {
                    // Add next nodes to queue
                    for i := 0; i < len(adjList[now]); i++ {
                        queue = append(queue, adjList[now][i])
                    }
                    // Add counting point
                    queue = append(queue, cntPoint)
                }
            }
        }

        if len(queue) == 0 {
            fmt.Println("Not found")
            break
        }

        // Move to top of the queue
        now = queue[0]
        queue = queue[1:]

        if now == cntPoint {
            if len(queue) == 0 {
                fmt.Println("Not found")
                break
            }
            cnt++
            now = queue[0]
            queue = queue[1:]
        }
    }
}



func test(mode string, links [][]int, start int, goal int) {
    adjList := make(map[int][]int)

    // Put link datas into adjacency list
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        adjList[from] = append(adjList[from], to)
    }

    fmt.Println(links)

    if mode == "bfs" {
        fmt.Println(strconv.Itoa(start) + " to " + strconv.Itoa(goal))
        bfs(adjList, start, goal)
    }
    fmt.Println("--------")
}


func runTest() {

    link1 := [][]int{{0, 1}}
    link2 := [][]int{{0, 1}, {1, 2}}
    link3 := [][]int{{0, 1}, {1, 2}, {0, 2}}
    link4 := [][]int{{0, 1}, {0, 2}}
    //link5 := [][]int{{0, 1}, {1, 0}}

    // bfs
    fmt.Println("testCase 1-1:")
    test("bfs", link1, 0, 1)

    fmt.Println("testCase 1-2:")
    test("bfs", link1, 0, 2)

    fmt.Println("testCase 1-3:")
    test("bfs", link2, 0, 2)

    fmt.Println("testCase 1-4:")
    test("bfs", link3, 0, 1)

    fmt.Println("testCase 1-5:")
    test("bfs", link4, 1, 2)

}

func run() {
    adjList := makeAdjacencyList()

    // Search and count step
    start := "Google"
    goal := "渋谷"

    snum, sexist := nameToNum[start]
    gnum, gexist := nameToNum[goal]

    // Check if the words exist in pages list
    if !sexist {
        fmt.Println("The word " + start + "doesn't extsts.")
        return
    }
    if !gexist {
        fmt.Println("The word " + goal + "doesn't extsts.")
        return
    }

    fmt.Println(start + " to " + goal)
    bfs(adjList, snum, gnum)

}

func main() {
    setNamesMap()
    runTest()
    run()
}
