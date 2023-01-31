package cart

import (
	"github.com/mikeletux/goboy/pkg/test"
	"reflect"
	"strings"
	"testing"
)

const (
	testRomPath          string = "/home/mikeletux/development/goboy/roms/tetris.gb"
	tetrisTitle          string = "TETRIS"
	licenseeCodeReadable string = "Nintendo"
)

func TestCartridge(t *testing.T) {
	cartridge, err := NewCartridge(testRomPath)
	if err != nil {
		t.Errorf("error while initiating cartridge - %s", err)
	}

	t.Parallel()

	t.Run("Test Nintendo Logo", func(t *testing.T) {
		got := test.AssessArrays(cartridge.CartridgeHeader.NintendoLogo[:], test.NintendoCartridgeLogo)
		if !got {
			t.Errorf("error while testing Nintendo logo in cartridge. Got %v expected %v",
				cartridge.CartridgeHeader.NintendoLogo, test.NintendoCartridgeLogo)
		}
	})

	t.Run("Test Tetris Title", func(t *testing.T) {
		got := strings.Contains(cartridge.CartridgeHeader.GetReadableTitle(), tetrisTitle)
		if !got {
			t.Errorf("error while testing Tetris title in cartridge. Got %s expected %s",
				cartridge.CartridgeHeader.GetReadableTitle(), tetrisTitle)
		}
	})

	t.Run("Test New Licensee code", func(t *testing.T) {
		got := reflect.DeepEqual(cartridge.CartridgeHeader.GetReadableLicenseeCode(), licenseeCodeReadable)
		if !got {
			t.Errorf("error while testing readable licensee code. Got %s expected %s",
				cartridge.CartridgeHeader.GetReadableLicenseeCode(), licenseeCodeReadable)
		}
	})

}
