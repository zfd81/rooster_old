package store

const (
	defaultBlockCapacity = 64 * 1024
)

type Block struct {
	Name  string
	lines []Line
}

func (b *Block) Add(line Line) *Block {
	if line != nil {
		b.lines = append(b.lines, line)
	}
	return b
}

func (b *Block) AddBatch(lines []Line) *Block {
	if lines != nil {
		b.lines = append(b.lines, lines...)
	}
	return b
}

func (b *Block) Remove(index int) *Block {
	b.lines = append(b.lines[:index], b.lines[index+1:]...)
	return b
}

func (b *Block) GetLine(index int) *Line {
	return &b.lines[index]
}

func (b *Block) SetLine(index int, line Line) {
	b.lines[index] = line
}

func (b *Block) GetField(lindex int, findex int) *Field {
	return &b.lines[lindex][findex]
}

func (b *Block) SetField(lindex int, findex int, field Field) {
	b.lines[lindex][findex] = field
}

func (b *Block) Length() int {
	return len(b.lines)
}

func (b *Block) Size() int {
	size := 0
	for _, line := range b.lines {
		size = size + line.Size()
	}
	return size
}

func NewBlock(name string) *Block {
	return &Block{name, make([]Line, 0, defaultBlockCapacity)}
}
