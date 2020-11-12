package main 
import "fmt" 
func main () {
	var y int
	fmt.Print("Input year ->")
	fmt.Scan(&y)
	if y%4==0 {
			if y%100==0 && y%400!=0 {
			fmt.Print("This year is common")
			}else{
				fmt.Print("The is a leap year")
			}
		}		
}
