package main 
import "fmt"

func main() {
	var ac,bc,ap,bp,k float64
	fmt.Println("Input sides of a paper ->")
	fmt.Scan(&ap,bp)
	fmt.Println("Input the sides of a convert")
	fmt.Scan(&ac,&bc)
	fmt.Println("Choose the mode 1 - common, 2 - fold in half")
	fmt.Scan(&k)
	if k==1 {
		if ac>=ap && bc>=bp {
			fmt.Print("Поместится")
		}
	}
if k==2 {
	if ac>=ap/2 && bc>=bp {
		fmt.Print("Поместится")
	}else{
		if bc>=bp/2 && ac>=ap {
			fmt.Print("Поместится")
		}else{
			fmt.Print("Не поместится")
		}
	}
}	
}

