package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func makeNamesMap() map[string]int {
    var fp *os.File
    var err error

    fp, err = os.Open("nicknames.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)

    // name : number
    nameToNum := map[string]int{}

    for scanner.Scan() {
        // Read a line
        line := scanner.Text()

        // Split A\tB to [A, B]
        array := strings.Split(line, "\t")

        num, _ := strconv.Atoi(array[0])
        name := array[1]

        nameToNum[name] = num
    }

    return nameToNum
}


func makeLinksArray() [][]int {
    var fp *os.File
    var err error

    fp, err = os.Open("links.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)
    links := [][]int{}
    cnt := 0

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

        cnt++
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return links
}

func bfs(matrix [49][49]bool, start int, goal int) {

    queue := make([]int, 0)

    isFound := false

    // Dummy node to count step
    cntPoint := -1

    // Step counter
    cnt := 1

    now := start
    target := goal

    for !isFound {
        fmt.Println(now)
        fmt.Println(queue)

        if now == target {
            fmt.Println("Found! step = " + strconv.Itoa(cnt))
            isFound = true
            break
        } else {
            nextNum := 0
            // Search root from 'now'
            for i := 0; i < 49; i++ {
                if matrix[now][i] {
                    queue = append(queue, i)
                    nextNum++
                }
            }

            if nextNum != 0 {
                // To record the point of step count
                queue = append(queue, cntPoint)
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
            queue = queue[1:]
        }
    }
}

func test(links [][]int, start int, goal int) {
    var matrix [49][49] bool

    // Put link datas into adjacency matrix
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        matrix[from][to] = true
    }

    bfs(matrix, start, goal)
    fmt.Println("--------")
}

func runTest() {
    fmt.Println("testCase 1:")
    link1 := [][]int{{0, 1}}
    test(link1, 0, 1)

    fmt.Println("testCase 2:")
    link2 := [][]int{{0, 1}, {0, 2}, {2, 1}}
    test(link2, 0, 1)

    fmt.Println("testCase 3:")
    link3 := [][]int{{0, 1}, {1, 2}}
    test(link3, 0, 2)

    fmt.Println("testCase 4:")
    link4 := [][]int{{0, 1}}
    test(link4, 0, 2)
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

    nameToNum := makeNamesMap()

    start := "jacob"
    goal := "alex"

    fmt.Println(start + " to " + goal)

    bfs(matrix, nameToNum[start], nameToNum[goal])
}

func main() {
    runTest()
    run()
}
