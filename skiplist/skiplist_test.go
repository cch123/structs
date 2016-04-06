package skiplist

import "testing"

var strArr = []string{"a", "b", "name", "xargin", "any", "one", "two", "three", "four"}

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

//TODO test random delete
func TestDelete(t *testing.T) {
	var x *Skiplist = SkiplistCreate(compare)
	var insRes *SkiplistNode
	for _, str := range strArr {
		insRes = x.SkiplistInsert(str)
		if insRes == nil {
			t.Errorf("insert %s for test delete failed\n", str)
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
	var x *Skiplist = SkiplistCreate(compare)
	var insRes *SkiplistNode
	for _, str := range strArr {
		insRes = x.SkiplistInsert(str)
		if insRes == nil {
			t.Errorf("insert %s for test pop head failed\n", str)
		}
	}

	//x.Print()

	// get all elements
	var skiplistStrs []string
	var tmpNode = x.Header
	for tmpNode = tmpNode.Level[0].Forward; tmpNode != nil; {
		skiplistStrs = append(skiplistStrs, tmpNode.Obj.(string))
		tmpNode = tmpNode.Level[0].Forward
	}

	for _, str := range skiplistStrs {
		var obj interface{} = x.SkiplistPopHead()
		if obj != nil && obj.(string) != str {
			t.Errorf("pop head error! the pop result is %v, expected is %s\n", obj, str)
		}
	}

	if x.Length != 0 {
		t.Errorf("after delete all there cannot be any element left\n")
	}
}

func TestPopTail(t *testing.T) {
	var x *Skiplist = SkiplistCreate(compare)
	var insRes *SkiplistNode
	for _, str := range strArr {
		insRes = x.SkiplistInsert(str)
		if insRes == nil {
			t.Errorf("insert %s for test pop head failed\n", str)
		}
	}

	//x.Print()
	// get all elements
	var skiplistStrs []string
	var tmpNode = x.Header
	for tmpNode = tmpNode.Level[0].Forward; tmpNode != nil; {
		skiplistStrs = append(skiplistStrs, tmpNode.Obj.(string))
		tmpNode = tmpNode.Level[0].Forward
	}

	//TODO 这里poptail应该倒排
	//for _, str := range skiplistStrs {
	for i := len(skiplistStrs) - 1; i >= 0; i-- {
		var str = skiplistStrs[i]
		var obj interface{} = x.SkiplistPopTail()
		if obj != nil && obj.(string) != str {
			t.Errorf("pop tail error! the pop result is %v, expected is %s\n", obj, str)
		}
		//x.Print()
	}

	if x.Length != 0 {
		t.Errorf("after delete all there cannot be any element left\n")
	}
}

func TestDuplicateNode(t *testing.T) {
	var x *Skiplist = SkiplistCreate(compare)
	var insRes *SkiplistNode
	for _, str := range strArr {
		insRes = x.SkiplistInsert(str)
		if insRes == nil {
			t.Errorf("insert %s for test pop head failed\n", str)
		}
	}

	for _, str := range strArr {
		insRes = x.SkiplistInsert(str)
		if insRes != nil {
			t.Errorf("node %s already exist, insert success is not valid\n", str)
		}
	}
}
