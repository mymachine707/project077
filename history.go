

func Calculate(num1, num2, signs string) string {

	a := strings.Split(num1, "")
	b := strings.Split(num2, "")

	// step-1: validation is sign number  == negative (true)  // positive (false)
	var a1 bool // if number one is negative a1 is true else a1 is false
	var b2 bool // if number two is negative b2 is true else b2 is false
	//
	a1 = Signnumbers(a)
	b2 = Signnumbers(b)
	// clean up sign !!!!!
	num1 = Clean(num1)
	num2 = Clean(num2)
	//=====//step-1-2: validation is number true else panic
	Isnumber(a)
	Isnumber(b)

	// step-2: validation is sign canculate // Bajarish amalini tekshiriladi +  -  *  / is sign true else panic
	SignValidation(signs)

	var answerFunction string // function answer

	//=====//step-2-1: sign numbers
	switch signs {
	case "+":
		fmt.Println("+")
		// step-3: yig'indini hisoblash
		//=====//step-3-1: a+b=a+b
		if !a1 && !b2 {
			// a+b=a+b
			// a > 0  and b > 0
			answerFunction = Add(num1, num2)
		} else if !a1 && b2 {
			//=====//step-3-2: a+(-b)=a-b  // kattasini aniqlab keyin ayiriladi
			// a > 0  and b < 0
			if Findtobig(num1, num2) == 1 { // a>b  == a-b
				answerFunction = Minus(num1, num2)
			} else if Findtobig(num1, num2) == 2 { // a<b == b-a  // answer is negative
				answerFunction = AddtoMinus(Minus(num1, num2)) // add to minus for answer
			}
		} else if a1 && !b2 {
			// a < 0  and b > 0
			//=====//step-3-3: -a+b=b-a  // kattasini aniqlab keyin ayiriladi
			if Findtobig(num1, num2) == 1 { // a>b  == a-b = -c
				answerFunction = AddtoMinus(Minus(num1, num2)) // add to minus for answer
			} else if Findtobig(num1, num2) == 2 { // a<b == b-a
				answerFunction = Minus(num1, num2)
			}
		} else if a1 && b2 {
			//=====//step-3-4: -a+(-b)=-(a+b)
			answerFunction = AddtoMinus(Add(num1, num2))
		}

	case "-":
		fmt.Println("-")
		// step-4: ayirmani hisoblash
		//=====//step-4-1: a-b=a-b // kattasini aniqlab keyin ayiriladi
		if !a1 && !b2 {
			if Findtobig(num1, num2) == 1 { // a>b
				answerFunction = Minus(num1, num2)
			} else if Findtobig(num1, num2) == 2 { //b>a
				answerFunction = AddtoMinus(Minus(num1, num2))
			}
		} else if !a1 && b2 {
			//=====//step-4-2: a-(-b)=a+b
			answerFunction = Add(num1, num2)
		} else if a1 && !b2 {
			//=====//step-4-3: -a-b=-(a+b)
			answerFunction = AddtoMinus(Add(num1, num2))
		} else if a1 && b2 {
			//=====//step-4-4: -a-(-b)=b-a // kattasini aniqlab keyin ayiriladi
			if Findtobig(num1, num2) == 2 { // a>b
				answerFunction = Minus(num1, num2)
			} else if Findtobig(num1, num2) == 1 { //b>a
				answerFunction = AddtoMinus(Minus(num1, num2))
			}
		}
	case "*":
		fmt.Println("*")
		// step-5: ko'paytmani hisoblash
		if (!a1 && !b2) || (a1 && b2) {
			//=====//step-5-1: a*b=a*b // kattasini aniqlab keyin ayiriladi
			//=====//step-5-3: (-a)*(-b)=a*b
			answerFunction = Multiply(num1, num2)
		} else if (a1 && !b2) || (!a1 && b2) {
			//=====//step-5-2: a*(-b)=-(a*b) \\ // -a*b=-(a*b)
			answerFunction = AddtoMinus(Multiply(num1, num2))
		}
	case "/":
		fmt.Println("/")
		// step-6: bo'linmani hisoblash
	case "%":
		fmt.Println("%")
		// step-7: qoldiqni hisoblash
	default:
		panic("This sign is not defined")
	}

	return answerFunction
}

// Function of add numbers
func Add(num1, num2 string) string {

	a := strings.Split(num1, "")
	b := strings.Split(num2, "")
	// kiruvchi stringlarni uzunligini tenglashtirish======================

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

	var answer string // javobni ansverga biriktirish
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
		answer += Int_str(k[i])
	}

	return answer
}

func Findtobig(num1, num2 string) int {

	a := strings.Split(Clean(num1), "")
	b := strings.Split(Clean(num2), "")

	var e int // a katta b dan
	var c int // a kichik b dan
	// qaysi son kattaligini tekshirish kerak
	if len(a) > len(b) {
		// a katta b dan
		e += 1
	} else if len(a) < len(b) {
		// a kichik b dan
		c += 1
	} else if len(a) == len(b) {
		// a teng b ga
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				// a katta b dan
				e += 1
				break
			} else if a[i] < b[i] {
				// a kichik b dan
				c += 1
				break
			}

		}

	}
	var answer int
	if e == 1 {
		answer = 1
	} else if c == 1 {
		answer = 2
	}
	return answer
}

func Minus(num1, num2 string) string {
	a := strings.Split(num1, "")
	b := strings.Split(num2, "")

	// kiruvchi stringlarni uzunligini tenglashtirish======================

	var s []string
	var y int

	var e int // a katta b dan
	var c int // a kichik b dan

	if Findtobig(num1, num2) == 1 {
		e = 1
	} else if Findtobig(num1, num2) == 2 {
		c = 1
	}

	// slicelarni bir biriga tenglashtirish jarayoni bu kod o'zgartirilmasin

	if e == 1 { // a katta b dan
		y = len(a) - len(b)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, b...)
		b = s
		s = nil

	} else if c == 1 { // a kichik b dan
		y = len(b) - len(a)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, a...)
		a = s
		s = nil
	}
	//==========================

	// asosiy shartlar
	var z []string

	for i := 0; i < len(a); i++ {
		z = append(z, a[i])
	}

	//a bilan b berilgan slice holatida
	if e == 1 {
		fmt.Println("\nnumber1 > number 2")
		// a > b bo'lsa:
		for i := len(a); i > 0; i-- {
			if Str_Int(a[i-1])-Str_Int(b[i-1]) > 0 { // Agar joriy indekslar ayirmasi musbat bo'lsa
				z[i-1] = Int_str(Str_Int(a[i-1]) - Str_Int(b[i-1])) // z joriy indeksiga ayirma qiymati beriladi
			} else if Str_Int(a[i-1])-Str_Int(b[i-1]) == 0 { // Agar joriy indekslar ayirmasi 0 ga teng bo'lsa
				z[i-1] = Int_str(0) // z joriy indeksiga 0 qiymat beriladi
			} else if Str_Int(a[i-1])-Str_Int(b[i-1]) < 0 { // agar joriy indekslar ayirmasi 0 dan kichik bo'lsa
				a[i-1] = Int_str(Str_Int(a[i-1]) + 10)              // katt sonning joriy indeksiga 10 soni qo'shiladi va qaytadan yenglashtiriladi
				z[i-1] = Int_str(Str_Int(a[i-1]) - Str_Int(b[i-1])) // z ning joriy qiymatini sonlarning joriy qiymatlari ayirmasiga tenglashtiriladi va...
				if Str_Int(a[i-2]) > 0 {                            // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 dan katta bo'lsa
					a[i-2] = Int_str(Str_Int(a[i-2]) - 1) // katta sonning joriy qiymatidan bitta oldingi sonnini bittaga kamaytiriladi
				} else if Str_Int(a[i-2]) == 0 { // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 ga teng bo'lsa...
					a[i-2] = "9"           // katta sonning joriy qiymatidan bitta oldingi sonnini 9 soniga tenglashtiriladi
					var boolen bool = true // for loop uchun boolen qiyamtini true holatida ochildi
					var m int = 3          // indekslarni tekshirish uchun m int shaklida ochildi va katta sonning joriy qiymatidan ikkita oldingi qiymatini tekshirish uchun m sonini 3 ga tenglandi
					for boolen {           // for loop ochildi
						if Str_Int(a[i-m]) > 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 dan katta bo'lsa
							a[i-m] = Int_str(Str_Int(a[i-m]) - 1) // katta sonning joriy indeksidan ikkita oldingi indeksidan 1 soni ayiriladi
							boolen = false                        // boolen false qiymatini oladi
							break                                 // for loopdan chiqiladi
						} else if Str_Int(a[i-m]) == 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 ga teng bo'lsa
							a[i-m] = "9" // Katta sonning joriy indeksidan ikkita oldingi indeksi "9" ga tenglanadi
							m++          // m soni 1 ga oshiriladi va for loop qaytadan tekshiruvni boshlaydi
						}

					}
				}
			}

		}

	} else if c == 1 {
		fmt.Println("\nnumber1 < number 2")
		// a < b bo'lsa:
		for i := len(b); i > 0; i-- {
			if Str_Int(b[i-1])-Str_Int(a[i-1]) > 0 { // Agar joriy indekslar ayirmasi musbat bo'lsa
				z[i-1] = Int_str(Str_Int(b[i-1]) - Str_Int(a[i-1])) // z joriy indeksiga ayirma qiymati beriladi
			} else if Str_Int(b[i-1])-Str_Int(a[i-1]) == 0 { // Agar joriy indekslar ayirmasi 0 ga teng bo'lsa
				z[i-1] = Int_str(0) // z joriy indeksiga 0 qiymat beriladi
			} else if Str_Int(b[i-1])-Str_Int(a[i-1]) < 0 { // agar joriy indekslar ayirmasi 0 dan kichik bo'lsa
				b[i-1] = Int_str(Str_Int(b[i-1]) + 10)              // katt sonning joriy indeksiga 10 soni qo'shiladi va qaytadan yenglashtiriladi
				z[i-1] = Int_str(Str_Int(b[i-1]) - Str_Int(a[i-1])) // z ning joriy qiymatini sonlarning joriy qiymatlari ayirmasiga tenglashtiriladi va...
				if Str_Int(b[i-2]) > 0 {                            // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 dan katta bo'lsa
					b[i-2] = Int_str(Str_Int(b[i-2]) - 1) // katta sonning joriy qiymatidan bitta oldingi sonnini bittaga kamaytiriladi
				} else if Str_Int(b[i-2]) == 0 { // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 ga teng bo'lsa...
					b[i-2] = "9"           // katta sonning joriy qiymatidan bitta oldingi sonnini 9 soniga tenglashtiriladi
					var boolen bool = true // for loop uchun boolen qiyamtini true holatida ochildi
					var m int = 3          // indekslarni tekshirish uchun m int shaklida ochildi va katta sonning joriy qiymatidan ikkita oldingi qiymatini tekshirish uchun m sonini 3 ga tenglandi
					for boolen {           // for loop ochildi
						if Str_Int(b[i-m]) > 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 dan katta bo'lsa
							b[i-m] = Int_str(Str_Int(b[i-m]) - 1) // katta sonning joriy indeksidan ikkita oldingi indeksidan 1 soni ayiriladi
							boolen = false                        // boolen false qiymatini oladi
							break                                 // for loopdan chiqiladi
						} else if Str_Int(b[i-m]) == 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 ga teng bo'lsa
							b[i-m] = "9" // Katta sonning joriy indeksidan ikkita oldingi indeksi "9" ga tenglanadi
							m++          // m soni 1 ga oshiriladi va for loop qaytadan tekshiruvni boshlaydi
						}

					}
				}
			}

		}
	}
	//natijani birlashtirish
	var answers string
	for i := 0; i < len(z); i++ {
		answers += z[i]
	}

	//==========================
	return Clean(answers)
}

// stringdan intga o'tqizadi
func Str_Int(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 0 // sho'tda tanish bilishchili bo'm ketti err bersa 0 qayatrmasligi kere boshqa javob so'ra
	}
	return intVar
}

// intdan stringga o'tkizadi

func Int_str(num int) string {
	str := strconv.Itoa(num)
	return str
}

func Clean(num string) string {
	a := strings.Split(num, "")
	//

	if a[0] == "+" {
		a = a[1:]
	} else if a[0] == "-" {
		a = a[1:]
	} else if !IsnumberStr(a[0]) {
		panic("This string is not a number!")
	}
	// o'zgartirish kiritish kerak!
	for i := 1; i < len(a); i++ {
		if !IsnumberStr(a[i]) {
			panic("This string is not a number!")
		}
	}

	var b []string
	count := 0

	for i := 0; i < len(a); i++ {

		if a[i] == "0" {
			count++
		} else {
			break
		}
	}
	b = a[count:]

	var answer string

	for i := 0; i < len(b); i++ {
		answer += b[i]
	}

	return answer
}

// slice is number
func Isnumber(num []string) bool {
	var answer bool

	for i := 0; i < len(num); i++ {
		if num[0] == "+" || num[0] == "-" {
			continue
		} else if IsnumberStr(num[i]) {
			answer = true
		} else if !IsnumberStr(num[i]) {
			panic("This string is not a number!")
		}
	}
	return answer
}

// entered string and validate
func IsnumberStr(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func Multiply(a, b string) string {
	var k []string

	var e int // a katta b dan
	var c int // a kichik b dan
	// qaysi son kattaligini tekshirish kerak
	if len(a) >= len(b) {
		// a katta b dan
		e += 1

	} else if len(a) < len(b) {
		// a kichik b dan
		c += 1

	}
	var ans string
	if e == 1 {
		f := strings.Split(b, "")
		for i := len(f); i > 0; i-- {
			ans = mult(a, f[i-1])
			for l := 0; l < len(f)-i; l++ {
				ans += "0"
			}

			k = append(k, ans)
		}

	} else if c == 1 {
		f := strings.Split(a, "")
		for i := len(f); i > 0; i-- {
			ans = mult(b, f[i-1])
			for l := 0; l < len(f)-i; l++ {
				ans += "0"
			}
			k = append(k, ans)

		}

	}

	var answers string

	for i := 0; i < len(k); i++ {
		answers = Add(answers, k[i])
	}

	return answers
}

// katta xonali sonni bir xonali songa ko'paytirish

func mult(w, b string) string {

	a := strings.Split(w, "")

	var l []string = a

	var p int
	var m int = 0
	for i := len(a); i > 0; i-- {
		p = Str_Int(a[i-1]) * Str_Int(b)
		if i-1 == 0 {
			p = Str_Int(a[i-1]) * Str_Int(b)
			p += m
			m = 0
			l[i-1] = Int_str(p)
		} else if p < 10 {
			if p+m >= 10 {
				p += m
				m = 0
				q := strings.Split(Int_str(p+m), "")
				l[i-1] = q[1]
				m = Str_Int(q[0])
			} else {
				l[i-1] = Int_str(p)
			}
		} else if p >= 10 {
			p += m
			m = 0
			q := strings.Split(Int_str(p), "")
			l[i-1] = q[1]
			m = Str_Int(q[0])
		}
	}
	answer := ""

	for i := 0; i < len(l); i++ {
		answer += l[i]
	}
	return answer
}

func Signnumbers(num []string) bool {

	//
	var negative bool

	if num[0] == "+" {
		negative = false
	} else if num[0] == "-" {
		negative = true
	} else if IsnumberStr((num[0])) {
		negative = false
	} else {
		panic("This string is not num number!{Signnumbers}")
	}

	return negative
}

func SignValidation(sign string) string {
	var answer string
	switch sign {
	case "+":
		answer = sign
	case "-":
		answer = sign
	case "*":
		answer = sign
	case "/":
		answer = sign
	case "%":
		answer = sign
	default:
		panic("This sign is not defined!")
	}
	return answer
}

func AddtoMinus(num string) string {
	var b []string
	b = append(b, "-")
	b = append(b, num)
	var answer string
	for i := 0; i < len(b); i++ {
		answer += b[i]
	}
	return answer

}
