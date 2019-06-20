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

const nodeNum int = 1483277

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


func search() [nodeNum]bool {

    var  [nodeNum]bool

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

        // Convert number to int
        from, _ := strconv.Atoi(array[0])
        to, _ := strconv.Atoi(array[1])

        connected[from] = true
        connected[to] = true
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return connected
}



func run() {

    connected := search()
    fmt.Println("Independent pages:")

    cnt := 0
    for i := 0; i < len(connected); i++ {
        if !connected[i] {
            cnt++
            fmt.Println(strconv.Itoa(cnt) + ": " + numToName[i])
        }
    }

}

func main() {
    setNamesMap()
    run()
}
