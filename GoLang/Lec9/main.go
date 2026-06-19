// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
// )

// func main() {
//     reader := bufio.NewReader(os.Stdin)

//     fmt.Print("Enter your full name: ")
//     name, _ := reader.ReadString('\n')

//     fmt.Println("Hello,", name)
// }

package main

import (
    "fmt"
)

func main() {
    var age int
    fmt.Print("Enter age: ")
    fmt.Scan(&age)

    fmt.Println("Your age is:", age)
}

// It reads input based on a format string.
// fmt.Print("Enter name and age: ")
// fmt.Scanf("%s %d", &name, &age)