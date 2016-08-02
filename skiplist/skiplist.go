package skiplist

import "math/rand"
import "fmt"

const skiplistMaxLevel = 32
const skiplistTP = 0.25

// LevelElem is the level array of a skiplist node
type LevelElem struct {
	Forward *Node
	Span    uint64
}

//Node is the node of a skiplist
type Node struct {
	Obj      interface{}
	Backward *Node
	Level    [skiplistMaxLevel]LevelElem
}

// Compare is used to compare node elem
type Compare func(a interface{}, b interface{}) int

// Skiplist is the main struct of skiplist
type Skiplist struct {
	Header *Node
	Tail   *Node
	// compare function
	comp   Compare
	Length uint64
	Level  int
}

func createNode(level int, obj interface{}) *Node {
	var zn = new(Node)
	zn.Obj = obj
	return zn
}

// CreateList is used to create a skiplist
func CreateList(p Compare) *Skiplist {
	sl := new(Skiplist)
	sl.Level = 1
	sl.Length = 0
	sl.Header = createNode(skiplistMaxLevel, nil)
	for j := 0; j < skiplistMaxLevel; j++ {
		sl.Header.Level[j].Forward = nil
		sl.Header.Level[j].Span = 0
	}
	// in face,
	// in golang there is no need to initialize pointer to nil
	sl.Header.Backward = nil
	sl.Tail = nil
	sl.comp = p

	return sl
}

// the random number is fake random, cuz there is no time seed
// so in fact, the level of each insert node can be predicted
// but it doesn't matter
func skiplistRandomLevel() int {
	var level = 1
	for i, j := rand.Int()&0xFFFF, 0.25*0xFFFF; i < int(j); i = rand.Int() & 0xFFFF {
		level++
	}

	if level > skiplistMaxLevel {
		return skiplistMaxLevel
	}

	return level
}

//Insert is used to insert node to skiplist
func (sl *Skiplist) Insert(obj interface{}) *Node {

	// record the nodes which forward pointer need to be updated
	var update [skiplistMaxLevel]*Node
	var rank [skiplistMaxLevel]uint64

	var x *Node

	x = sl.Header

	for i := sl.Level - 1; i >= 0; i-- {
		if i == sl.Level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		// find the node before the insert node
		for x.Level[i].Forward != nil && sl.comp(x.Level[i].Forward.Obj, obj) < 0 {
			rank[i] += x.Level[i].Span
			x = x.Level[i].Forward
		}
		update[i] = x
	}

	// if the node exists, give up insert，and return nil pointer
	if x.Level[0].Forward != nil && sl.comp(x.Level[0].Forward.Obj, obj) == 0 {
		x = nil
		return x
	}

	level := skiplistRandomLevel()
	// if the generated level is bigger than the current skiplist level
	// need some special ops
	if level > sl.Level {
		for i := sl.Level; i < level; i++ {
			rank[i] = 0
			update[i] = sl.Header
			update[i].Level[i].Span = sl.Length
		}
		sl.Level = level
	}
	x = createNode(level, obj)

	for i := 0; i < level; i++ {
		x.Level[i].Forward = update[i].Level[i].Forward
		update[i].Level[i].Forward = x

		/* update span covered by update[i] as x is inserted here */
		x.Level[i].Span = update[i].Level[i].Span - (rank[0] - rank[i])
		update[i].Level[i].Span = (rank[0] - rank[i]) + 1
	}

	/* increment span for untouched levels */
	for i := level; i < sl.Level; i++ {
		update[i].Level[i].Span++
	}

	if update[0] == sl.Header {
		x.Backward = nil
	} else {
		x.Backward = update[0]
	}

	if x.Level[0].Forward != nil {
		x.Level[0].Forward.Backward = x
	} else {
		sl.Tail = x
	}

	sl.Length++

	return x
}

// internal function to implement the insert logic
func (sl *Skiplist) skiplistDeleteNode(x *Node, update [skiplistMaxLevel](*Node)) {
	for i := 0; i < sl.Level; i++ {
		if update[i].Level[i].Forward == x {
			update[i].Level[i].Span += x.Level[i].Span - 1
			update[i].Level[i].Forward = x.Level[i].Forward
		} else {
			update[i].Level[i].Span--
		}
	}

	if x.Level[0].Forward != nil {
		x.Level[0].Forward.Backward = x.Backward
	} else {
		sl.Tail = x.Backward
	}

	for sl.Level > 1 && sl.Header.Level[sl.Level-1].Forward == nil {
		sl.Level--
	}

	sl.Length--
}

//Delete is used to delete the node
func (sl *Skiplist) Delete(obj interface{}) int {
	var update [skiplistMaxLevel]*Node

	x := sl.Header

	for i := sl.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil &&
			sl.comp(x.Level[i].Forward.Obj, obj) < 0 {
			x = x.Level[i].Forward
		}
		update[i] = x
	}
	x = x.Level[0].Forward
	if x != nil && sl.comp(x.Obj, obj) == 0 {
		sl.skiplistDeleteNode(x, update)
		// skiplistfreenode
		return 1
	}
	return 0 // not found
}

//Find is used to find the node in skiplist
func (sl *Skiplist) Find(obj interface{}) *Node {
	var x *Node
	x = sl.Header
	for i := sl.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil && sl.comp(x.Level[i].Forward.Obj, obj) < 0 {
			x = x.Level[i].Forward
		}
	}
	x = x.Level[0].Forward
	if x != nil && sl.comp(x.Obj, obj) == 0 {
		return x
	}
	x = nil
	return x

}

//PopHead is used to pop the head node from skiplist
func (sl *Skiplist) PopHead() interface{} {
	var res interface{}
	var x = sl.Header

	x = x.Level[0].Forward
	if x == nil {
		res = nil
		return res
	}
	res = x.Obj
	sl.Delete(res)
	return res
}

//PopTail is used to pop the tail elem from skiplist
func (sl *Skiplist) PopTail() interface{} {
	var res interface{}
	var x = sl.Tail
	if x == nil {
		res = nil
		return res
	}

	res = x.Obj
	sl.Delete(res)
	return res
}

// SkiplistLength is used to show the length
func (sl *Skiplist) SkiplistLength() uint64 {
	return sl.Length
}

// Print is used to print the skiplist
func (sl *Skiplist) Print() {
	var x *Node
	for i := 0; i < sl.Level; i++ {
		x = sl.Header
		fmt.Printf("========= level %d ==========\n", i)
		for x != sl.Tail && x.Level[i].Forward != nil {
			x = x.Level[i].Forward
			fmt.Printf("--- %v --- ", x.Obj)
		}
		fmt.Println()
	}
}
