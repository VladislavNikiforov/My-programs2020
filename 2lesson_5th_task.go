package main
import "fmt"
func main() {
	var a,b,c int
fmt.Println("Input 3 sides for a triangle")
fmt.Scan(&a,&b,&c)
if a+b>c {
	if a-b<c || b-a <c {
		fmt.Print("This triangle exists")
		}
}else {
		fmt.Print("The triangle doesn't exist")
}
}
