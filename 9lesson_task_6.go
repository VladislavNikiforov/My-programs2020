//6. Вывести в файл все простые числа вплоть до N.
// Если хочется сделать хорошо(но это сложнее), то через решето Эратосфена

package main
import "fmt"
import "math"
import "os"
import "bufio"

func main(){
	var N int
	fmt.Print("Input value N>0, program outputs prime numbers before value ->")
	fmt.Scan(&N)
	
	sl:=make([]bool,0)
	PrimeNumbers:=make([]int,0)
	for i:=0;i<=N;i++{
		f:=true
		sl=append(sl,f)
	}
	
	for i:=2;i<=int(math.Sqrt(float64(N)))+1;i++{
			for j:=i*i;j<=N;{
				if sl[i]{
					sl[j]=false
				}
					j+=i
			}	
	}
	
	for i:=1;i<=N;i++{
		if sl[i]{
			PrimeNumbers=append(PrimeNumbers,i)
		}
	}
	fout,err:=os.Create("output96.txt")
	if err!=nil{
		fmt.Print("File write error")
		return
	}
	defer fout.Close()
	out:=bufio.NewWriter(fout)
	defer out.Flush()
	fmt.Fprint(out,PrimeNumbers)
}
