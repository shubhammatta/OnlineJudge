package main

import (
      "fmt"
      "os"
      "time"
      "strings"
      //"strconv"
      "os/exec"
      "syscall"
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


func compile_code(path string, number int)  {
    //binary := strconv.Itoa(number)
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }
    args := []string{ "-a", "-l", "-h"}
    env := os.Environ()
    fmt.Println(binary)
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}

func compile_code1(path string, number int)  {
    //binary := strconv.Itoa(number)
    var (
		cmdOut []byte
		err    error
	)
	cmdName := "ls"
	cmdArgs := []string{"-l", "-a", "-h"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	//firstSix := sha[:6]
	fmt.Println(sha)
}




func main(){
    arguments := os.Args
    fmt.Println(arguments)
    f := file{}
    f.file_name = os.Args[1]
    f.file_path = os.Args[2]
    fmt.Println(detect_lang(f.file_name))
    compile_code1("sff",44)
}
