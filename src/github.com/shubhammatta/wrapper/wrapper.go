package main

import (
      "fmt"
      "os"
      "time"
      "strings"
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

//0 - C ,  1 - C++ , 2 -  Java  , 3 - Python
func detect_lang(s string) int{
    arr := strings.Split(s , ".")
    n := len(arr)
    if arr[n-1] == "cpp" || arr[n-1] == "CPP" {
        return 1
    }
    if arr[n-1]=="c" || arr[n-1] == "C"{
        return 0
    }
    if arr[n-1] == "java"{
        return 2
    }
    if arr[n-1] == "py" {
        return 3
    }
    return -1
}



func main(){
    arguments := os.Args
    fmt.Println(arguments)
    f := file{}
    f.file_name = os.Args[1]
    f.file_path = os.Args[2]
    fmt.Println(detect_lang(f.file_name))

        fmt.Println("Main commit");
}
