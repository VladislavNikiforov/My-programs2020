//5) Найти все пары простых чисел-близнецов (т.е. простых чисел, отличающихся на 2), не превосходящие 30000.
package main
import "fmt"
func prime(){
	k:=0
		for i:=3;i<30000;i+=2 {
			l:=true
			for j:=3;j<i;j+=2{
				if i%j==0{
					l=false
				}
				
			}
			if l{
				if i-k==2 {
					fmt.Println(i,k)
				}
				k=i
			}
		}
}
func main(){
	prime()
}
