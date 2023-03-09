# 🎮 GoBoy 🎮
Experimental Game Boy emulator written in Go.

## Purpose
The main purpose of this project is to understand better how the original Game Boy
worked by implementing a minimal viable emulator.

## References for this project
- [gbdev.io Pan Docs](https://gbdev.io/pandocs/)
- [lowleveldevel1712 YouTube channel](https://www.youtube.com/@lowleveldevel1712)
- [Gameboy CPU (LR35902) instruction set](https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html)
- [Gameboy Doctor](https://github.com/robert/gameboy-doctor) 
- [retrio/gb-test-roms](https://github.com/retrio/gb-test-roms)
- ...

## Roadmap
- [x] Implement CPU emulation and testing. 
- [ ] Implement PPU emulation and testing.

## Testing
Regarding testing, some components have unit tests written. It is a very small percentage tbh.
It should cover much more, but I'm just a dev, PRs are welcome :D.  
  
Regarding CPU testing, I've used [retrio/gb-test-roms](https://github.com/retrio/gb-test-roms) testing 
ROM `cpu_instrs`. Probably I should use more, probably as the development continues I'll use some more!  
  
/Miguel Sama 2023