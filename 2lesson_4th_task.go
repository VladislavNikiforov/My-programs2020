package main
import "fmt"

func main() {
	var a,b,c,d,e,max,min float64
fmt.Println("Input 5 numbers and the programm will calculate the max and min")
fmt.Scan(&a,&b,&c,&d,&e)
if a>b {
	max=a
}else{
	max=b
}
if c>max  {
	max=c
}	
if d>max {
	max=d
}
if e>max {
	max=e
}

if a>b {
	min=b
}else{
	min=a
}
if c<min  {
	min=c
}	
if d<min {
	min=d
}
if e<min {
	min=e
}
fmt.Printf("the max is %f and the min is %f",max,min)

}

