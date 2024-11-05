package simulator

import "fmt"

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

func (cpu *Cpu) DecodeAndExecute(mem *Memory) {
	switch cpu.IR.opcode {
	case "LDA":
		cpu.AC = mem.Read(cpu.IR.operand).operand
	case "STO":
		mem.Write(cpu.IR.operand, Instruction{"DATA", cpu.AC})
	case "ADD":
		cpu.AC += mem.Read(cpu.IR.operand).operand
	case "SUB":
		cpu.AC -= mem.Read(cpu.IR.operand).operand
	case "MPY":
		cpu.AC *= mem.Read(cpu.IR.operand).operand
	case "BRU":
		cpu.PC = cpu.IR.operand
	case "BRM":
		if cpu.AC < 0 {
			cpu.PC = cpu.IR.operand
		}
	case "WWD":
		value := mem.Read(cpu.IR.operand).operand
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

		mem.Write(cpu.IR.operand, Instruction{"DATA", input})
	case "HLT":
		cpu.Halted = true
	case "NOP":
		// do nothing
	default:
		fmt.Println("Unknown instruction: ", cpu.IR.opcode)
		cpu.Halted = true

	}
}
