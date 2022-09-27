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
		Value: "-100",
	}
	var add bigint.Bigint = bigint.Bigint{
		Value: "+",
	}
	var sub bigint.Bigint = bigint.Bigint{
		Value: "*",
	}

	var min bigint.Bigint = bigint.Bigint{
		Value: "-",
	}
	
/*
	var num3 bigint.Bigint = bigint.Bigint{
		Value: "+0000000000000000000000000010000002122000131564",
	}*/
	

	fmt.Printf("num1: %v\n", num1)
	fmt.Printf("num2: %v\n\n", num2)


	
	fmt.Printf("%s%s%s=%s\n",num1.Value,add.Value,num1.Value,bigint.Calculate(num1,num1,add).Value)
	fmt.Printf("%s%s(%s)=(%s)\n",num1.Value,add.Value,num2.Value,bigint.Calculate(num1,num2,add).Value)
	fmt.Printf("(%s)%s%s=(%s)\n",num2.Value,add.Value,num1.Value,bigint.Calculate(num2,num1,add).Value)
	fmt.Printf("(%s)%s(%s)=(%s)\n\n",num2.Value,add.Value,num2.Value,bigint.Calculate(num2,num2,add).Value)

	fmt.Printf("%s%s%s=%s\n",num1.Value,sub.Value,num1.Value,bigint.Calculate(num1,num1,sub).Value)
	fmt.Printf("%s%s(%s)=(%s)\n",num1.Value,sub.Value,num2.Value,bigint.Calculate(num1,num2,sub).Value)
	fmt.Printf("(%s)%s%s=(%s)\n",num2.Value,sub.Value,num1.Value,bigint.Calculate(num2,num1,sub).Value)
	fmt.Printf("(%s)%s(%s)=(%s)\n\n",num2.Value,sub.Value,num2.Value,bigint.Calculate(num2,num2,sub).Value)

	fmt.Printf("%s%s%s=%s\n",num1.Value,min.Value,num1.Value,bigint.Calculate(num1,num1,min).Value)
	fmt.Printf("%s%s(%s)=(%s)\n",num1.Value,min.Value,num2.Value,bigint.Calculate(num1,num2,min).Value)
	fmt.Printf("(%s)%s%s=(%s)\n",num2.Value,min.Value,num1.Value,bigint.Calculate(num2,num1,min).Value)
	fmt.Printf("(%s)%s(%s)=(%s)\n\n",num2.Value,min.Value,num2.Value,bigint.Calculate(num2,num2,min).Value)


}
