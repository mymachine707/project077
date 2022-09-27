package bigint

import (
"unicode"
	"strconv"
	"strings"
	"errors"
)

type Bigint struct {
	Value string
}
// Ro'yxatda bor funksiyalar:

// Multiply
// Vmult
// Add
// Minus
// bigStrf
// minuss
// addZero
// zeroAdd
// Str_Int
// Int_Str
// clean up
// sign validation
// AddtoMinus


func Validationstr( num string ) error {
	ErrorValidation:=errors.New("Invalid number!")
	start:=0
	if string(num[0])=="-" || string(num[0])=="+"{
	
	start = 1
		} 
	allowed := "1234567890"

	for i := start; i < len(num); i++ {
		if !strings.Contains(allowed, string(num[i])){
			return ErrorValidation
		}
}
	return nil



}


func Clean(num Bigint) Bigint {
	num1:= num.Value
	if err:=Validationstr(num1); err!=nil{
		panic(err)
	}
	start:=0

// 	a := strings.Split(num.Value, "")
// 	if a[0]=="-"{
// 		a=a[1:]
// 		m=1
// 	}
// 	if a[0]=="+"{
// 		a=a[1:]
// 	}


// 	for i := 0; i < len(a); i++ {
// 		if a[i]!="0"{
// 			a=a[i:]
// 			break
// 		}
// 		if a[i]=="0" && i+1< len(a){
// 			a=a[i+1:]

// 		}
// 	}
// 	var answer Bigint

// 	if m==1{
// 		answer.Value+="-"
// 	}
// 	fmt.Println(a)
// 	for i := 0; i < len(a); i++ {
// 		answer.Value+=a[i]
// 	}

// 	return answer

	if string(num1[0])=="-" {
		num1=num1[1:]
		start=1
	}
	if string(num1[0])=="+" {
		num1=num1[1:]
	}

	for strings.HasPrefix(num1, "0"){
		num1=num1[1:]
	}

	if start==1{
		num1="-"+num1	
	}

	

	return Bigint{Value: num1}

}

func addMinus(sign Bigint) Bigint{
	var answer Bigint
	
	answer.Value+="-"
	answer.Value+=sign.Value

return answer
}

func SignValidation(sign Bigint) Bigint{
	var answer Bigint
	switch sign.Value {

		case "+":
			answer=Bigint{Value: sign.Value}
		case "-":
			answer=Bigint{Value: sign.Value}
		case "*":
			answer=Bigint{Value: sign.Value}
		case "/":
			answer=Bigint{Value: sign.Value}
		case "%":
			answer=Bigint{Value: sign.Value}
	default:
			answer.Value= "Error: You cannot use this sign!"
			break
	}
return answer
}

func IsnumberStr(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}


func Multiply(num1, num2 Bigint) Bigint {
	a := strings.Split(num1.Value, "")
	b := strings.Split(num2.Value, "")

	var k []string

	bigstr := bigStrf(a, b)

	if bigstr == 1 || bigstr == 0 {
		for i := len(b); i > 0; i-- {
			k = append(k, zeroAdd(Vmult(num1.Value, b[i-1]), len(b)-i))
		}
	}
	if bigstr == -1 || bigstr == 0 {
		for i := len(a); i > 0; i-- {
			k = append(k, zeroAdd(Vmult(num2.Value, a[i-1]), len(a)-i))
		}
	}

	var answers Bigint

	for i := 0; i < len(k); i++ {
		m := Bigint{Value: k[i]}
		answers = Add(answers, m)
	}

	return answers
}

// katta xonali sonni bir xonali songa ko'paytirish
func Vmult(w, b string) string {

	a := strings.Split(w, "")
	var l []string
	l = append(l, a...)
	t := Str_Int(b)
	reminder := 0
	answer := ""
	k := 0
	for i := len(a); i > 0; i-- {
		if i-1 == 0 {
			k += Str_Int(a[i-1])*t + reminder
			l[i-1] = Int_str(k)
			k = 0
		}
		if Str_Int(a[i-1])*t+reminder < 10 {
			k += Str_Int(a[i-1])*t + reminder
			l[i-1] = Int_str(k)
			k = 0
			reminder = 0
		}
		if Str_Int(a[i-1])*t+reminder > 10 {
			k += Str_Int(a[i-1])*t + reminder
			reminder = k / 10
			l[i-1] = Int_str(k % 10)
			k = 0
		}
	}
	// a*b

	for i := 0; i < len(l); i++ {
		answer += l[i]
	}
	return answer
}

func Add(num1, num2 Bigint) Bigint {
	a := strings.Split(num1.Value, "")
	b := strings.Split(num2.Value, "")
	// kiruvchi stringlarni uzunligini tenglashtirish

	var s []string
	var y int

	if len(a) > len(b) {
		y = len(a) - len(b)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, b...)
		b = s
		s = nil
	} else {
		y = len(b) - len(a)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, a...)
		a = s
		s = nil
	}

	//stringlarni slice holatida har bir elementini qo'shib yangi slice hosil qilish

	var answer Bigint // javobni ansverga biriktirish
	var k []int

	var m int
	for i := 0; i < len(a); i++ {
		z := Str_Int(a[i]) + Str_Int(b[i])
		k = append(k, z)

		for i := len(k); i > 0; i-- {
			if k[i-1] > 10 {
				if i-1 > 0 {
					m = k[i-1] - 10
					k[i-1] = m
					if i-2 >= 0 {
						k[i-2]++
					}
				} else if i-1 == 0 {
					break
				}

			}
		}
	}

	for i := 0; i < len(k); i++ {
		answer.Value += Int_str(k[i])
	}

	return answer
}

func Minus(num1, num2 Bigint) Bigint {
	a := strings.Split(num1.Value, "")
	b := strings.Split(num2.Value, "")
	// kiruvchi stringlarni uzunligini tenglashtirish

	var bigStr int = bigStrf(a, b)
	// kattasi aniqlandi
	var k []string
	var answers Bigint
	if bigStr == 1 { // if a is big then b
		b = addZero(a, b)
		k = minuss(a, b)
	}
	if bigStr == -1 {
		a = addZero(b, a)
		k = minuss(b, a)
	}
	if bigStr == 0 {
		answers.Value = "0"
	}
	// 0-9 [i-1]

	// 0==9 [i-2]



	for i := 0; i < len(k); i++ {
		answers.Value += k[i]
	}

	return answers
}

func bigStrf(a, b []string) int {
	var bigStr int
	// kattasi aniqlandi
	if len(a) > len(b) {
		bigStr = 1
	}
	if len(a) < len(b) {
		bigStr = -1
	}
	if len(a) == len(b) {
		count := 0
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				bigStr = 1
				break
			}
			if a[i] < b[i] {
				bigStr = -1
				break
			}
			count++
		}
		if len(a)==count{
			bigStr = 0
		}
	}
	return bigStr
}

func minuss(a, b []string) []string {
	var answer []string
	answer = append(answer, a...)
	reminder := 0
	for i := len(a); i > 0; i-- {
		if a[i-1] > b[i-1] {
			answer[i-1] = Int_str(Str_Int(a[i-1]) - Str_Int(b[i-1]) - reminder)
			reminder = 0
		}
		if a[i-1] == b[i-1] {
			if reminder != 1 {
				answer[i-1] = "0"
			}
			if reminder == 1 {
				answer[i-1] = "9"
			}
		}
		if a[i-1] < b[i-1] {
			answer[i-1] = Int_str(10 + Str_Int(a[i-1]) - Str_Int(b[i-1]) - reminder)
			reminder = 1
		}
	}
	return answer
}

func addZero(a, b []string) []string {
	adLen := len(a) - len(b)
	var s []string
	for i := 0; i < adLen; i++ {
		s = append(s, "0")
	}
	s = append(s, b...)
	return s
}
func zeroAdd(a string, k int) string {
	var s string
	s += a
	for i := 0; i < k; i++ {
		s += "0"
	}
	return s
}

func Str_Int(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 0 // ko'ddi o'zgartiradigan joyi bor
	}
	return intVar
}

// intdan stringga o'tkizadi

func Int_str(num int) string {
	str := strconv.Itoa(num)
	return str
}
