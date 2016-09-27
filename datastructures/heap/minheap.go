package heap

//import "fmt"

type Entry struct {
	W  int
	Id int
}

type MinHeap interface {
	Min() Entry
	Remove(id int)
	Insert(Entry)
}

type ArrMinHeap []Entry

func NewMinHeap() MinHeap {
	minHeap := ArrMinHeap(make([]Entry, 1))
	minHeap[0] = Entry{Id: -1}
	return &minHeap
}

func (minh *ArrMinHeap) Min() Entry {
	min := (*minh)[1]

	return min
}

func (minh *ArrMinHeap) sinkDown(index int) {

}

func (minh *ArrMinHeap) flyUp(index int) {
	for curIndex := index; curIndex > 1; curIndex = int(curIndex / 2) {
		println(curIndex)
		if (*minh)[curIndex].W < (*minh)[curIndex/2].W {
			minh.swap(curIndex, curIndex/2)
		} else {
			return
		}
	}

	return
}

func (minh *ArrMinHeap) swap(index1 int, index2 int) {
	temp := (*minh)[index1]
	(*minh)[index1] = (*minh)[index2]
	(*minh)[index2] = temp
}

func (minh *ArrMinHeap) Insert(entry Entry) {
	*minh = append(*minh, entry)
	minh.flyUp(len(*minh) - 1)
}

func (minh *ArrMinHeap) Remove(id int) {
	//findI
}
