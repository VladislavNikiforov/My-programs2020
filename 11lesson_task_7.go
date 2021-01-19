//7. Числа храним в строках. Написать функцию, которая:
//a. сравнивает два числа
//b. увеличивает число на 1
//c. умножает число на 2
package main
import "fmt"
func Compare(s1,s2 string)(bool,bool){
	//input
	var f bool
	var m bool
	sl:=[]rune(s1) 
	sl2:=[]rune(s2)
	equa:=false
	//minus check
	if sl[0]=='-' && sl2[0]!='-'{
		return f,equa
	}
	if sl2[0]=='-' && sl[0]!='-'{
		f=true
		return f,equa
	}
	if sl[0]=='-' && sl2[0]=='-'{
		m=true
	}
	//len check
	if !m{
		if len(sl)>len(sl2){
			f=true
			return f,equa
		}
		if len(sl)<len(sl2){
			return f,equa
		}
	}else{
		if len(sl)>len(sl2){
			return f,equa
		}
		if len(sl)<len(sl2){
			f=true
			return f,equa
		}
	}
	//main
	if len(sl)==len(sl2){
		for i:=0;i<len(sl);i++{
			if sl[i]>sl2[i]{
				f=true
				break
			}else{
				if sl[i]<sl2[i]{
					break
				}
			}
			if i==len(sl)-1 && sl[len(sl)-1]==sl[len(sl2)-1] {
				equa=true
				return f,equa
			}
		}
	}
	//both minus check 
	if m{
		if f{
			f=false
		}else{
			f=true
		}
	}
	fmt.Print(m)
	return f,equa
}
func Inc(s1 string)string{
	sl:=[]rune(s1)
	n:=len(sl)-1
	var k int
	var s string
	
	if sl[0]!='-'{//алгоритм для неотрицательных чисел
		for i:=n;i>=0;i--{
			if sl[i]=='9'{
				sl[i]='0'
				k++
			}else{
				sl[i]++
				break
			}
		}
		if k==len(sl){
			s="1"
			s+=string(sl)
		}else{
			s=string(sl)
		}
	}else{ //алгоритм отрицательных чисел
		first:=sl[1]-1
		for i:=n;i>0;i--{
			if sl[i]=='0'{
				sl[i]='9'
				k++
			}else{
				sl[i]--
				break
			}
		}
		if k==len(sl)-2 && k>=1{
			s="-"
			if string(first)!="0"{
				s+=string(first)
			}
			for i:=0;i<len(sl)-2;i++{
				s+="9"
			}
		}else{
			if string(sl)!="-0"{
				s=string(sl)
			}else{
				s="0"
			}
		}
	}
	return s
}
func Mult(s1 string)string{ //не работает, не закончено
	sl:=[]rune(s1)
	var mark bool
	for i:=len(sl)-1;i>=0;i--{
		n:=sl[i]
		
		if sl[i]+sl[i]-'0'>'9'{ //алгоритм умножения
			
			for k:=1;k<int(n-'0');k++{
				if sl[i]=='9'{
					sl[i]='0'
					mark=true
				}
				sl[i]++
			}
				
			
		}else{
			sl[i]+=sl[i]-'0'
		}
		if mark{ //алгоритм прибавления единицы, если предыдушее произведение цифр >9
			if i==0{
				s:="1"
				s+=string(sl)
				return s
			}
			if i!=0{
				sl[i-1]++
			}
			mark=false
		}
	}
	return string(sl) 
}
func main(){
	var s1,s2 string
	var k int
	fmt.Printf("Choose operation -> 1-3\n1)Compare\n2)Increase by 1\n3)Multiply by 2\n")
	fmt.Scan(&k)
	if k==1 {
		fmt.Print("Input two values, program will work with them as with strings ->")
		fmt.Scan(&s1,&s2)
		f,equa:=Compare(s1,s2)
		if !equa{
			if f {
				fmt.Printf("%s>%s",s1,s2)
			}else{
				fmt.Printf("%s<%s",s1,s2)
			}
		}else{
			fmt.Print("Values are equal")
		}
	}
	if k==2 || k==3 {
		fmt.Print("Input value  program will work with it as with string ->")
		fmt.Scan(&s1)
	if k==2{
			fmt.Print(Inc(s1))
		}
		if k==3{
			fmt.Print(Mult(s1))
		}
	}
	if k<1 || k>3 {
		fmt.Print("ERROR.Unexpected operation.")
		return
	}
	
}
