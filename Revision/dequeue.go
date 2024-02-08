package revision

type Deque struct {
	items []int
}

func (d *Deque) PeekLeft() int {

	if d.Size() > 0 {
		return d.items[0]
	}

	return -99
}

func (d *Deque) PopLeft() int {

	item := d.items[d.Size()-1]

	d.items = d.items[1:]
	return item
}

func (d *Deque) PeekRight() int {

	if d.Size() > 0 {
		return d.items[d.Size()-1]
	}

	return -99
}

func (d *Deque) PopRight() int {

	item := d.items[d.Size()-1]

	d.items = d.items[:d.Size()-1]

	return item
}

func (d *Deque) Push(data int) {

	d.items = append(d.items, data)
}

func (d *Deque) Size() int {

	return len(d.items)
}
