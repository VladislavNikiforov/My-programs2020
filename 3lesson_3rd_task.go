package main
import "fmt"
func main () {
	var x,y float64
	fmt.Print("Input point coordinates ->")
	fmt.Scan(&x,&y)
	if x>0 && y>0 {
		fmt.Print("	I quadrant")
	}else{
		if x<0 && y>0 {
			fmt.Print("	II quadrant")
		}else{
			if x<0 && y<0 {
				fmt.Print("	III quadrant")
			}else{ 
				if x>0 && y<0 {
					fmt.Print("	IV quadrant")
				}else{
					if x==0 && y!=0{
						fmt.Print("	the point is on Oy")
					}else{
					if x!=0 && y==0 {
						fmt.Print("	the points is on Ox")
					}else{
						fmt.Print("	The point is on centre")
					}
					}
				}
			}
		}
	}	

}
