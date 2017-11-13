package main
import "fmt"

type hourGlass struct {
    l1 [3]int
    l2 [1]int
    l3 [3]int
}
  
func (hg *hourGlass) newHourGlass(a [3][3]int) {
    for i := 0; i < 3; i++ {
        hg.l1[i] = a[0][i]
        hg.l3[i] = a[2][i]
    }
    hg.l2[0] = a[1][1]
}

func main() {
    a := readArray()
    var hga []hourGlass
    b := 0
    hga = getHourGlasses(a)
    for i, hg := range hga {
        s := sumHourGlass(hg)
        if i == 0 {
            b = s
        }
        if b < s {
            b = s
        }
    }
    fmt.Printf("%v", b)
}

func readArray() [6][6]int {
    var a [6][6]int
    for i := 0; i < 6; i++{
        for j := 0; j < 6; j++ {
            fmt.Scanf("%v", &a[i][j])
        }
    }
    return a
}

func getHourGlasses(a [6][6]int) []hourGlass {
    var hga []hourGlass
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            var hg hourGlass
            hg.newHourGlass([3][3]int{{a[i][j], a[i][j+1], a[i][j+2]}, {a[i+1][j], a[i+1][j+1], a[i+1][j+2]}, {a[i+2][j], a[i+2][j+1], a[i+2][j+2]}})
            hga = append(hga, hg)
        }
    }
    return hga
}

func sumHourGlass(hg hourGlass) int {
    s := 0
    for _, v := range hg.l1 {
        s = s + v
    }
    for _, v := range hg.l3 {
        s = s + v
    }
    s = s + hg.l2[0]
    return s
}
