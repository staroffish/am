package downloadmanager

type DownloadTaskInfo struct {
	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Regexp        string `protobuf:"bytes,3,opt,name=regexp,proto3" json:"regexp,omitempty"`
	LatestChapter int32  `protobuf:"varint,4,opt,name=latest_chapter,json=latestChapter,proto3" json:"latest_chapter,omitempty"`
	StorePath     string `protobuf:"bytes,5,opt,name=store_path,json=storePath,proto3" json:"store_path,omitempty"`
	UpdateTime    string `protobuf:"bytes,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	AnimeId       string `protobuf:"bytes,7,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`

	CreateRevision int64 `json:"-"`
	ModRevision    int64 `json:"-"`
	Version        int64 `json:"-"`
}
