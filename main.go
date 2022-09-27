package main

import (
	"bigint/bigint"
	"fmt"
)

func main() {
	//Bigint strukturada o'zgaruvchi qiymatlari olindi
	var num1 bigint.Bigint = bigint.Bigint{
		Value: "20",
	}
	var num2 bigint.Bigint = bigint.Bigint{
		Value: "100",
	}
	var sign bigint.Bigint = bigint.Bigint{
		Value: "-",
	}
/*
	var num3 bigint.Bigint = bigint.Bigint{
		Value: "+0000000000000000000000000010000002122000131564",
	}*/
	
	Calculates := bigint.Calculate(num1,num2,sign)
	fmt.Printf("num1: %v\n", num1)
	fmt.Printf("sign: %v\n", sign)
	fmt.Printf("num2: %v\n", num2)
	fmt.Printf("Calculates: %v\n", Calculates)  //
	
	// var sign bigint.Bigint = bigint.Bigint{
	// 	Value: "+",
	// }
	// num2.Value=num2.Value[1:]
	// fmt.Println(num2)
	// AddNumber := bigint.Add(num1,num2)
	// MultiplyNumber := bigint.Multiply(num1, num2)
	//Clenmethod := bigint.Clean(num3)
	
	// fmt.Printf("\nAddNumber: %v\n", AddNumber.Value)
	// fmt.Printf("MultiplyNumber: %v\n", MultiplyNumber.Value)
	//fmt.Printf("Clenmethod: %v\n", Clenmethod)
	
	// fmt.Printf("\nnum1: %v\n", num1.Value)
	// fmt.Printf("num2: %v\n", num2.Value)
	// fmt.Printf("sign: %v\n", sign.Value)
	// num2.Value=num2.Value[1:]
	// MinusNumber :=bigint.Minus(num1, num2)
	// fmt.Printf("MinusNumber: %v\n", MinusNumber.Value)
}
