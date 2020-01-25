package md

type ListInfo struct {
	CurrentNumber int
	IsOrdered     bool
	Marker        string
}

type ListState struct {
	ListPath []*ListInfo
}

func NewListState() *ListState {
	return &ListState{}
}

func (ls *ListState) CurrentList() *ListInfo {
	index := len(ls.ListPath) - 1
	if index < 0 {
		return nil
	}

	return ls.ListPath[index]
}

func (ls *ListState) ListLevel() int {
	return len(ls.ListPath)
}

func (ls *ListState) StartBulletList(marker string) {
	newList := &ListInfo{Marker: marker}
	ls.ListPath = append(ls.ListPath, newList)
}

func (ls *ListState) StartOrderedList(startingNumber int) {
	newList := &ListInfo{CurrentNumber: startingNumber, IsOrdered: true}
	ls.ListPath = append(ls.ListPath, newList)
}


func (ls *ListState) CompleteList() {
	index := len(ls.ListPath) - 1
	if index >= 0 {
		ls.ListPath = ls.ListPath[:index]
	}
}