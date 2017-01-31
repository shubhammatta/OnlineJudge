package main

import (
      "fmt"
      "os"
      "time"
)

type file struct{
    file_name            string
    file_path           string
    number              int
    result              int8
    username            string
    timestamp           time.Time
    currTestcase        int
    language            string
}





func main(){
    arguments := os.Args
    fmt.Println(arguments)
        fmt.Println("Main commit");
}
