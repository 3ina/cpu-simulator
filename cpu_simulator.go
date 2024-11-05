package main

import (
	"bufio"
	"fmt"
	"github.com/3ina/cpu-simulator/simulator"
	"os"
	"strconv"
	"strings"
)

func main() {
	mem := simulator.NewMemory()
	cpu := &simulator.Cpu{}

	err := LoadInstructions("instructions.txt", mem)

	if err != nil {
		fmt.Println("Error loading instructions:", err)
		return
	}

	for !cpu.Halted {
		cpu.Fetch(mem)
		cpu.DecodeAndExecute(mem)

		fmt.Printf("PC: %d, AC: %d, IR: {%s %d}\n",
			cpu.PC, cpu.AC, cpu.IR.Opcode, cpu.IR.Operand)
	}
}

func LoadInstructions(filename string, mem *simulator.Memory) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	address := 0
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "+") {
			line = line[1:]
		}

		if len(line) < 3 {
			continue
		}

		opcode := strings.TrimSpace(line[:3])
		operandStr := strings.TrimSpace(line[3:])

		operand, err := strconv.Atoi(operandStr)
		if err != nil {
			if opcode == "000" {
				operand = 0
			} else {
				return fmt.Errorf(
					"invalid operand at address %d: %v",
					address, operandStr)
			}
		}

		if opcode == "000" {
			mem.Write(address, simulator.Instruction{
				Opcode:  "DATA",
				Operand: operand,
			})
		} else {
			mem.Write(address, simulator.Instruction{
				Opcode:  opcode,
				Operand: operand,
			})
		}
		address++
	}

	return scanner.Err()
}
