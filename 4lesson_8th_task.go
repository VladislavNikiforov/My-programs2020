package main
import "fmt"
func main () {
	var a,re,k10 int
	fmt.Println("The programm outputs a digitally-reversed number")
	fmt.Scan(&a)
	for a!=0 {
		re=a%10
		a=a/10
		k10*=10
		k10+=re
	}
	fmt.Printf("Reversed number %d is %d",a,k10)
}
//ivan method 
/*
  package main 
import ("fmt"
       )

func main () {
    var i1, i2, dig int
    fmt.Println("Введите число")
    fmt.Scanln(&i1)
    for i1>0 {
        dig = i1%10
        i1 /= 10
        i2 *= 10
        i2 += dig
    }
fmt.Println(i2)
}
*/
