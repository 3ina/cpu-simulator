package simulator

type Memory struct {
	data map[int]Instruction
}

func NewMemory() *Memory {
	return &Memory{data: make(map[int]Instruction)}
}
