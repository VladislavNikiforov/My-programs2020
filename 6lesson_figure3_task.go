package main
import "fmt"
func drawFigure(n float64){
	for i:=1.0; i<=n;i++{
		for k:=1.0; k<=n-i; k++{
			fmt.Print(" ")
		}
		for j:=1.0;j<=i;j+=0.5{
			fmt.Print(i)
		}
		fmt.Println()
	}
	for i:=n-1; i>0;i--{
		for k:=1.0; k<=n-i; k++{
			fmt.Print(" ")
		}
		for j:=i;j>0.5;j-=0.5{
			fmt.Print(i)
		}
		fmt.Println()
	}
}
func main() {
	drawFigure(9)
}



 /*  1
 *  222
 * 33333
 *  222
 *   1*/
