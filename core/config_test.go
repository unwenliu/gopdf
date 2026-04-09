package core

import (
	"testing"
)

func TestSetPage_A4_P(t *testing.T) {
	r := CreateReport()
	r.SetPage("A4", "P")
	if r.pageWidth != 595.28 || r.pageHeight != 841.89 {
		t.Errorf("A4 P: expected 595.28x841.89, got %fx%f", r.pageWidth, r.pageHeight)
	}
}

func TestSetPage_A4_L(t *testing.T) {
	r := CreateReport()
	r.SetPage("A4", "L")
	if r.pageWidth != 841.89 || r.pageHeight != 595.28 {
		t.Errorf("A4 L: expected 841.89x595.28, got %fx%f", r.pageWidth, r.pageHeight)
	}
}

func TestSetPage_A3_P(t *testing.T) {
	r := CreateReport()
	r.SetPage("A3", "P")
	if r.pageWidth != 841.89 || r.pageHeight != 1190.55 {
		t.Errorf("A3 P: expected 841.89x1190.55, got %fx%f", r.pageWidth, r.pageHeight)
	}
}

func TestSetPage_A3_L(t *testing.T) {
	r := CreateReport()
	r.SetPage("A3", "L")
	if r.pageWidth != 1190.55 || r.pageHeight != 841.89 {
		t.Errorf("A3 L: expected 1190.55x841.89, got %fx%f", r.pageWidth, r.pageHeight)
	}
}

func TestSetPage_LTR_P(t *testing.T) {
	r := CreateReport()
	r.SetPage("LTR", "P")
	if r.pageWidth != 612 || r.pageHeight != 792 {
		t.Errorf("LTR P: expected 612x792, got %fx%f", r.pageWidth, r.pageHeight)
	}
}

func TestSetPage_LTR_L(t *testing.T) {
	r := CreateReport()
	r.SetPage("LTR", "L")
	if r.pageWidth != 792 || r.pageHeight != 612 {
		t.Errorf("LTR L: expected 792x612, got %fx%f", r.pageWidth, r.pageHeight)
	}
}

func TestSetPage_CustomSize(t *testing.T) {
	Register("CUSTOM", &Config{
		startX:        50,
		startY:        50,
		endX:          550,
		endY:          750,
		width:         600,
		height:        800,
		contentWidth:  500,
		contentHeight: 700,
	})

	r := CreateReport()
	r.SetPage("CUSTOM", "P")
	if r.pageWidth != 600 || r.pageHeight != 800 {
		t.Errorf("CUSTOM P: expected 600x800, got %fx%f", r.pageWidth, r.pageHeight)
	}

	r2 := CreateReport()
	r2.SetPage("CUSTOM", "L")
	if r2.pageWidth != 800 || r2.pageHeight != 600 {
		t.Errorf("CUSTOM L: expected 800x600, got %fx%f", r2.pageWidth, r2.pageHeight)
	}
}

func TestSetPage_UnknownPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for unknown size")
		}
	}()
	r := CreateReport()
	r.SetPage("UNKNOWN", "P")
}

func TestSetPage_InvalidOrientationPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for invalid orientation")
		}
	}()
	r := CreateReport()
	r.SetPage("A4", "X")
}

func TestSetPage_AtomicCellFormat(t *testing.T) {
	r := CreateReport()
	r.SetPage("A4", "P")
	cells := r.GetAtomicCells()

	expected := "P|pt|A4|P"
	found := false
	for _, cell := range *cells {
		if cell == expected {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected atomicCell %q not found in cells", expected)
	}
}
