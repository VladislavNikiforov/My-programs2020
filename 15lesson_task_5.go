/*5. Функция получает номер месяца (вариант: название или краткое трёхбуквенное название).
 Возвращает количество дней в этом месяце. */
package main
import "fmt"
func DayAmo(num int)int{
	switch num{
	case 2: return 28
	case 4,6,9,11: return 30
	default: return 31	
	}
}
func main(){
	var num int
	fmt.Print("Input a number of a month ->")
	fmt.Scan(&num)
	fmt.Printf("The amount of days in this month is %d",DayAmo(num))
}
