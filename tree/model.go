package tree

import "fmt"

type NodeItem struct {
	Data  []byte	`json:"data"`
	Left  *NodeItem `json:"left"`
	Right *NodeItem	`json:"right"`
}

func InitNode(data []byte) *NodeItem {
	node := new(NodeItem)
	node.Data = data
	return node
}

func (ni *NodeItem) AddData(data []byte) error {
	if ni.insertDataIsBigger(data) {
		if ni.Right == nil {
			ni.Right = InitNode(data)
		} else {
			ni.Right.AddData(data)
		}
	} else {
		if ni.Left == nil {
			ni.Left = InitNode(data)
		} else {
			ni.Left.AddData(data)
		}
	}

	return nil
}

func (ni *NodeItem) insertDataIsBigger(data []byte) bool {
	var minLen int
	if len(ni.Data) > len(data) {
		minLen = len(data)
	} else {
		minLen = len(ni.Data)
	}

	for ind := 0; ind < minLen; ind++ {
		if ni.Data[ind] < data[ind] {
			return true
		}
	}
	return false
}

func (ni *NodeItem) String() string {
	return fmt.Sprintf("{data: %q, left: %v right: %v}", ni.Data, ni.Left, ni.Right)
}
