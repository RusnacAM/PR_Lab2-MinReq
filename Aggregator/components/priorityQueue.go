package components

type Queue struct {
	Elements []Order
}

func (ol *Queue) Enqueue(order Order) {
	ol.Elements = append(ol.Elements, order)
}

func (ol *Queue) Dequeue() *Order {
	if ol.IsEmpty() {
		return nil
	}
	order := ol.Elements[0]
	if ol.GetSize() == 1 {
		ol.Elements = nil
		return &order
	}
	ol.Elements = ol.Elements[1:]
	return &order
}

func (ol *Queue) GetSize() int {
	return len(ol.Elements)
}

func (ol *Queue) IsEmpty() bool {
	return len(ol.Elements) == 0
}
