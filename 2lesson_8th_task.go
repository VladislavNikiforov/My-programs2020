package main 
import "fmt"
func main () {
	var a,b,c,d,e, med,m1,m2,m4,m5 float64
	fmt.Println("Input 5 numbers and the programm calculates the median")
	fmt.Scan(&a,&b,&c,&d,&e)
	if a>b {
		m5,m4=a,b
	}if c>m5 {
		m5,m4,m3=c,a,b
	}else{
		m4,m3=c,b
	}else{
		m3=c
	}
	
	
	}else{
		m5,m4=b,a
	}if c>m5{
		m5,m4,m3=c,b,a
	}if d>m5{
		m5,m4,m3,m2=d,c,b,a
	}else{
		m4,m3=c,a
	}if d>m4{
		m5,m4,m3,m2=d,c,b,a
	}else{
		m3=c
	}if d>m3 {
		m5,m4,m3,m2=d,b,a,c













}
