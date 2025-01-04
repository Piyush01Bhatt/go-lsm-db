package skiplist

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
)

const MAXLEVEL = 16

type Node struct {
	key   string
	value string
	next  []*Node
}

type Skiplist struct {
	head  *Node
	level int
}

func NewNode(key string, value string, level int) *Node {
	return &Node{
		key:   key,
		value: value,
		next:  make([]*Node, level),
	}
}

func NewSkiplist() *Skiplist {
	return &Skiplist{
		head:  NewNode("", "", MAXLEVEL),
		level: 1,
	}
}

// randomLevel generates a random level for a new node
func randomLevel() int {
	level := 1
	for rand.Float32() < 0.5 && level < MAXLEVEL {
		level++
	}
	return level
}

func (sl *Skiplist) search(key string) (*Node, [MAXLEVEL]*Node) {
	var journey [MAXLEVEL]*Node
	start := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for start.next[i] != nil && start.next[i].key < key {
			start = start.next[i]
		}
		journey[i] = start
	}

	if start != nil && start.next[0] != nil && start.next[0].key == key {
		return start.next[0], journey
	}
	return nil, journey
}

func (sl *Skiplist) Find(key string) (string, error) {
	node, _ := sl.search(key)
	if node == nil {
		return "", errors.New("key not found")
	}
	return node.value, nil
}

func (sl *Skiplist) Insert(key, value string) {
	found, journey := sl.search(key)
	if found != nil {
		found.value = value
		return
	}
	level := randomLevel()
	found = NewNode(key, value, level)
	for i := level - 1; i >= 0; i-- {
		if journey[i] == nil {
			journey[i] = sl.head
		}
		found.next[i] = journey[i].next[i]
		journey[i].next[i] = found
	}
	if sl.level < level {
		sl.level = level
	}
}

func (sl *Skiplist) Print() {
	for i := sl.level - 1; i >= 0; i-- {
		start := sl.head
		var data []string
		for start != nil {
			if start.key == "" {
				data = append(data, "HEAD")
			} else {
				data = append(data, start.key)
			}
			start = start.next[i]
		}
		data = append(data, "NIL")
		result := strings.Join(data, "->")
		fmt.Println(result)
	}
}
