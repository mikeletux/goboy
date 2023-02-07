package cpu

import "fmt"

func (c *CPU) fetchData() error {
	c.MemoryDestination = 0
	c.DestinationIsMemory = false

	switch c.CurrentInstruction.AddressingMode {
	case amImp:
		return nil

	case amR:
		fetchedData, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
		if err != nil {
			return err
		}
		c.FetchedData = fetchedData
		return nil

	case amRnD8:
		c.FetchedData = uint16(c.bus.BusRead(c.registers.GetPCAndIncrement()))
		c.emulateCpuCycles(1)
		return nil

	case amD16:
		var low = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		var high = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		c.FetchedData = uint16(low) | uint16(high)<<8
		return nil

	// To be done still
	default:
		return fmt.Errorf("addressing mode %d doesn't exist", c.CurrentInstruction.AddressingMode)
	}

	return nil // This return should not be reached ever
}
