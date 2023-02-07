package cpu

import "fmt"

func (c *CPU) fetchData() error {
	c.MemoryDestination = 0
	c.DestinationIsMemory = false

	switch c.CurrentInstruction.AddressingMode {
	case amImp:
		return nil

	case amRnR:
		fetchedData, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}
		c.FetchedData = fetchedData
		return nil

	case amMRnR:
		fetchedData, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}
		c.FetchedData = fetchedData

		fetchedData, err = c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
		if err != nil {
			return err
		}
		c.MemoryDestination = fetchedData

		c.DestinationIsMemory = true

		if c.CurrentInstruction.RegisterType1 == rtC { // I don't understand this quite well
			c.MemoryDestination |= 0xFF00
		}

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

	case amRnMR:
		memoryAddress, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}

		if c.CurrentInstruction.RegisterType1 == rtC {
			memoryAddress |= 0xFF00
		}

		c.FetchedData = uint16(c.bus.BusRead(memoryAddress))
		c.emulateCpuCycles(1)

		return nil

	case amRnHLI:
		memoryAddress, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}

		c.FetchedData = uint16(c.bus.BusRead(memoryAddress))
		c.emulateCpuCycles(1)
		c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType2, memoryAddress+1)

		return nil

	case amRnHLD:
		memoryAddress, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}

		c.FetchedData = uint16(c.bus.BusRead(memoryAddress))
		c.emulateCpuCycles(1)
		c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType2, memoryAddress-1)

		return nil

	case amHLInR:
		fetchedData, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}

		c.FetchedData = fetchedData

		memoryDestination, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
		if err != nil {
			return err
		}
		c.DestinationIsMemory = true
		c.MemoryDestination = memoryDestination

		err = c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, memoryDestination+1)
		if err != nil {
			return err
		}

		return nil

	case amHLDnR:
		fetchedData, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}

		c.FetchedData = fetchedData

		memoryDestination, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
		if err != nil {
			return err
		}
		c.DestinationIsMemory = true
		c.MemoryDestination = memoryDestination

		err = c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, memoryDestination-1)
		if err != nil {
			return err
		}

		return nil

	case amRnA8:
		var memoryAddress = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)
		c.FetchedData = uint16(c.bus.BusRead(uint16(memoryAddress)))

		return nil

	case amA8nR:
		fetchedData, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}

		c.FetchedData = fetchedData

		var memoryAddress = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.DestinationIsMemory = true
		c.emulateCpuCycles(1)
		c.MemoryDestination = uint16(memoryAddress) | 0xFF00

		return nil

	case amHLnSPR, amD8:
		c.FetchedData = uint16(c.bus.BusRead(c.registers.GetPCAndIncrement()))
		c.emulateCpuCycles(1)
		return nil

	case amD16, amRnD16:
		var low = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		var high = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		c.FetchedData = uint16(low) | uint16(high)<<8
		return nil

	case amD16nR, amA16nR:
		var low = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		var high = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		c.DestinationIsMemory = true
		c.MemoryDestination = uint16(low) | uint16(high)<<8

		fetchedData, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			return err
		}

		c.FetchedData = fetchedData
		return nil

	case amMRnD8:
		c.FetchedData = uint16(c.bus.BusRead(c.registers.GetPCAndIncrement()))
		c.emulateCpuCycles(1)

		memoryDestination, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
		if err != nil {
			return err
		}

		c.DestinationIsMemory = true
		c.MemoryDestination = memoryDestination

		return nil

	case amMR:
		memoryDestination, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
		if err != nil {
			return err
		}

		c.DestinationIsMemory = true
		c.MemoryDestination = memoryDestination

		c.FetchedData = uint16(c.bus.BusRead(memoryDestination))

		return nil

	case amRnA16:
		var low = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		var high = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		memoryDestination := uint16(low) | uint16(high)<<8
		c.FetchedData = uint16(c.bus.BusRead(memoryDestination))
		c.emulateCpuCycles(1)

		return nil
		
	default:
		return fmt.Errorf("addressing mode %d doesn't exist", c.CurrentInstruction.AddressingMode)
	}

	return nil // This return should not be reached ever
}
