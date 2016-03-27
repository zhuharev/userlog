package userlog

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUserLog(t *testing.T) {
	Convey("test new ul", t, func() {
		ul, _ := NewLog("log")

		a, _ := NewAction([]byte("asdfghkgiotdforf"), "127.0.0.1", 1, AVisit, OPost, 12)
		e := ul.Add(a)
		So(e, ShouldBeNil)
	})

	Convey("test read", t, func() {
		ul, _ := ReadLog("log")
		ul.PrintTMP()
	})
}
