package model

type TorrentInfo struct {
	Hash        string  `json:"hash"`
	Name        string  `json:"name"`
	Size        int64   `json:"size"`
	Progress    float32 `json:"progress"`
	Status      string  `json:"status"`
	StorePath   string  `json:"store_path"`
	CreatedTime string  `json:"created_time"`
}
