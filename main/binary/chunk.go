package binary

type Chunk struct {
	ChunkHeader
	entry *Prototype 
}
