package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func makeNamesMap() map[string]int{
    var fp *os.File
    var err error

    fp, err = os.Open("nicknames.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)

    // name : number
    names := map[string]int{}

    for scanner.Scan() {
        // Read a line
        line := scanner.Text()

        // Split A\tB to [A, B]
        array := strings.Split(line, "\t")

        num, _ := strconv.Atoi(array[0])
        name := array[1]

        names[name] = num
    }

    return names
}

func makeLinksArray() [][]int{
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


func main() {
    var matrix [49][49] bool
    links := makeLinksArray()

    // Put link datas into adjacency matrix
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        matrix[from][to] = true
    }

    nameToNum := makeNamesMap()
    myName := "alex"

    queue := make([]int, 0)

    now := nameToNum[myName]
    queue = append(queue, now)
    fmt.Println(queue)
}
