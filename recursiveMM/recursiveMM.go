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

func add(a, b [][]int) (res[][]int){
    l := len(a)
    res = make([][]int, l)
    for i := range res{
        res[i] = make([]int,l)
    }
    for i := 0; i < l; i++{
        for j := 0; j < l; j++{
            res[i][j] = a[i][j] + b[i][j]
        }
    }
    return
}
func sub(a, b [][]int) (res[][]int){
    l := len(a)
    res = make([][]int, l)
    for i := range res{
        res[i] = make([]int,l)
    }
    for i := 0; i < l; i++{
        for j := 0; j < l; j++{
            res[i][j] = a[i][j] - b[i][j]
        }
    }
    return
}

func stitch4(a, b, c ,d [][]int)([][]int){
    l := len(a)
    res := a
    for i := 0; i < l; i++{
	res = append(a, b[i])
        res[i] = append(res[i], c[i]...)
        res[i+l] = append(res[i+l], d[i]...)
    }
    return res
}

func strassen(a, b[][]int)([][]int){
    l := len(a)
    if l <= 2 {
        return multiply(l, a, b)
    } else {
        a11 := a[0:l/2][0:l/2]
        a12 := a[0:l/2][l/2:]
        a21 := a[l/2:][0:l/2]
        a22 := a[l/2:][l/2:]
        b11 := b[0:l/2][0:l/2]
        b12 := b[0:l/2][l/2:]
        b21 := b[l/2:][0:l/2]
        b22 := b[l/2:][l/2:]
        m1 := strassen(add(a11, a22),add(b11,b22))
        m2 := strassen(add(a21,a22),b11)
        m3 := strassen(a11,sub(b12, b22))
        m4 := strassen(a22,sub(b21, b11))
        m5 := strassen(add(a11, a12),b22)
        m6 := strassen(sub(a21, a11), add(b11, b12))
        m7 := strassen(sub(a12, a22), add(b21, b22))
        c11 := add(m1, add(sub(m4, m5), m7))
        c12 := add(m3, m5)
        c21 := add(m2, m4)
        c22 := add(add(sub(m1, m2), m3), m6)
        return stitch4(c11,c12,c21,c22)
    }
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

    print_correct(strassen(mtx_a, mtx_b))
}
