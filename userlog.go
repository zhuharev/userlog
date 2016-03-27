package userlog

import (
	"encoding/binary"
	"github.com/fatih/color"
	"io"
	"net"
	"os"
	"time"
)

type UserLog struct {
	f *os.File
}

type Action struct {
	Time       []byte     //8
	Session    []byte     //16
	Ip         []byte     //4
	UserId     uint64     //8
	ActionType ActionType //2
	ObjectType ObjectType //1
	ObjectId   uint64     //8
} // 39

func NewAction(ssid []byte, ip string, uid uint64, at ActionType, ot ObjectType, oid uint64) (*Action, error) {
	a := new(Action)
	a.Time = time2b(time.Now())

	a.Session = ssid
	a.Ip = []byte(net.ParseIP(ip).To4())
	a.UserId = uid
	a.ActionType = at
	a.ObjectType = ot
	a.ObjectId = oid
	return a, nil
}

func time2b(t time.Time) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(t.UnixNano()))
	return buf
}

func NewLog(fpath string) (*UserLog, error) {
	f, e := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if e != nil {
		return nil, e
	}

	ul := new(UserLog)
	ul.f = f
	return ul, nil
}

func NewLogWriter(writer io.Writer) (*UserLog, error) {
	return nil, nil
}

// todo do only write to os file
func (l *UserLog) Add(a *Action) (e error) {
	e = l.writeErr(a.Time, e)
	e = l.writeErr(a.Session, e)
	e = l.writeErr(a.Ip, e)
	e = l.writeErr(i2b(a.UserId), e)
	e = l.writeErr(a.ActionType.ToBytes(), e)
	e = l.writeErr(a.ObjectType.ToBytes(), e)
	e = l.writeErr(i2b(a.ObjectId), e)
	return e
}

func (l *UserLog) writeErr(b []byte, e error) error {
	if e != nil {
		return e
	}
	color.Green("Write %d", len(b))
	_, e = l.f.Write(b)
	return e
}

func i2b(i uint64) []byte {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func ReadLog(fpath string) (*UserLog, error) {
	f, e := os.OpenFile(fpath, os.O_RDONLY, 0777)
	if e != nil {
		return nil, e
	}

	ul := new(UserLog)
	ul.f = f
	return ul, nil
}

func (ul *UserLog) PrintTMP() {
	var buf = make([]byte, 48)
	for _, e := ul.f.Read(buf); e == nil; _, e = ul.f.Read(buf) {

		color.Green("%d ", NewActionType(buf[36:38]))
		buf = make([]byte, 48)
	}
}
