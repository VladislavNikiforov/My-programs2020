//6) Найти первые (и вторые) N подряд идущих составных (не простых) чисел.
package main
import "fmt"
func CompositeNumber(i int)bool{
	isComp:=false
	for j:=2;j<i;j++{
		if i%j==0 {
			isComp=true
		}
	}
	return isComp	
}
func main(){
	var n int
	fmt.Println("The programm outputs all the composite numbers in a row to n")
	fmt.Print("Input n ->")
	fmt.Scanln(&n)
	k:=0
	p:=0
		for i:=4;i<5000;i++{ 
			if !CompositeNumber(i) {
				k=0
			}
		k++
				if k==n+1 {
					m:=i
					for l:=n;l>0;l--{
						fmt.Println(i)
						i--
					}
					fmt.Println()
					for ;m<5000;m++{ 
						if !CompositeNumber(m) {
							p=0
						}
						p++
						if p==n+1 {
							for o:=n;o>0;o--{
								fmt.Println(m)
								m--
							}
						break
						}
					}
				break
				}
		}
}

		
