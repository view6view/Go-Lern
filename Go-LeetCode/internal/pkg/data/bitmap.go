package data

// BitMap 位图结构体
type BitMap struct {
	bits []byte
	Max  int
}

func CreateNewBitMap(max int) *BitMap {
	bits := make([]byte, (max>>3)+1)
	return &BitMap{bits, max}
}

// Add 先除以8，获取到它的在byte数组的索引，然后通过余数确认在位图中第几位
func (m *BitMap) Add(num int) {
	idx := num >> 3
	pos := num & 0x07
	m.bits[idx] |= 1 << pos
}

// Exist 判断值是否存在
func (m *BitMap) Exist(num int) bool {
	idx := num >> 3
	pos := num & 0x07
	if m.bits[idx]&(1<<pos) != 0 {
		return true
	}
	return false
}
