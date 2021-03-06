//Variable

package main
import "fmt"
var i = 8
var x, y int = 6,10
var m, n, o = 1,2,3
//The := called initializer operator automatically defines and initialized variables with the given value.
/*Let's see what other common types Go supports.
float32 - a single-precision floating point value.
float64 - a double-precision floating point value.
string - a text value.
bool - Boolean true or false.*/
var q int = 19
var u string = "Vikram"
var w float32 = 1.49
var t bool = true

/*interesting feature of Go are zero values: variables that are declared without a value take the zero value of their type:
0 for numeric types,
false for the boolean type, 
"" for strings.*/
var qt int
var ut string
var wt float32
var tt bool

//In some cases your program may need values that are preserved during the program. These are called constants and they cannot be changed from their initial value.
const pi = 3.14
func main() {
    v := 20
    vat := true
    //Constants cannot be declared using the := syntax.
    fmt.Println(i)
    fmt.Println(y,x)
    fmt.Println(m, n, o)
    fmt.Println(v)
    fmt.Println(q)
    fmt.Println(u)
    fmt.Println(w)
    fmt.Println(t)
    fmt.Println(vat)
    fmt.Println(qt, ut, wt, tt)
    fmt.Println(pi)

}
