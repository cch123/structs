package skiplist

import "math/rand"
import "fmt"

const SKIPLIST_MAXLEVEL = 32
const SKIPLIST_P = 0.25

type SkiplistLevel struct {
	Forward *SkiplistNode
	Span    uint64
}

type SkiplistNode struct {
	Obj      interface{}
	Backward *SkiplistNode
	Level    [SKIPLIST_MAXLEVEL]SkiplistLevel
}

type Compare func(a interface{}, b interface{}) int

type Skiplist struct {
	Header *SkiplistNode
	Tail   *SkiplistNode
	// compare function
	comp   Compare
	Length uint64
	Level  int
}

func skiplistCreateNode(level int, obj interface{}) *SkiplistNode {
	var zn = new(SkiplistNode)
	zn.Obj = obj
	return zn
}

func SkiplistCreate(p Compare) *Skiplist {
	sl := new(Skiplist)
	sl.Level = 1
	sl.Length = 0
	sl.Header = skiplistCreateNode(SKIPLIST_MAXLEVEL, nil)
	for j := 0; j < SKIPLIST_MAXLEVEL; j++ {
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
	var level int = 1
	for i, j := rand.Int()&0xFFFF, 0.25*0xFFFF; i < int(j); i = rand.Int() & 0xFFFF {
		level += 1
	}

	if level > SKIPLIST_MAXLEVEL {
		return SKIPLIST_MAXLEVEL
	}

	return level
}

func (sl *Skiplist) SkiplistInsert(obj interface{}) *SkiplistNode {

	// record the nodes which forward pointer need to be updated
	var update [SKIPLIST_MAXLEVEL]*SkiplistNode
	var rank [SKIPLIST_MAXLEVEL]uint64

	var x *SkiplistNode

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
	x = skiplistCreateNode(level, obj)

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
func (sl *Skiplist) skiplistDeleteNode(x *SkiplistNode, update [SKIPLIST_MAXLEVEL](*SkiplistNode)) {
	for i := 0; i < sl.Level; i++ {
		if update[i].Level[i].Forward == x {
			update[i].Level[i].Span += x.Level[i].Span - 1
			update[i].Level[i].Forward = x.Level[i].Forward
		} else {
			update[i].Level[i].Span -= 1
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

func (sl *Skiplist) SkiplistDelete(obj interface{}) int {
	var update [SKIPLIST_MAXLEVEL]*SkiplistNode

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

func (sl *Skiplist) SkiplistFind(obj interface{}) *SkiplistNode {
	var x *SkiplistNode
	x = sl.Header
	for i := sl.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil && sl.comp(x.Level[i].Forward.Obj, obj) < 0 {
			x = x.Level[i].Forward
		}
	}
	x = x.Level[0].Forward
	if x != nil && sl.comp(x.Obj, obj) == 0 {
		return x
	} else {
		x = nil
		return x
	}

}

// pop, and return the value
func (sl *Skiplist) SkiplistPopHead() interface{} {
	var res interface{}
	var x *SkiplistNode = sl.Header

	x = x.Level[0].Forward
	if x != nil {
		res = nil
		return res
	}
	res = x.Obj
	sl.SkiplistDelete(res)
	return res
}

func (sl *Skiplist) SkiplistPopTail() interface{} {
	var res interface{}
	var x *SkiplistNode = sl.Tail
	if x != nil {
		res = nil
		return res
	}

	res = x.Obj
	sl.SkiplistDelete(res)
	return res
}

func (sl *Skiplist) SkiplistLength() uint64 {
	return sl.Length
}

func (sl *Skiplist) Print() {
	var x *SkiplistNode
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
