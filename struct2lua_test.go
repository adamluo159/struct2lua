package struct2lua

import (
	"fmt"
	"testing"
)

type CWD struct {
	A string
	B int
	C []int
	D []string
	E bool
}

type LUA struct {
	CW CWD
	ID int
	IP string
}

func TestToLuaConfig(t *testing.T) {
	g := LUA{
		CW: CWD{
			A: "aaaaa",
			B: 123,
			C: []int{1, 2, 3},
			D: []string{"dfdfdfddf", "5", "6"},
			E: true,
		},
		ID: 1,
		IP: "192.168.1.1",
	}
	result := ToLuaConfig(g)
	fmt.Println(result)
}
