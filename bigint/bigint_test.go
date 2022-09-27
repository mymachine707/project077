package bigint


import 	"testing"

func TestMultiply(t *testing.T)  {
	num1:=Bigint{Value: "123"}
	num2:=Bigint{Value: "321"}

	num3:=Bigint{Value: "-00000000000000000000000000000000000198"}

	Multiplyans:=Bigint{Value: "39483"}
	AddAns:=Bigint{Value: "444"}
	Minusans:=Bigint{Value: "198"}
	clear:=Bigint{Value: "-198"}

	if Multiply(num1,num2)!=Multiplyans{
		t.Error("It's false")
	}
	if Add(num1,num2)!=AddAns{
		t.Error("It's false")
	}
	if Minus(num1,num2)!=Minusans{
		t.Error("It's false")
	}
	if Clean(num3)!=clear{
		t.Error("It's false")
	}
}
// go test bigint/bigint_test.go