package bus

type Dma struct {
	active     bool
	byte       byte
	value      byte
	startDelay byte
}

func initDma() *Dma {
	return &Dma{}
}

func (d *Dma) start(start byte) {
	d.active = true
	d.byte = 0
	d.startDelay = 2
	d.value = start
}

func (d *Dma) isTransferring() bool {
	return d.active
}
