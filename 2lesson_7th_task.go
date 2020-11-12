package main 
import "fmt"
func main() {
	var a,b,c,max,min,mid,k float64
	fmt.Println("The program sort the numbers from min to max")
	fmt.Scan(&a,&b,&c)
	if a>b {
		max,min=a,b
	}else {
		max,min=b,a
	}
	if c>max {
		k=max
		max=c
		mid=k
	}else{
		if c>min {
			mid=c
		}else{
			k=min
			min=c
			mid=k
		}
	}		
	fmt.Printf("From the min to the max the numbers are %f,%f,%f",min,mid,max)

}
