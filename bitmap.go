package virtuaposp

type BitMap struct {
	m [8]byte
}

func (this *BitMap) SetBit(n int) error {
	return nil
}

func (this *BitMap) RemoveBit(n int) error {
	return nil
}

func (this *BitMap) Clear() {
	for i := 0; i < 8; i++ {
		this.m[i] = 0
	}
}
