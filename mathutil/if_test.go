package mathutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIfIntXX(t *testing.T) {
	Convey("TestIfIntXX", t, func() {
		So(1, ShouldEqual, IfInt(true, 1, 2))
		So(2, ShouldEqual, IfInt(false, 1, 2))

		So(1, ShouldEqual, IfInt64(true, 1, 2))
		So(2, ShouldEqual, IfInt64(false, 1, 2))

		So(1, ShouldEqual, IfInt32(true, 1, 2))
		So(2, ShouldEqual, IfInt32(false, 1, 2))

		So(1, ShouldEqual, IfInt16(true, 1, 2))
		So(2, ShouldEqual, IfInt16(false, 1, 2))

		So(1, ShouldEqual, IfInt8(true, 1, 2))
		So(2, ShouldEqual, IfInt8(false, 1, 2))
	})
}

func TestIfUintXX(t *testing.T) {
	Convey("TestIfUintXX", t, func() {
		So(1, ShouldEqual, IfUint(true, 1, 2))
		So(2, ShouldEqual, IfUint(false, 1, 2))

		So(1, ShouldEqual, IfUint64(true, 1, 2))
		So(2, ShouldEqual, IfUint64(false, 1, 2))

		So(1, ShouldEqual, IfUint32(true, 1, 2))
		So(2, ShouldEqual, IfUint32(false, 1, 2))

		So(1, ShouldEqual, IfUint16(true, 1, 2))
		So(2, ShouldEqual, IfUint16(false, 1, 2))

		So(1, ShouldEqual, IfUint8(true, 1, 2))
		So(2, ShouldEqual, IfUint8(false, 1, 2))
	})
}

func TestIfFloatXX(t *testing.T) {
	Convey("TestIfFloatXX", t, func() {
		So(1.0, ShouldEqual, IfFloat32(true, 1.0, 2.0))
		So(2.0, ShouldEqual, IfFloat32(false, 1.0, 2.0))

		So(1.0, ShouldEqual, IfFloat64(true, 1.0, 2.0))
		So(2.0, ShouldEqual, IfFloat64(false, 1.0, 2.0))
	})
}

func TestIfStr(t *testing.T) {
	Convey("TestIfStr", t, func() {
		So("1.0", ShouldEqual, IfStr(true, "1.0", "2.0"))
		So("2.0", ShouldEqual, IfStr(false, "1.0", "2.0"))
	})
}

func TestIfObj(t *testing.T) {
	Convey("TestIfObj", t, func() {
		So(1, ShouldEqual, IfObj(true, 1, 2))
		So(2, ShouldEqual, IfObj(false, 1, 2))

		So(1.0, ShouldEqual, IfObj(true, 1.0, 2.0))
		So(2.0, ShouldEqual, IfObj(false, 1.0, 2.0))

		So("1.0", ShouldEqual, IfObj(true, "1.0", "2.0"))
		So("2.0", ShouldEqual, IfObj(false, "1.0", "2.0"))

		So([]string{"1", "2"}, ShouldResemble, IfObj(true, []string{"1", "2"}, []string{"3", "4"}))
		So([]string{"3", "4"}, ShouldResemble, IfObj(false, []string{"1", "2"}, []string{"3", "4"}))

		So(map[string]string{"1": "2"}, ShouldResemble, IfObj(true, map[string]string{"1": "2"}, map[string]string{"3": "4"}))
		So(map[string]string{"3": "4"}, ShouldResemble, IfObj(false, map[string]string{"1": "2"}, map[string]string{"3": "4"}))

		type _Info struct {
			Name string
		}

		So(&_Info{Name: "1"}, ShouldResemble, IfObj(true, &_Info{Name: "1"}, &_Info{Name: "2"}))
		So(&_Info{Name: "2"}, ShouldResemble, IfObj(false, &_Info{Name: "1"}, &_Info{Name: "2"}))
	})
}
