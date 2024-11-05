package simulator

import "fmt"

type Instruction struct {
	Opcode  string
	Operand int
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

func (cpu *Cpu) DecodeAndExecute(mem *Memory) {
	switch cpu.IR.Opcode {
	case "LDA":
		cpu.AC = mem.Read(cpu.IR.Operand).Operand
	case "STO":
		mem.Write(cpu.IR.Operand, Instruction{"DATA", cpu.AC})
	case "ADD":
		cpu.AC += mem.Read(cpu.IR.Operand).Operand
	case "SUB":
		cpu.AC -= mem.Read(cpu.IR.Operand).Operand
	case "MPY":
		cpu.AC *= mem.Read(cpu.IR.Operand).Operand
	case "BRU":
		cpu.PC = cpu.IR.Operand
	case "BRM":
		if cpu.AC < 0 {
			cpu.PC = cpu.IR.Operand
		}
	case "WWD":
		value := mem.Read(cpu.IR.Operand).Operand
		fmt.Println("OUTPUT:", value)
	case "RWD":
		var input int
		fmt.Print("Enter input value: ")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Invalid input, Please enter an integer.")
			cpu.Halted = true
			return
		}

		mem.Write(cpu.IR.Operand, Instruction{"DATA", input})
	case "HLT":
		cpu.Halted = true
	case "NOP":
		// do nothing
	default:
		fmt.Println("Unknown instruction: ", cpu.IR.Opcode)
		cpu.Halted = true

	}
}
