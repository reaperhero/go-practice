package lux

import (
	"github.com/iawia002/lux/downloader"
	"github.com/iawia002/lux/extractors"
	"testing"
)

func TestDownloadBiliBili(t *testing.T) {
	down := downloader.New(downloader.Options{
		//Silent:         c.Bool("silent"),
		//InfoOnly:       c.Bool("info"),
		//Stream:         c.String("stream-format"),
		//Refer:          c.String("refer"),
		//OutputPath:     c.String("output-path"),
		//OutputName:     c.String("output-name"),
		//FileNameLength: int(c.Uint("file-name-length")),
		//Caption:        c.Bool("caption"),
		//MultiThread:    c.Bool("multi-thread"),
		//ThreadNumber:   int(c.Uint("thread")),
		//RetryTimes:     int(c.Uint("retry")),
		//ChunkSizeMB:    int(c.Uint("chunk-size")),
		//UseAria2RPC:    c.Bool("aria2"),
		//Aria2Token:     c.String("aria2-token"),
		//Aria2Method:    c.String("aria2-method"),
		//Aria2Addr:      c.String("aria2-addr"),
	})
	testCase := struct {
		name string
		data *extractors.Data
	}{
		name: "normal test",
		data: &extractors.Data{
			Site:  "douyin",
			Title: "test",
			Type:  extractors.DataTypeVideo,
			URL:   "https://www.douyin.com",
			Streams: map[string]*extractors.Stream{
				"default": {
					ID: "default",
					Parts: []*extractors.Part{
						{
							URL:  "https://aweme.snssdk.com/aweme/v1/playwm/?video_id=v0200f9a0000bc117isuatl67cees890&line=0",
							Size: 4927877,
							Ext:  "mp4",
						},
					},
				},
			},
		},
	}
	err := down.Download(testCase.data)
	if err != nil {
		t.Error(err)
	}

}
