package main
import "fmt"
func drawFigure(n float64){
	for i:=1.0;i<=n;i++{
		for k:=1.0; k<=n-i; k++{
			fmt.Print(" ")
		}
		for j:=1.0;j<=2.0;j++{
			if i==1.0{
				j=2.0
			}
			fmt.Print("x")
			for l:=1.0;l<=i-1.0&&l<=3;l+=0.5{
				fmt.Print(" ")
			}
		}
		
		
		fmt.Println()
	}
	for i:=2.0;i<=n;i++{
		for k:=1.0;k>=n-i;k--{
			fmt.Print(" ")
		}
		for j:=2.0;j>0;j--{
			if i==n{
				j=1.0
			}
			fmt.Print("x")
			for l:=1.0;l<=i-1.0&&l>0.5;l-=0.5{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func main(){
drawFigure(3)
}
