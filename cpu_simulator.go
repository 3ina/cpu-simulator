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
