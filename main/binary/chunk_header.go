package binary

const (
	GOO_SIGNATURE = "Goo\x00"
	GOO_VERSION = 0x10
)

type ChunkHeader struct {
	signature [4]byte
	version byte
}
