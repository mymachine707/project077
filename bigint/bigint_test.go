package bigint


import 	"testing"

func TestMultiply(t *testing.T)  {
	num1:=bigint.Bigint{Value: "123"}
	num2:=bigint.Bigint{Value: "321"}

	num3:=bigint.Bigint{Value: "-00000000000000000000000000000000000198"}

	Multiplyans:=bigint.Bigint{Value: "39483"}
	AddAns:=bigint.Bigint{Value: "444"}
	Minusans:=bigint.Bigint{Value: "198"}
	Clean:=bigint.Bigint{Value: "-198"}

	if bigint.Multiply(num1,num2)!=Multiplyans{
		t.Error("It's false")
	}
	if bigint.Add(num1,num2)!=AddAns{
		t.Error("It's false")
	}
	if bigint.Minus(num1,num2)!=Minusans{
		t.Error("It's false")
	}
	if bigint.Clean(num1,num2)!=Minusans{
		t.Error("It's false")
	}
}
// go test bigint/bigint_test.go