package copy

type LinkedItem struct {
	Data     []byte
	NextItem *LinkedItem
}

func InitLinkedItem(data []byte, link *LinkedItem) *LinkedItem {
	li := new(LinkedItem)
	li.Data = data
	li.NextItem = link
	return li
}
