package api

type MessageType string

const (
	MessageTypeIntroduction     MessageType = "introduction"
	MessageTypeAcknowledgement  MessageType = "acknowledgement"
	MessageTypeNodeInfo         MessageType = "node_info"
	MessageTypeFileRequest      MessageType = "file_request"
	MessageTypeMasterFileHeader MessageType = "master_file_header"
	MessageTypeNodeFileHeader   MessageType = "node_file_header"
	MessageTypeError            MessageType = "error"
)

type Message struct {
	MessageType MessageType `json:"message_type,omitempty"`

	MessageID           string `json:"message_id,omitempty"`
	ResponseToMessageID string `json:"response_to_message_id,omitempty"`

	Introduction     Introduction     `json:"introduction,omitempty"`
	Acknowledgement  Acknowledgement  `json:"acknowledgement,omitempty"`
	NodeInfo         NodeInfo         `json:"node_info,omitempty"`
	FileRequest      FileRequest      `json:"file_request,omitempty"`
	MasterFileHeader MasterFileHeader `json:"master_file_header,omitempty"`
	NodeFileHeader   NodeFileHeader   `json:"node_file_header,omitempty"`
	Error            Error            `json:"error,omitempty"`
}

type Introduction struct {
	IsNew        bool   `json:"is_new,omitempty"`
	MasterSecret string `json:"master_secret,omitempty"`
	NodeID       string `json:"node_id,omitempty"`
}

type Acknowledgement struct {
	NodeID string `json:"node_id,omitempty"`
}

type StorageInfo struct {
	StorageIndex  uint64 `json:"storage_index,omitempty"`
	TotalCapacity uint64 `json:"total_capacity,omitempty"`
	FreeCapacity  uint64 `json:"free_capacity,omitempty"`
}

type NodeInfo struct {
	IsPersistent      bool          `json:"is_persistent,omitempty"`
	AcceptedChunkSize uint64        `json:"accepted_chunk_size,omitempty"`
	StorageInfo       []StorageInfo `json:"storage_info,omitempty"`
}

type FileRequest struct {
	FileID            string `json:"file_id,omitempty"`
	AcceptedChunkSize uint64 `json:"accepted_chunk_size,omitempty"`
}

type MasterFileHeader struct {
	FileID         string   `json:"file_id,omitempty"`
	Size           uint64   `json:"size,omitempty"`
	Name           string   `json:"name,omitempty"`
	ChunksNumber   uint64   `json:"chunks_number,omitempty"`
	StorageIndexes []uint64 `json:"storage_indexes,omitempty"`
}

type NodeFileHeader struct {
	FileID       string `json:"file_id,omitempty"`
	Size         uint64 `json:"size,omitempty"`
	ChunksNumber uint64 `json:"chunks_number,omitempty"`
}
