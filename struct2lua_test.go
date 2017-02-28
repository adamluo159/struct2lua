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

func TestToLuaConfig(t *testing.T) {
	g := LUA{
		t: T{
			A: "aaaaa",
			B: 123,
			C: []int{1, 2, 3},
			//D: []string{"dfdfdfddf", "5", "6"},
			E: true,
		},
		//		ID: 1,
		//IP: "192.168.1.1",
		CW: make(map[string]interface{}),
	}

	//d := R{
	//	F: "luoluoj",
	//}

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

	sucess := ToLuaConfig("", "testlua", g)

	if sucess == false {
		t.Error("test TestToLuaConfig fail~")
	}
}
