/* Задачи на заполнение:
a. 0 1 0 1 0 1 ...
b. 0 2 0 4 0 6 0 8 ...
c. 1 0 2 0 3 0 4 0 ...
d. 1 2 2 3 3 3 4 4 4 4 5 5 5 5 5 6 6 6 6 6 6 ...
e. 1 0 1 1 0 1 1 1 0 1 1 1 1 0 1 1 1 1 1 0 ...
f. 1 4 7 10 13 16 19 22 25 ...
g. 1 3 7 13 21 31 43 57 73 89 ...
*/
package main
import "fmt"
const MaxLength=30
type Array[MaxLength]int
func FillA(a Array){
	for i:=0;i<MaxLength;i++{
		i++
		a[i]=1
	}
	fmt.Print(a)
}
func FillB(a Array){
	for i:=0;i<MaxLength;i++{
		i++
		a[i]=i+1
	}
	fmt.Print(a)
}
func FillC(a Array){
	k:=0
	for i:=range a{
		if i%2!=0{
			k++
			a[i]=k
		}
	}
	fmt.Print(a)
}
func FillD(a Array) {
	k:=0
	for i:=range a{
		for j:=0;j<=i;j++{
			if k<MaxLength{
				a[k]=i+1
				k++
			}
		}
	}
	fmt.Print(a)
}
func FillE(a Array) {
	k:=-1
	for i:=range a{
		for j:=0;j<i;j++{
			if i+k<MaxLength{
				a[i+k]=1
				k++
				
			}
		}
	}
	fmt.Println(a)
}
func FillF(a Array) {
	for i:=range a{
		
		for j:=0;j<i;j++{
			a[i]+=3
		}
		a[i]++
		if i==0{
			a[i]=1
		}
	}
	fmt.Print(a)
}
func FillG(a Array){
	k:=0
	a[0]=1
	for i:=range a{
		for j:=0;j<i;j++{
			k++
			a[i]=2*k
			a[i]++
		}
	}
	fmt.Print(a)
}
func main () {
	var a Array
	var k string
	fmt.Print("Choose the operation, the program fills th array: a-g ->")
	fmt.Scanln(&k)
	if k=="a"{
		FillA(a)
	}
	if k=="b"{
		FillB(a)
	}
	if k=="c"{
		FillC(a)
	}
	
	if k=="d"{
		FillD(a)
	}
	if k=="e"{
		FillE(a)
	}
	
	if k=="f"{
		FillF(a)
	}
	if k=="g"{
		FillG(a)
	}
}
