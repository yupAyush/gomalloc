package allocator

import "encoding/binary"

const (
	headerSize = 8
	align      = 8
	minBlock   = headerSize + align
)

func roundUp(n int) int {
	return (n + align - 1) &^ (align - 1)
}

func blockSize(n int) int {
	need := roundUp(n + headerSize)
	if need < minBlock {
		need = minBlock
	}
	return need

}

func (h *Heap) header(off int) (size int, used bool) {
	v := binary.LittleEndian.Uint64(h.mem[off:]) // takes 8 bytes then combines them into single number
	return int(v >> 1), v&1 == 1

}

func (h *Heap) setHeader(off, size int, used bool) {
	v := uint64(size) << 1
	if used {
		v |= 1
	}
	binary.LittleEndian.PutUint64(h.mem[off:], v)

}
