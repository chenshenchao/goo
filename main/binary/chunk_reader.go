package binary

import (
	"encoding/binary"
	"math"
)

type ChunkReader struct {
	data []byte
}

func (i *ChunkReader) ReadByte() byte {
	result := i.data[0]
	i.data = i.data[1:]
	return result
}

func (i *ChunkReader) ReadUint32() uint32 {
	result := binary.LittleEndian.Uint32(i.data)
	i.data = i.data[4:]
	return result
}

func (i *ChunkReader) ReadUint64() uint64 {
	result := binary.LittleEndian.Uint64(i.data)
	i.data = i.data[8:]
	return result
}

func (i *ChunkReader) ReadFloat64() float64 {
	return math.Float64frombits(i.ReadUint64())
}

func (i *ChunkReader) ReadBytes(size uint64) []byte {
	end := size - 1
	result := i.data[:end]
	i.data = i.data[end:]
	return result
}

func (i *ChunkReader) ReadString() string {
	size := uint64(i.ReadByte())
	if 0 == size {
		return ""
	}
	if 0xFF == size {
		size = uint64(i.ReadUint64())
	}
	bytes := i.ReadBytes(size)
	return string(bytes)
}

