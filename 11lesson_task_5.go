//5. В файле есть строка, состоящая из букв и пробелов: 
//    a. подсчитать количество слов (словом считаем произвольную последовательность букв, не содержащую пробелов, и ограниченную с обоих сторон пробелами (или находящихся на краях строки)_)
//    b. извлечь из такой строки K-ое слово
//    c. удалить из такой строки К-ое слово и записать получившуюся строку обратно в файл
package main

import "fmt"
import "os"
import "bufio"
import "strings"


func main(){
	const FileLenght=8
	f,err:=os.Open("input115.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	defer f.Close()
	scanner:=bufio.NewScanner(f)
	
	var strNumber int
	fmt.Print("Input string number (number from file) 1-8 ->")
	fmt.Scan(&strNumber)
	if strNumber<1 || strNumber>8 {
			fmt.Print("Incorrect number of string")
			return
		}
	k:=0
	
	for scanner.Scan(){
		k++
		str:=scanner.Text()
		if k==strNumber{
			wordlist:=strings.Split(str," ")
			for i:=0; i < len(wordlist); {
				if  wordlist[i]>="a" && wordlist[i]<="z" || wordlist[i]>="A" && wordlist[i]<="Z" || wordlist[i]>="а" && wordlist[i]<="я" || wordlist[i]>="А" && wordlist[i]<="Я" {
					i++
				}else{
					if wordlist[i]!=" "{
						wordlist=append(wordlist[:i],wordlist[i+1:]...)	
					}
				}
			}
			fmt.Printf("The amount of words in the string is %d\n",len(wordlist))//кол-во слов  в строке
			
			var wordN int
			var opN int
			fmt.Printf("You are able to do some operations with selected string:\n1)Extract word\n2)Delete word\nChoose operation ->")
			fmt.Scan(&opN)
			if opN==1{ //extract
				fmt.Print("Input the word number you want to extract from the string ->")
				fmt.Scan(&wordN)
				fmt.Printf("Extracted word is {%s}",wordlist[wordN-1])
				break
			}
			if opN==2{ //delete - не работает
				fmt.Print("Input the word number you want to delete from string ->")
				fmt.Scan(&wordN)
				f.Close()
				
				fin,err:=os.Open("input115.txt")
				if err!=nil{
					fmt.Print("File write error")
					return
				}
				defer fin.Close()
				scanner2:=bufio.NewScanner(fin)
				slStr:=make([]string,FileLenght)//Filelenght=8 (const)
				sN:=0//номер строчки
				
				for scanner2.Scan(){
					sN++
					sT:=scanner2.Text()//строчка
					slStr=append(slStr,sT)//слайс со всеми строчками
					if sN==strNumber{ //находим нужную нам строчку
						slStrWord:=strings.Split(slStr[sN-1]," ") //разбиваем ее на слова
						for i:=0;i<len(slStrWord);{ //избавляемся от пустот
							if slStrWord[i]==""{
								slStrWord=append(slStrWord[:i],slStrWord[i+1:]...) 
							}else{
								i++
							}
						}
						slStrWord=append(slStrWord[:wordN],slStrWord[wordN+1:]...) //удаляем нужное нам слово
						fmt.Print(slStrWord)
					}
				}	
				writer:=bufio.NewWriter(fin) //печатаем тот же файл, но с удаленным словом в выбранной строчке
				for j:=range slStr{
					fmt.Fprint(writer, slStr[j])
				}
				fmt.Print("Check the file")
				break
			}
			if opN<1 || opN>2{
				fmt.Print("ERROR. INVALID OPERAATION.")
			}
			
		}
		
	}
}










