package gops

import "testing"

func TestProcInfo(t *testing.T) {
	p, err := NewProcess(1)
	if err != nil {
		t.Fatal(err)
	}
	pi, err := p.ProcInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pi)
}
