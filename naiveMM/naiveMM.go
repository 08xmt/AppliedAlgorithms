package main

import (
    "fmt"
    "os"
    "bufio"
    "encoding/csv"
    "strconv"
)

func multiply(l int, a, b [][]int)(res [][]int){
    res = make([][]int, l)
    for i := range res{
        res[i] = make([]int,l)
    }
    for i := 0; i < l; i++{
        for j := 0; j < l; j++{
            for x := 0; x < l; x++{
                res[i][j] += a[i][x] * b[x][j]
            }
        }
    }
    return
}
func parse(l int, filename string)(parsed [][]int){
    file, _ := os.Open(filename)
    reader := csv.NewReader(bufio.NewReader(file))
    reader.Comment= '#'
    parsed = make([][]int, l)
    for i := range parsed{
        parsed[i] = make([]int,l)
    }
    for {
        line, err := reader.Read()
        if err != nil {
            break
        }
	i,_ := strconv.ParseInt(line[0],10,64)
	j,_ := strconv.ParseInt(line[1],10,64)
	k,_ := strconv.ParseInt(line[2],10,64)
	parsed[i][j] = int(k)
    }
    return
}
func print_correct( res [][]int){
    for i := range res{
	for j := range res[i]{
	    fmt.Print(res[i][j], " ")
	}
    }
    fmt.Println()
}

func main(){
    big, _ := strconv.ParseInt(os.Args[1],10,64)
    l := int(big)
    filename1 := os.Args[2]
    filename2 := os.Args[3]
    mtx_a := parse(l, filename1)
    mtx_b := parse(l, filename2)

    print_correct(multiply(l, mtx_a, mtx_b))
}
