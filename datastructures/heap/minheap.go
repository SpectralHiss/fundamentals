package heap

type Entry struct {
	W  int
	Id int
}

type MinHeap interface {
	Min() Entry
	Remove(id int)
	Insert(Entry)
	ExtractMin() Entry
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

func (minh *ArrMinHeap) Insert(entry Entry) {
	*minh = append(*minh, entry)
	minh.flyUp(len(*minh) - 1)
}

// iterative style here..
func (minh *ArrMinHeap) flyUp(index int) {
	heap := (*minh)

	for curIndex := index; curIndex > 1; curIndex = int(curIndex / 2) {
		if heap[curIndex].W < heap[curIndex/2].W {
			minh.swap(curIndex, curIndex/2)
		} else {
			return
		}
	}

	return
}

func (minh *ArrMinHeap) Remove(id int) {
	heap := (*minh)
	pos := minh.scan(id)
	if pos != -1 {
		lastPos := len(*minh) - 1
		heap[pos] = heap[lastPos]
		*minh = heap[:lastPos]
		minh.sinkDown(pos)
	}
}

// for variety we do this recursively
func (minh *ArrMinHeap) sinkDown(index int) {
	heap := (*minh)

	if index > len(heap)/2 {
		return
	} else {
		leftPos := 2 * index
		left := heap[leftPos]

		rightPos := leftPos + 1

		if rightPos > len(heap)-1 {
			// no right elem at bottom, tree is "complete"
			if left.W < heap[index].W {
				minh.swap(index, leftPos)
			}
			return
		} else {
			right := heap[rightPos]

			if left.W < heap[index].W || right.W < heap[index].W {

				var swapPosition int
				if left.W < right.W {
					swapPosition = leftPos
				} else {
					swapPosition = rightPos
				}

				minh.swap(index, swapPosition)
				minh.sinkDown(swapPosition)
			}
		}
	}
}

func (minh *ArrMinHeap) ExtractMin() Entry {
	ret := minh.Min()
	minh.Remove((*minh)[1].Id)
	return ret
}

func (minh *ArrMinHeap) swap(index1 int, index2 int) {
	heap := (*minh)

	temp := heap[index1]
	heap[index1] = heap[index2]
	heap[index2] = temp
}

func (minh *ArrMinHeap) scan(id int) int {
	heap := (*minh)

	for i := 1; i < len(heap); i++ {
		if heap[i].Id == id {
			return i
		}
	}
	return -1
}
