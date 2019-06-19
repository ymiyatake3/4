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

func bfs(matrix [49][49]bool, start int, goal int) {

    queue := make([]int, 0)
    var isConnected[49] bool
    isFound := false

    // Dummy node to count step
    cntPoint := -1

    // Step counter
    cnt := 1

    now := start
    target := goal

    for !isFound {

        // For debugging
        //fmt.Println(now)
        //fmt.Println(queue)

        if !isConnected[now] {

            isConnected[now] = true

            if now == target {
                fmt.Println("Found! step = " + strconv.Itoa(cnt))
                isFound = true
                break
            } else {

                // If there is next node
                existNext := false

                // Search root from 'now'
                for i := 0; i < 49; i++ {
                    if matrix[now][i] {
                        queue = append(queue, i)
                        existNext = true
                    }
                }

                if existNext {
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

func searchAllConnected(isTest bool, matrix [49][49]bool, start int) {

    queue := make([]int, 0)
    var isConnected[49] bool

    // Steps to reach each node
    var count[49] int

    now := start

    // Dummy node to count step
    cntPoint := -1

    // Step counter
    cnt := 1

    for {
        // For debugging
        //fmt.Println(now)
        //fmt.Println(queue)

        if !isConnected[now] {

            isConnected[now] = true
            count[now] = cnt

            // If there is next node
            existNext := false

            // Search root from 'now'
            for i := 0; i < 49; i++ {
                if matrix[now][i] {
                    queue = append(queue, i)
                    existNext = true
                }
            }

            if existNext {
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
            } else {
                fmt.Println(numToName[i] + " : Not Connected")
            }
        }
    }
}



func test(mode string, links [][]int, start int, goal int) {
    var matrix [49][49] bool

    // Put link datas into adjacency matrix
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        matrix[from][to] = true
    }


    if mode == "bfs" {
        bfs(matrix, start, goal)
    } else if mode == "connected" {
        searchAllConnected(true, matrix, start)
    }
    fmt.Println("--------")
}


func runTest() {
    fmt.Println("testCase 1:")
    link1 := [][]int{{0, 1}}
    test("bfs", link1, 0, 1)

    fmt.Println("testCase 2:")
    link2 := [][]int{{0, 1}, {0, 2}, {2, 1}}
    test("bfs", link2, 0, 1)

    fmt.Println("testCase 3:")
    link3 := [][]int{{0, 1}, {1, 2}}
    test("bfs", link3, 0, 2)

    fmt.Println("testCase 4:")
    link4 := [][]int{{0, 1}}
    test("bfs", link4, 0, 2)

    fmt.Println("testCase 5:")
    link5 := [][]int{{0, 1}}
    test("connected", link5, 0, 0)

    fmt.Println("testCase 6:")
    link6 := [][]int{{0, 1}, {0, 2}}
    test("bfs", link6, 1, 2)

    fmt.Println("testCase 6.5:")
    test("connected", link6, 0, 0)

    fmt.Println("testCase 7:")
    link7 := [][]int{{0, 1}, {1, 2}}
    test("connected", link7, 0, 0)

    fmt.Println("testCase 8:")
    link8 := [][]int{{0, 1}, {1, 0}}
    test("connected", link8, 0, 0)

}

func run() {
    links := makeLinksArray()
    var matrix [49][49] bool

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
    searchAllConnected(false, matrix, nameToNum[start])

}

func main() {
    setNamesMap()
    runTest()
    run()
}
