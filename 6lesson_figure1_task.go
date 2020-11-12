package main
import "fmt"
func drawFigure(n int){
	for i:=1; i<=n;i++{
		for j:=0;j<i;j++{
			fmt.Print(i)
		}
		fmt.Println()
	}
}
func main() {
	drawFigure(9)

}
