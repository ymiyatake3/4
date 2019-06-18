package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func makeNamesMap() (map[string]int, map[int]string) {
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

    // number : name
    numToName := map[int]string{}

    for scanner.Scan() {
        // Read a line
        line := scanner.Text()

        // Split A\tB to [A, B]
        array := strings.Split(line, "\t")

        num, _ := strconv.Atoi(array[0])
        name := array[1]

        nameToNum[name] = num
        numToName[num] = name
    }

    return nameToNum, numToName
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

func bfs(matrix [49][49]bool, nameToNum map[string]int, numToName map[int]string) {
    start := "jacob"
    goal := "alex"

    queue := make([]int, 0)

    isFound := false

    // Step counter
    cnt := 1

    now := nameToNum[start]
    target := nameToNum[goal]

    for !isFound {
        fmt.Println(now)
        fmt.Println(queue)
        if now == target {
            fmt.Println("Found " + goal + "! step = " + strconv.Itoa(cnt))
            isFound = true
        } else {
            // Search root from 'now'
            for i := 0; i < 49; i++ {
                if matrix[now][i] {
                    queue = append(queue, i)
                }
            }

            // To record the point of step count
            queue = append(queue, 100)
        }

        if len(queue) == 0 {
            fmt.Println("Not found")
            break
        }

        // Move to top of the queue
        now = queue[0]
        queue = queue[1:]

        if now == 100 {
            cnt++
            now = queue[0]
            queue = queue[1:]
        }
    }
}


func test() {
    var matrix [49][49] bool
    links := makeLinksArray()

    // Put link datas into adjacency matrix
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        matrix[from][to] = true
    }

    nameToNum, numToName := makeNamesMap()

    bfs(matrix, nameToNum, numToName)
}

func run() {
    var matrix [49][49] bool
    links := makeLinksArray()

    // Put link datas into adjacency matrix
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        matrix[from][to] = true
    }

    nameToNum, numToName := makeNamesMap()

    bfs(matrix, nameToNum, numToName)
}


func main() {
    run()
}
