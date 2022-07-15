package v2

import "errors"

type BitMap struct {
	vec  []byte
	size int
}

func New(size int) *BitMap {
	return &BitMap{
		size: size,
		vec:  make([]byte, size),
	}
}

func (bm *BitMap) Set(num int) (ok bool, err error) {
	if num/8 >= bm.size {
		return false, errors.New("the num overflows the size of bitmap")
	}
	bm.vec[num/8] |= 1 << (num % 8)
	return true, nil
}

func (bm *BitMap) Exist(num int) bool {
	if num/8 >= bm.size {
		return false
	}
	return bm.vec[num/8]&(1<<(num%8)) > 0
}

func (bm *BitMap) Sort() (ret []int) {
	ret = make([]int, 0)
	for i := 0; i < (8 * bm.size); i++ {
		if bm.Exist(i) {
			ret = append(ret, i)
		}
	}
	return
}
