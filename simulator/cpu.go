package simulator

type Instruction struct {
	opcode  string
	operand int
}

type Cpu struct {
	PC     int         // program counter
	AC     int         // Accumulator
	IR     Instruction // instruction register
	Halted bool        // cpu halt status
}

func (cpu *Cpu) Fetch(mem *Memory) {
	cpu.IR = mem.Read(cpu.PC)
	cpu.PC++
}
