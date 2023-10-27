package mathutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxXX(t *testing.T) {
	Convey("TestMaxXX", t, func() {
		So(2, ShouldEqual, MaxInt(1, 2))
		So(2, ShouldEqual, MaxInt64(1, 2))
		So(2, ShouldEqual, MaxInt32(1, 2))
		So(2, ShouldEqual, MaxInt16(1, 2))

		So(2, ShouldEqual, MaxUint(1, 2))
		So(2, ShouldEqual, MaxUint64(1, 2))
		So(2, ShouldEqual, MaxUint32(1, 2))
		So(2, ShouldEqual, MaxUint16(1, 2))

		So(2, ShouldEqual, MaxFloat64(1, 2))
		So(2, ShouldEqual, MaxFloat32(1, 2))
	})
}

func TestMinXX(t *testing.T) {
	Convey("TestMinXX", t, func() {
		So(1, ShouldEqual, MinInt(1, 2))
		So(1, ShouldEqual, MinInt64(1, 2))
		So(1, ShouldEqual, MinInt32(1, 2))
		So(1, ShouldEqual, MinInt16(1, 2))

		So(1, ShouldEqual, MinUint(1, 2))
		So(1, ShouldEqual, MinUint64(1, 2))
		So(1, ShouldEqual, MinUint32(1, 2))
		So(1, ShouldEqual, MinUint16(1, 2))

		So(1, ShouldEqual, MinFloat64(1, 2))
		So(1, ShouldEqual, MinFloat32(1, 2))
	})
}
