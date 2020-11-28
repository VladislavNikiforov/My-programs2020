//Есть два возрастающих слайса, склеить их в один, чтобы он тоже возрастал.

//ЕСЛИ ПРИ ТЕСТРИРОВАНИИ ВЫ ПОЛУЧАЕТЕ НЕЖЕЛАННЫЙ ОТВЕТ, ПОПРОБУЙТЕ ВМЕСТО 1 ВВЕСТИ 2 И НАОБОРОТ 
package main
import "fmt"
func main() {
	sl:=make([]int,5)
	sl2:=make([]int,5)
	sl3:=make([]int,0)
	var n int
	
	fmt.Println("Input 5 growing slice elements ->")
	for i:=range sl{
		fmt.Scan(&sl[i])
	}
	fmt.Println("Input another 5 growing slice elements ->")
	for i:=range sl2{
		fmt.Scan(&sl2[i])
	}
	fmt.Println("Which slice starts with smaller digit? Input 1 or 2, it will help programm to work ->")
	fmt.Scan(&n)
	fmt.Println("The programm appends two growing slices in one that is also growing")
	
	if n==1{
		for i:=range sl{
			sl3=append(sl3,sl[i])
			for j:=range sl2{
				if sl[i]==sl2[j]{
					sl3=append(sl3,sl2[j])
				}
			}
		}
		for k:=range sl2{
			if sl2[k]>sl[len(sl)-1]{
				sl3=append(sl3,sl2[k])
			}
		}
		fmt.Print(sl3)
	}
	if n==2{
		for i:=range sl2{
			sl3=append(sl3,sl2[i])
			for j:=range sl{
				if sl2[i]==sl[j]{
					sl3=append(sl3,sl[j])
				}
			}
		}
		for k:=range sl{
			if sl[k]>sl2[len(sl2)-1]{
				sl3=append(sl3,sl[k])
			}
		}
		fmt.Print(sl3)
	}
	if n!=1 && n!=2 {
		fmt.Printf("\n\n\n")
		fmt.Println("WRONG NUMBER")
		fmt.Print("ERROR")
	}
}
