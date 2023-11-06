package xface

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"

	"github.com/Kagami/go-face"
)

// ================================================================
//
// ================================================================
const DimensionCount = 128

func SquaredDist(f1, f2 face.Descriptor) float64 {
	sum, diff := float64(0), float64(0)
	for i := 0; i < DimensionCount; i += 1 {
		diff = float64(f1[i] - f2[i])
		sum += diff * diff
	}

	return sum
}

// ================================================================
//
// ================================================================
type Descriptor face.Descriptor

func (d Descriptor) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, d); err != nil {
		return nil, err
	} else {
		return buf.Bytes(), nil
	}
}

func (d *Descriptor) Scan(src any) error {
	if src != nil {
		buf := bytes.NewReader(src.([]byte))
		return binary.Read(buf, binary.LittleEndian, d)
	}

	return nil
}

func (d Descriptor) DistWithFace(f Descriptor) float64 {
	sum, diff := float64(0), float64(0)
	for i := 0; i < DimensionCount; i += 1 {
		diff = float64(d[i] - f[i])
		sum += diff * diff
	}

	return sum
}

// ================================================================
//
// ================================================================
const FaceDistThreshold = 0.15

type Threshold float64

func (t *Threshold) Validate() {
	if *t == 0.0 {
		*t = FaceDistThreshold
	} else if *t < 0.0 {
		*t = 0.01
	} else if *t > 0.99 {
		*t = 0.99
	}
}
