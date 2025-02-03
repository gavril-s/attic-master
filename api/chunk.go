package api

type Chunk struct {
	FileID     string `json:"file_id,omitempty"`
	ChunkIndex uint64 `json:"chunk_index,omitempty"`
	Payload    []byte `json:"payload,omitempty"`
}
