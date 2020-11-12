/*
 * 1) Найти все решения в целых числах уравнения
a.  x^2^ – xy + y^2^ = N
b.  x^3^ + y^3^ = N
c.  sqrt(x) + sqrt(y) = sqrt(N)
d.  2x^2^ – 5xy + 2y^2^ = N
e.  x(N + x) = y^2^
 * */
package main
import "fmt"
import "math"
func EquA(n int){
	for   x:=-1000; x<= 1000;x++{
		for y:=-1000;y<= 1000;y++{
			if n==((x*x)-(x*y)+(y*y)){
				fmt.Println(x,y)
			}
		}
	}
}
func EquB(n int){
		for x:=-1000;x<=1000;x++{
			for y:=-1000;y<=1000;y++{
				if n==((x*x*x)+(y*y*y)){
					fmt.Println(x,y)
					}
				}
			}
}
func EquC(n int){
		for x:=-1000;x<=1000;x++{
			for y:=-1000;y<=1000;y++{
				if math.Sqrt(float64(n))==math.Sqrt(float64(x))+math.Sqrt(float64(y)){
					fmt.Println(x,y)
					}
				}
			}
}
func EquD(n int){
		for x:=-1000;x<=1000;x++{
			for y:=-1000;y<=1000;y++{
				if n==2*x*x-5*x*y+2*y*y{
					fmt.Println(x,y)
					}
				}
			}
}
func EquE(n int){
		for x:=-1000;x<=1000;x++{
			for y:=-1000;y<=1000;y++{
				if x==0{
					x++
				}
				if y==0{
				y++	
				}
				if float64(n)==((float64(y)*float64(y))/float64(x))-float64(x){
					fmt.Println(x,y)
				}
			}
		}
}
func main (){
	var n,k int 
	fmt.Print("Choose the equation to solve ->")
	fmt.Scan(&k)
	fmt.Print("Input the value of function, the program will all the roots ->")
	fmt.Scan(&n)
	fmt.Println()
	if k==1{
		EquA(n)
	}
	if k==2 {
		EquB(n)
		}
	if k==3 {
		EquC(n)
	}
	if k==4 {
		EquD(n)
	}
	if k==5 {
		EquE(n)
	}	
}

