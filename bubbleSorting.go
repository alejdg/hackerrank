package main
import "fmt"

func main() {
    var n int
    fmt.Scanf("%v", &n)
    var a []int
    
    for i := 0; i < n; i++ {
        var b int
        fmt.Scanf("%v", &b)
        a = append(a, b)
    }
    
    s := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n - 1; j++ {
            if a[j] > a[ j + 1 ] {
                a[j] = a[j] + a[j+1]
                a[j+1] = a[j] - a[j+1]
                a[j] = a[j] - a[j+1]
                s++
            }
        }
        if s == 0 { 
            break 
        }
    }
    fmt.Printf("Array is sorted in %v swaps.\n", s)
    fmt.Printf("First Element: %v\n", a[0])
    fmt.Printf("Last Element: %v\n", a[len(a)-1])
}
