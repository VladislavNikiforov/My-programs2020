//3. Четное ли число (возвращаем 0/1)
package main 
import "fmt"
func pair(a int) int{
	if a%2==0 {
		return 1
	}else{
		return 0
	}
}
func main () {
	var a int 
	fmt.Scanln(&a)
	fmt.Print(pair(a))
}
