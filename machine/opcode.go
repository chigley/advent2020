package machine

type Opcode int

const (
	OpAcc Opcode = iota
	OpJmp
	OpNop
)

var ops = map[string]Opcode{
	"acc": OpAcc,
	"jmp": OpJmp,
	"nop": OpNop,
}
