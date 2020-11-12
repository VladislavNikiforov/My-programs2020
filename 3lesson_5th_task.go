package main
import "fmt"
//Шахматы на полставки
func main () {
	var x1,x2,y1,y2 int
	fmt.Print("Input the coordinates of the cell at the chess desk -> ")
	fmt.Scan(&x1,&y1,&x2,&y2)
	if (x1>0) && (y1>0) && (x1<=8) && (y1<=8) {
		if (x1+y1) % 2==0 {
			fmt.Println("First figure's cell is black")
		}else{
			fmt.Println("First figure's cell is white")
		}
	}else{
		fmt.Print("WRONG COORDINATES")
	}
	
	if (x2>0) && (y2>0) && (x2<=8) && (y2<=8) {
		if (x2+y2) % 2==0 {
			fmt.Println("Second figure's cell is black")
		}else{
			fmt.Println("Second figure's cell is white")
		}
	}	
	fmt.Println()
	fmt.Println("Possible figure moves:")	
	if x1+1==x2 || y1+1==y2 || x1-1==x2|| y1-1==y2 ||x1+1==x2 && y1+1==y2 ||x1+1==x2 && y1-1==y2||x1-1==x2 && y1+1==y2||x1-1==x2 && y1-1==y2{
		fmt.Println("King can eat the figure") 
	}
	if x1==x2||y1==y2 {
		fmt.Println("Rook can eat the figure")
	}
	if x1+1==x2 && y1+1==y2 ||x1+1==x2 && y1-1==y2||x1-1==x2 && y1+1==y2||x1-1==x2 && y1-1==y2 || x1+2==x2 && y1+2==y2 ||x1+2==x2 && y1-2==y2||x1-2==x2 && y1+2==y2||x1-2==x2 && y1-2==y2 || x1+3==x2 && y1+3==y2 ||x1+3==x2 && y1-3==y2||x1-3==x2 && y1+3==y2||x1-3==x2 && y1-3==y2 || x1+4==x2 && y1+4==y2 ||x1+4==x2 && y1-4==y2||x1-4==x2 && y1+4==y2||x1-4==x2 && y1-4==y2 || x1+5==x2 && y1+5==y2 ||x1+5==x2 && y1-5==y2||x1-5==x2 && y1+5==y2||x1-5==x2 && y1-5==y2 || x1+6==x2 && y1+6==y2 ||x1+6==x2 && y1-6==y2||x1-6==x2 && y1+6==y2||x1-6==x2 && y1-6==y2 || x1+7==x2 && y1+7==y2 ||x1+7==x2 && y1-7==y2||x1-7==x2 && y1+7==y2||x1-7==x2 && y1-7==y2 || x1+8==x2 && y1+8==y2 ||x1+8==x2 && y1-8==y2||x1-8==x2 && y1+8==y2||x1-8==x2 && y1-8==y2 {
		fmt.Println("Bishop can eat the figure")
	}
	if x1==x2||y1==y2 || x1+1==x2 && y1+1==y2 ||x1+1==x2 && y1-1==y2||x1-1==x2 && y1+1==y2||x1-1==x2 && y1-1==y2 || x1+2==x2 && y1+2==y2 ||x1+2==x2 && y1-2==y2||x1-2==x2 && y1+2==y2||x1-2==x2 && y1-2==y2 || x1+3==x2 && y1+3==y2 ||x1+3==x2 && y1-3==y2||x1-3==x2 && y1+3==y2||x1-3==x2 && y1-3==y2 || x1+4==x2 && y1+4==y2 ||x1+4==x2 && y1-4==y2||x1-4==x2 && y1+4==y2||x1-4==x2 && y1-4==y2 || x1+5==x2 && y1+5==y2 ||x1+5==x2 && y1-5==y2||x1-5==x2 && y1+5==y2||x1-5==x2 && y1-5==y2 || x1+6==x2 && y1+6==y2 ||x1+6==x2 && y1-6==y2||x1-6==x2 && y1+6==y2||x1-6==x2 && y1-6==y2 || x1+7==x2 && y1+7==y2 ||x1+7==x2 && y1-7==y2||x1-7==x2 && y1+7==y2||x1-7==x2 && y1-7==y2 || x1+8==x2 && y1+8==y2 ||x1+8==x2 && y1-8==y2||x1-8==x2 && y1+8==y2||x1-8==x2 && y1-8==y2 {
		fmt.Println("Queen can eat the figure") 
	}
	if x1+3==x2 && y1+1==y2||x1+3==x2 && y1-1==y2||x1-3==x2 && y1+1==y2||x1-3==x2 && y1-1==y2|| x1-1==x2 && y1+3==y2||x1+1==x2 && y1+3==y2||x1+1==x2 && y1-3==y2||x1-1==x2 && y1-3==y2 {
		fmt.Println("Knight can eat the figure")
	}
}
