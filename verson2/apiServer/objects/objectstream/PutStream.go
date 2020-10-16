package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream struct{
	writer *io.PipeWriter
	c chan error
}

func NewPutStream(server, object string) *PutStream {
	// 这一对管道是希望能以写入数据流的方式操作HTTP的PUT请求
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		request, _ := http.NewRequest("PUT", "http://" +
			"/objects/" + object, reader)
		client := http.Client{}
		r, e := client.Do(request)
		if e == nil && r.StatusCode != http.StatusOK{
			e = fmt.Errorf("dataserver return http code %d", r.StatusCode)
		}
		c <- e
	}()
	return &PutStream{writer, c}
}

func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *PutStream) Close() error{
	w.writer.Close()
	return <- w.c
}














































