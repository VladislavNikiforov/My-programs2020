//0) вывести все простые числа от 1 до N (или от A до  B – тогда у функции два параметра)
package main
import "fmt"
func prime(a,b int){
	for x:=a; x<b;x++{
		k:=0
		for i:=2; x>i;i++{
			if x%i==0{
				k++
			}
		}
		if k==0 && x>1{
			fmt.Println(x)
		}
	
	}
}		
func main () {
	var a,b int
	fmt.Scanln(&a,&b)
	prime(a,b)
			
}
//ИВАН НОРМ КОД
/*package main

import "fmt"

func Primezero(x int) {
    fmt.Println("Prime numbers from 0 to x")
    for i := 0; i <= x; i++ {
        isPrime := true
        if i == 0  i == 1  i == 2 {
            isPrime = false
        }
        for j := 2; j < i; j++ {
            if i%j == 0 {
                isPrime = false
            }
        }
        if isPrime == true {
            fmt.Println(i)
        }
    }
    fmt.Print("\n \n \n \n \n \n \n")
}
func Primedistance(x, y int) {
    fmt.Println("Prime numbers from a to b")
    for i := x; i <= y; i++ {
        isPrime := true
        for j := 2; j < i; j++ {
            if i%j == 0 {
                isPrime = false
            }
        }
        if isPrime == true {
            if i < x {
                fmt.Print("")
            } else {
                fmt.Println(i)
            }
        }
    }
}
func main() {
    Primezero(100)
    Primedistance(23, 100)
}*/
