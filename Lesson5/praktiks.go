package Lesson5

func lesson() {
	// 1
	// var name string
	// name = "NIikta"
	// fmt.Println(name)
	//
	//
	// 2
	// const pi = 3.1415926
	// fmt.Println(pi)
	//
	//
	// 3
	// var a, b int = 6, 4
	//
	//	if a > b {
	//		fmt.Println("Наибольшее число", a)
	//	}
	//
	//
	// 4
	// var a int = 67
	//
	//	if a%2 == 0 {
	//		fmt.Println("это число четное", a)
	//	} else {
	//
	//		fmt.Println("Это число нечетное", a)
	//	}
	//
	//
	// 5
	// age := 5
	//
	//	if age >= 18 {
	//		fmt.Println("Ты взрослый")
	//	} else {
	//
	//		fmt.Println("Ты маленький")
	//	}
	//
	// 6
	//
	//
	//	for i := 1; i <= 10; i++ {
	//		fmt.Println(i)
	//	}
	//
	//
	// 7
	// x := 1
	//
	//	for x < 9 {
	//		fmt.Println(x)
	//		x++
	//	}
	//
	//
	// 8
	// day := 889
	// switch day {
	// case 1:
	//
	//	fmt.Println("понедельник")
	//
	// case 2:
	//
	//	fmt.Println("вторник")
	//
	// case 3:
	//
	//	fmt.Println("среда")
	//
	// case 4:
	//
	//	fmt.Println("четверг")
	//
	// case 5:
	//
	//	fmt.Println("пятница")
	//
	// case 6:
	//
	//	fmt.Println("суббота")
	//
	// case 7:
	//
	//	fmt.Println("воскресенье")
	//
	// default:
	//
	//		fmt.Println("такого дня нет")
	//	}
	//
	//
	// 9-10
	//
	//		func divide (a,b int)(result int, err error){
	//		if b == 0{
	//			err = fmt.Errorf("клкние на ноль")
	//			return
	//		}
	//		result = a / b
	//		return
	//	}
	//
	//
	//11
	//
	//for i := 12; i <= 45; i++ {
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	//
	//
	//12
	//for i := 1; i <= 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//
	//
	//13
	//var i int
	//fmt.Print("Введитt число ")
	//fmt.Scanln(&i)
	//fmt.Println("Ваше число", i)
	//
	//
	//14
	//var name string = "Никита"
	//age := 29
	//
	//fmt.Printf("Привет меня завут %s, мне %d лет/n", name, age)
	//
	//
	//15
	//a, b, c := 3, 7, 32
	//fmt.Println(a + b + c)
	//
	//
	//16
	//func math(){
	//	var a, b int
	//	fmt.Print("Введите числа ")
	//	fmt.Scanln(&a, &b)
	//	c := a + b
	//	fmt.Printf("Ваше число %d", c)
	//}
	//
	//
	//17
	//var number int
	//fmt.Println("Введите число: ")
	//fmt.Scanln(&number)
	//
	//isGreaterTen := number > 10
	//
	//fmt.Printf("Число больше 10? %t\n", isGreaterTen)
	//
	//
	//18
	//number := -12
	//if number < 0 {
	//	fmt.Println("Число меньше нуля")
	//} else if number == 0 {
	//	fmt.Println("Равно нулю")
	//} else {
	//	fmt.Println("Число больше нуля")
	//}
	//
	//
	//19
	//if x := 20; x >= 20 {
	//	fmt.Println("x больше 20")
	//} else {
	//	fmt.Println("х больше или равно 20")
	//}
	//
	//
	// 20
	//func divide(a, b int) (result int, err error) {
	//	if b == 0 {
	//		err = fmt.Errorf("деление на ноль")
	//		return
	//	}
	//	result = a / b
	//	return
	//}
	//
	//func main() {
	//	result, err := divide(30, 3)
	//
	//	if err != nil {
	//		fmt.Println("Ошибка:", err)
	//	} else {
	//		fmt.Println("Результат:", result)
	//	}
	//}
}
