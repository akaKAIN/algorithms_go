package copy

type LinkedItem struct {
	Data     []byte
	NextItem *LinkedItem
}

func (l *LinkedItem) GetCopy() LinkedItem {
	return *l
}

func InitLinkedItem(data []byte, link *LinkedItem) *LinkedItem {
	li := new(LinkedItem)
	li.Data = data
	li.NextItem = link
	return li
}
