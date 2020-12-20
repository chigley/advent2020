package day20

type Noncer struct {
	nonce int
}

func NewNoncer() *Noncer {
	return &Noncer{nonce: 0}
}

func (n *Noncer) Nonce() int {
	n.nonce++
	return n.nonce - 1
}
