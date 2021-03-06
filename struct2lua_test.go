package struct2lua

import "testing"

type T struct {
	A string
	B int
	C []int
	//D []string
	E bool
}
type R struct {
	F string
}

type LUA struct {
	CW map[string]interface{}
	ID int
	IP string
	t  T
}

type Ld struct {
	NET_TIMEOUT_MSEC  int
	NET_MAX_CONNETION int
	StartService      []int
}

func TestToLuaConfig(t *testing.T) {
	g := LUA{
		t: T{
			A: "aaaaa",
			B: 123,
			C: []int{1, 2, 3},
			//D: []string{"dfdfdfddf", "5", "6"},
			E: true,
		},
		CW: make(map[string]interface{}),
	}

	g.CW["aaa"] = R{
		F: "a is a",
	}
	g.CW["bbb"] = T{
		A: "is a",
		B: 2222,
		E: true,
	}
	g.CW["ccc"] = "realccc"

	submap := []int{1, 2}
	submap[0] = 1
	g.CW["sub"] = submap

	//array := []R{
	//	{
	//		F: "aaaaa",
	//	},
	//	{
	//		F: "ccccc",
	//	},
	//}
	array := make([]R, 1)
	array[0] = R{F: "ccwwr"}

	g.CW["wwww"] = array

	lst := Ld{
		StartService:      []int{0, 1, 3},
		NET_TIMEOUT_MSEC:  100,
		NET_MAX_CONNETION: 300,
	}

	sucess := ToLuaConfig("", "testlua", g, lst, 1)

	if sucess == false {
		t.Error("test TestToLuaConfig fail~")
	}

}
