package simulator

type Memory struct {
	data map[int]Instruction
}

func NewMemory() *Memory {
	return &Memory{data: make(map[int]Instruction)}
}

func (mem *Memory) Read(address int) Instruction {
	if instr, exists := mem.data[address]; exists {
		return instr
	}
	// Return a default NOP (No Operation)
	return Instruction{"NOP", 0}
}

func (mem *Memory) Write(address int, instruction Instruction) {
	mem.data[address] = instruction
}
