/*
6. Вводится натуральное число от 1 до 999. Вывести это число словами
*/
package main
import "fmt"
import "time"
func main(){
	var s string // число словами
	var num int
	fmt.Print("Input natural number n, 0>n>=999 ->")
	fmt.Scan(&num)
	u:=num%10
	num/=10
	d:=num%10
	h:=num/10
	switch h{
	case 1: s+="сто "
	case 2: s+="двести "
	case 3:	s+="триста "
	case 4:	s+="четыреста "
	case 5:	s+="пятьсот "
	case 6:	s+="шетьсот "
	case 7: s+="семьсот "
	case 8:	s+="восемьсот "
	case 9:	s+="девятьсот "
	}
	switch d{
	case 1: 
		switch u{
		case 1: s+="одиннадцать "
		case 2: s+="двенадцать "
		case 3: s+="тринадцать "
		case 4:	s+="четырнадцать "
		case 5: s+="пятнадцать "
		case 6:	s+="шестнадцать "
		case 7:	s+="семнадцать "
		case 8:	s+="восемнадцать "
		case 9: s+="девятнадцать "
		default: s+="десять "
		}
		fmt.Print(s)
		time.Sleep(3*time.Second)
		return
	case 2: s+="двадцать "
	case 3: s+="тридцать "
	case 4: s+="сорок "
	case 5:	s+="пятьдесят "
	case 6: s+="шестьдесят "
	case 7:	s+="семьдесят "
	case 8:	s+="восемьдесят "
	case 9:	s+="девяносто "
	}
	switch u{
	case 0: 
		if h==0 && d==0{
			s+="ноль"
		}
	case 1: s+="один"
	case 2: s+="два"
	case 3: s+="три"
	case 4: s+="четыре"
	case 5:	s+="пять"
	case 6: s+="шесть"
	case 7:	s+="семь"
	case 8:	s+="восемь"
	case 9:	s+="девять"
	}
	fmt.Print(s)
	time.Sleep(3*time.Second)
}
