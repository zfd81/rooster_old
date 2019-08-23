package store

const (
	defaultBlockCapacity = 64 * 1024
)

type Block []Row

func (b *Block) Add(row Row) *Block {
	if row != nil {
		*b = append(*b, row)
	}
	return b
}

func (b *Block) AddBatch(rows []Row) *Block {
	if rows != nil {
		*b = append(*b, rows...)
	}
	return b
}

func (b *Block) Remove(index int) *Block {
	*b = append((*b)[:index], (*b)[index+1:]...)
	return b
}

func (b *Block) GetRow(index int) *Row {
	return &(*b)[index]
}

func (b *Block) SetRow(index int, row Row) {
	(*b)[index] = row
}

func (b *Block) GetRowField(rindex int, findex int) *Field {
	return &(*b)[rindex][findex]
}

func (b *Block) SetRowField(rindex int, findex int, field Field) {
	(*b)[rindex][findex] = field
}

func (b *Block) Length() int {
	return len(*b)
}

func (b *Block) Size() int {
	size := 0
	for _, row := range *b {
		size = size + row.Size()
	}
	return size
}

func NewBlock() *Block {
	var block Block = make([]Row, 0, defaultBlockCapacity)
	return &block
}
