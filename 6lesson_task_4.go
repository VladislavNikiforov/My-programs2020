//4) Найти все четырех-значные квадраты, запись которых состоит только из четных цифр.
package main
import "fmt"
func FSq(){
	for s:=1;s<100;s++{
		a:=1
		if (s*s)/1000>0{
			dig:=s*s
			for dig>0{
				if dig%2==0 {
					dig/=10	
				}else{
					break
				}
				a=dig
			}
			if a==0{
			fmt.Println(s*s)
			}
		}
		
	}
}
func main(){
	FSq()
}

