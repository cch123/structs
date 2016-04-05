package skiplist

import "testing"

var strArr = []string{"a", "b", "name", "xargin"}

func compare(a interface{}, b interface{}) int {
	if a.(string) > b.(string) {
		return 1
	} else if a.(string) < b.(string) {
		return -1
	}
	return 0
}

func TestInsert(t *testing.T) {
	var x *Skiplist = SkiplistCreate(compare)
	var insRes *SkiplistNode
	for _, str := range strArr {
		insRes = x.SkiplistInsert(str)
		if insRes == nil {
			t.Errorf("insert %s failed\n", str)
		}
	}
}

func TestDelete(t *testing.T) {
	var x *Skiplist = SkiplistCreate(compare)
	var insRes *SkiplistNode
	for _, str := range strArr {
		insRes = x.SkiplistInsert(str)
		if insRes == nil {
			t.Errorf("insert %s failed\n", str)
		}
	}

	//x.Print()
	var lastLength = x.Length
	var delRes int
	for _, str := range strArr {
		delRes = x.SkiplistDelete(str)
		if delRes != 1 || lastLength != x.Length+1 {
			t.Errorf("delete str %s error !\n", str)
		}
		lastLength = x.Length
		//x.Print()
	}

	if x.Length != 0 {
		t.Errorf("after delete all there cannot be any element left\n")
	}
}

func TestPopHead(t *testing.T) {
}

func TestPopTail(t *testing.T) {
}

func TestLength(t *testing.T) {
}

func TestCreate(t *testing.T) {
}

func TestDuplicateNode(t *testing.T) {
}
