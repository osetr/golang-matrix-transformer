package main
 
import "fmt"
 
func MatrixToSlice(s [][]int) {
    if s == nil {
        panic("Matrix is empty")
	}
	ColAmount := len(s[0])
	RowAmount := len(s)

	if ColAmount != RowAmount {
		panic("Not NxN matrix")
	}

    x, y := 0, 0
    rows := len(s) - 1
    columns := len(s[0]) - 1
    if rows == 0 {
        for _, v := range s[0] {
            fmt.Println(v)
        }
        return
    }
    if columns == 0 {
        for _, v := range s {
            fmt.Println(v[0])
        }
        return
    }
    for x <= rows && y <= columns {
        ax, ay := x, y
        for ay < columns {
            fmt.Println(s[ax][ay])
            ay++
        }
        for ax < rows {
            fmt.Println(s[ax][ay])
            ax++
        }
        for ay > y {
            fmt.Println(s[ax][ay])
            ay--
        }
        for ax > x {
            fmt.Println(s[ax][ay])
            ax--
        }
        x++
        y++
        rows--
        columns--
    }
}
 

func main() {
    s := [][]int{
		{1, 2, 4, 5},
		{1, 2, 4, 5}, 
		{1, 2, 4, 5}, 
		{1, 2, 4, 5}, 
	}
 
    MatrixToSlice(s)
}