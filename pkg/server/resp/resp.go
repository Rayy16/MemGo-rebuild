package resp

import (
	respintf "memgo/pkg/server/resp/types"
	"strconv"
)

const (
	CRLF = "\r\n"
)

type SimpleStringReply struct {
	str string
}

func (r *SimpleStringReply) ToBytes() []byte {
	return []byte("+" + r.str + CRLF)
}

type ErrorReply struct {
	err error
}

func (r *ErrorReply) ToBytes() []byte {
	return []byte("-" + r.err.Error() + CRLF)
}

type IntegerReply struct {
	data int
}

func (r *IntegerReply) ToBytes() []byte {
	return []byte(":" + strconv.Itoa(r.data) + CRLF)
}

type BulkReply struct {
	data []byte
}

func (r *BulkReply) ToBytes() []byte {
	if r.data == nil {
		return []byte("$-1" + CRLF)
	}
	return []byte("$" + strconv.Itoa(len(r.data)) + CRLF + r.GetData() + CRLF)
}

func (r *BulkReply) GetData() string {
	return string(r.data)
}

type MultiBulkReply struct {
	datas []respintf.ReplyIntf
}

func (r *MultiBulkReply) ToBytes() []byte {
	bs := []byte("*" + strconv.Itoa(len(r.datas)) + CRLF)
	for _, data := range r.datas {
		bs = append(bs, data.ToBytes()...)
	}
	return bs
}
