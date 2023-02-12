package deque

// Dequeue implementation
type Deque struct {
	items []int
}

func (d *Deque) Push(data int) {

	d.items = append(d.items, data)
}

func (d *Deque) PopLeft() int {

	item := d.items[0]

	d.items = d.items[1:]

	return item
}

func (d *Deque) PeekLeft() int {

	return d.items[0]
}

func (d *Deque) PopRight() int {

	item := d.items[d.Size()-1]

	d.items = d.items[:d.Size()-1]

	return item
}

func (d *Deque) PeekRight() int {

	return d.items[d.Size()-1]
}

func (d *Deque) Size() int {

	return len(d.items)
}
