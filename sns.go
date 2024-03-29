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

const nodeNum int = 49

func setNamesMap() {

    var fp *os.File
    var err error

    // Open file
    fp, err = os.Open("nicknames.txt")
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


func makeLinksArray() [][]int {

    links := [][]int{}

    var fp *os.File
    var err error

    // Open file
    fp, err = os.Open("links.txt")
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

        link := []int{from, to}
        links = append(links, link)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return links
}

func bfs(matrix [nodeNum][nodeNum]bool, start int, goal int) {

    // Dummy node to count step
    cntPoint := -1

    // Step counter
    cnt := 0

    queue := make([]int, 0)
    queue = append(queue, cntPoint)

    var visited[nodeNum] bool

    now := start

    for {
        // For debugging
        //fmt.Println(now)
        //fmt.Println(queue)

        visited[now] = true

        if now == goal {
            fmt.Println("Found! step = " + strconv.Itoa(cnt))
            break
        } else {
            // Search root from 'now'
            for i := 0; i < nodeNum; i++ {
                if matrix[now][i] {
                    if !visited[i] {
                        queue = append(queue, i)
                    }
                }
            }
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

            // remove first element and add cntPoint at last
            queue = append(queue[1:], cntPoint)
        }
    }
}

func showAllSteps(isTest bool, matrix [nodeNum][nodeNum]bool, start int) {

    // Dummy node to count step
    cntPoint := -1

    // Step counter
    cnt := 0

    queue := make([]int, 0)
    queue = append(queue, cntPoint)

    var visited[nodeNum] bool
    var count[nodeNum] int

    now := start

    for {
        // For debugging
        //fmt.Println(now)
        //fmt.Println(queue)

        if !visited[now] {

            visited[now] = true
            count[now] = cnt

            // Search root from 'now'
            for i := 0; i < nodeNum; i++ {
                if matrix[now][i] {
                    queue = append(queue, i)
                }
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

            // remove first element and add cntPoint at last
            queue = append(queue[1:], cntPoint)
        }
    }

    // Print result
    if isTest {
        fmt.Println("Connected nodes: ")
        for i := 0; i < len(visited); i++ {
            if i == start {
                fmt.Println(strconv.Itoa(i) + " : me")
            } else if visited[i] && i != start {
                fmt.Println(strconv.Itoa(i) + " : step = " + strconv.Itoa(count[i]))
            }
        }
    } else {
        for i := 0; i < len(visited); i++ {
            if i == start {
                fmt.Println(strconv.Itoa(i) + " : me")
            } else if visited[i] && i != start {
                fmt.Println(strconv.Itoa(i) + " " + numToName[i] + " : step = " + strconv.Itoa(count[i]))
            } else {
                fmt.Println(strconv.Itoa(i) + " " + numToName[i] + " : Not Connected")
            }
        }
    }
}



func test(mode string, links [][]int, start int, goal int) {
    var matrix [nodeNum][nodeNum] bool

    // Put link datas into adjacency matrix
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        matrix[from][to] = true
    }

    fmt.Println(links)

    if mode == "bfs" {
        fmt.Println(strconv.Itoa(start) + " to " + strconv.Itoa(goal))
        bfs(matrix, start, goal)
    } else if mode == "all" {
        fmt.Println("from " + strconv.Itoa(start))
        showAllSteps(true, matrix, start)
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

    // showAllSteps
    fmt.Println("testCase 2-1:")
    test("all", link1, 0, 0)

    fmt.Println("testCase 2-2:")
    test("all", link2, 0, 0)

    fmt.Println("testCase 2-3:")
    test("all", link3, 0, 0)

    fmt.Println("testCase 2-4:")
    test("all", link4, 0, 0)

    fmt.Println("testCase 2-5:")
    test("all", link5, 0, 0)
}

func run() {
    links := makeLinksArray()
    var matrix [nodeNum][nodeNum] bool

    // Put link datas into adjacency matrix
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        matrix[from][to] = true
    }

    // Count step from 'jacob' to 'alex'
    start := "jacob"
    goal := "alex"
    fmt.Println(start + " to " + goal)
    bfs(matrix, nameToNum[start], nameToNum[goal])

    fmt.Println("--------")

    // Search all steps to the other nodes
    start = "alex"
    showAllSteps(false, matrix, nameToNum[start])

}

func main() {
    setNamesMap()
    runTest()
    run()
}
