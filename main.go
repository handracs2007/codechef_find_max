// https://www.codechef.com/problems/LOSTMAX
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    nCases, _ := strconv.Atoi(scanner.Text())

    for i := 0; i < nCases; i++ {
        scanner.Scan()
        line := scanner.Text()

        sNums := strings.Split(line, " ")

        skip := true
        max := -1
        for j := 0; j < len(sNums); j++ {
            num, _ := strconv.Atoi(sNums[j])

            if skip && num == len(sNums)-1 {
                // Skip once if the number equals to the count of numbers
                skip = false
                continue
            }

            if num > max {
                max = num
            }
        }

        fmt.Println(max)
    }
}
