//4. Код Цезаря - каждую букву заменить на букву, стоящую на К позиций вправо в алфавите.
//  Сдвиг на К позиций - циклический. Разумеется, кодируются только английские| только русские | только латышские буквы.
// Надо написать и программу кодировщик, и декодеровщик, хоят в данном случае они практически совпадают.
package main
import "fmt"
import "os"
import "bufio"
import "unicode"
func main(){
	var s string
	Eng:="abcdefghijklmnopqrstuvwxyz"
	EngC:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lav:="aābcčdeēfgģhiījkķlļmnņoprsštuūvzž"
	LavC:="AĀBCČDEĒFGĢHIĪJKĶLĻMNŅOPRSŠTUŪVZŽ"
	Rus:="абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	RusC:="АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	var lang int
	fmt.Print("Program can work with files on different languages.\nChoose the language on which file is written (Input number)\n1) Eng\n2) Lav\n3) Rus\n->")
	fmt.Scan(&lang)
	if lang < 1 || lang > 3{
		fmt.Print("ERROR. INVALID NUMBER.")
	}
	var sl[]rune 
	var slC[]rune
	if lang == 1{
		sl=[]rune(Eng)
		slC=[]rune(EngC)
	}
	if lang == 2{
		sl=[]rune(Lav)
		slC=[]rune(LavC)
	}
	if lang == 3{
		sl=[]rune(Rus)
		slC=[]rune(RusC)
	}
	
	fin,err:=os.Open("input124.txt")
	if err!=nil{
		fmt.Print("file read error")
		return
	}
	in:=bufio.NewScanner(fin)
	for in.Scan(){
		str:=in.Text()
		s+=str
	}
	slS:=[]rune(s)
	var k int
	var n int
	fmt.Printf("Program can code and decode the text, input operation number:\n1) Code\n2) Decode\n->")
	fmt.Scan(&n)
	if n<1 || n>2{
		fmt.Print("ERROR. INVALID OPERATION.")
		return
	}
	fmt.Print("Input number to code/decode the text\n(by which amount each letter moves in alphabet (code - right, decode - left)) ->")
	fmt.Scan(&k)
	for ;k>=len(sl);{
		k=k%len(sl)+k/len(sl)
	}
	/*//кусок кода, который не понадобился, но мне нужен
	 var kInd int //запоминает k
	for ;k>len(sl);{
		kInd=k
		k/=len(sl)
		k+=kInd%len(sl)
		if k<=len(sl){
			k=kInd%len(sl)
			break
		}
		kInd=k
	}*/
	if n==1{
		for i:=range slS{
			if unicode.IsUpper(slS[i]){
					for j:=range slC{
						if slS[i]==slC[j]{
							if j+k>=len(slC){
								j-=len(slC)
							}
							slS[i]=slC[j+k]
							break
						}
					}
			}else{
				for j:=range sl{
					if slS[i]==sl[j]{
						if j+k>=len(sl){
							j-=len(sl)
						}
						slS[i]=sl[j+k]
						break
					}
				}
			}
		}
	}
	if n==2{
		for i:=range slS{
			if unicode.IsUpper(slS[i]){
				for j:=range slC{
					if slS[i]==slC[j]{
						if j-k<0{
							j+=len(slC)
						}
						slS[i]=slC[j-k]
						break
					}
				}
			}else{
				for j:=range sl{
					if slS[i]==sl[j]{
						if j-k<0{
							j+=len(sl)
						}
						slS[i]=sl[j-k]
						break
					}
				}
			}
		}
	}

	

	fout,err:=os.Create("input124.bak")
	if err!=nil{
		fmt.Print("file write error")
		return
	}
	out:=bufio.NewWriter(fout)
	formed:=string(slS)
	fmt.Fprint(out,formed)
	out.Flush()
	fout.Close()
	fin.Close()
	os.Remove("input124.txt")
	os.Rename("input124.bak","input124.txt")
}
