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
        fmt.Println(now)
        fmt.Println(queue)

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


func searchAllConnected(isTest bool, adjList map[int][]int, start int) {

    queue := make([]int, 0)
    var isConnected[pageNum] bool

    // Steps to reach each node
    var count[pageNum] int

    now := start

    // Dummy node to count step
    cntPoint := -1

    // Step counter
    cnt := 1

    for {
        // For debugging
        fmt.Println(now)
        fmt.Println(queue)

        if !isConnected[now] {

            isConnected[now] = true
            count[now] = cnt

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

        // Move to top of the queue
        now = queue[0]
        queue = queue[1:]

        if now == cntPoint {
            if len(queue) == 0 {
                break
            }
            cnt++
            now = queue[0]
            queue = queue[1:]
        }
    }

    // Print result
    if isTest {
        fmt.Println("Connected nodes: ")
        for i := 0; i < len(isConnected); i++ {
            if i == start {
                fmt.Println(strconv.Itoa(i) + " : me")
            } else if isConnected[i] && i != start {
                fmt.Println(strconv.Itoa(i) + " : step = " + strconv.Itoa(count[i]))
            }
        }
    } else {
        for i := 0; i < len(isConnected); i++ {
            if i == start {
                fmt.Println(strconv.Itoa(i) + " : me")
            } else if isConnected[i] && i != start {
                fmt.Println(numToName[i] + " : step = " + strconv.Itoa(count[i]))
            }
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

    if mode == "bfs" {
        bfs(adjList, start, goal)
    } else if mode == "connected" {
        searchAllConnected(true, adjList, start)
    }
    fmt.Println("--------")
}


func runTest() {

    link1 := [][]int{{0, 1}}
    link2 := [][]int{{0, 1}, {1, 2}}
    link3 := [][]int{{0, 1}, {1, 2}, {0, 2}}
    link4 := [][]int{{0, 1}, {0, 2}}
    link5 := [][]int{{0, 1}, {1, 0}}

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

    // searchAllConnected
    fmt.Println("testCase 2-1:")
    test("connected", link1, 0, 0)

    fmt.Println("testCase 2-2:")
    test("connected", link2, 0, 0)

    fmt.Println("testCase 2-3:")
    test("connected", link3, 0, 0)

    fmt.Println("testCase 2-4:")
    test("connected", link4, 0, 0)

    fmt.Println("testCase 2-5:")
    test("connected", link5, 0, 0)

}

func run() {
    adjList := makeAdjacencyList()

    // Search and count step
    start := numToName[210038]
    goal := numToName[37428]
    fmt.Println(start + " to " + goal)
    //bfs(adjList, nameToNum[start], nameToNum[goal])
    bfs(adjList, 210038, 37428)

    /*
    //fmt.Println("--------")

    // Search all steps to the other nodes
    start := "アカマダラハナムグリ"
    searchAllConnected(false, adjList, nameToNum[start])
    */
}

func main() {
    setNamesMap()
    runTest()
    run()
}
