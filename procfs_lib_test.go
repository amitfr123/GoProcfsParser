package procfsgo

import (
	"testing"
)

func TestPrintMeminfo(t *testing.T) {
	GetProcReader().PrintMeminfo()
}

func TestPrintNetDev(t *testing.T) {
	GetProcReader().PrintNetDev()
}

func TestPrintArp(t *testing.T) {
	GetProcReader().PrintArp()
}
