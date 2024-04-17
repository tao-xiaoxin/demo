package uploads

import (
	"gotest"
	"testing"
)

//	func TestAdd(t *testing.T) {
//		type args struct {
//			a int
//			b int
//		}
//		tests := []struct {
//			name string
//			args args
//			want int
//		}{
//			{
//				name: "Adding positive numbers",
//				args: args{a: 2, b: 3},
//				want: 5,
//			},
//			{
//				name: "Adding negative numbers",
//				args: args{a: -2, b: -3},
//				want: -5,
//			},
//			{
//				name: "Adding positive and negative numbers",
//				args: args{a: 2, b: -3},
//				want: -1,
//			},
//			{
//				name: "Adding zero",
//				args: args{a: 2, b: 0},
//				want: 2,
//			},
//		}
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				if got := Add(tt.args.a, tt.args.b); got != tt.want {
//					t.Errorf("Add() = %v, want %v", got, tt.want)
//				}
//			})
//		}
//	}
//
//	func TestMul(t *testing.T) {
//		type args struct {
//			a int
//			b int
//		}
//		tests := []struct {
//			name string
//			args args
//			want int
//		}{
//			{
//				name: "结果为0",
//				args: args{a: 2, b: 0},
//				want: 0,
//			},
//			{
//				name: "结果为-1",
//				args: args{a: -1, b: 1},
//				want: -1,
//			},
//			{
//				name: "结果为1",
//				args: args{a: -1, b: -1},
//				want: 1,
//			},
//		}
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				if got := Mul(tt.args.a, tt.args.b); got != tt.want {
//					t.Errorf("Mul() = %v, want %v", got, tt.want)
//				}
//			})
//		}
//	}
//
// func TestSomething(t *testing.T) {
//
//	// assert equality
//	assert.Equal(t, 123, 123, "they should be equal")
//
//	// assert inequality
//	assert.NotEqual(t, 123, 456, "they should not be equal")
//
//	// assert for nil (good for errors)
//	assert.Nil(t, object)
//
//	// assert for not nil (good when you expect something)
//	if assert.NotNil(t, object) {
//
//		// now we know that object isn't nil, we are safe to make
//		// further assertions without causing any errors
//		assert.Equal(t, "Something", object.Value)
//
//	}
//
// }

//type addCase struct{ A, B, want int }
//
//func createAddTestCase(t *testing.T, c *addCase) {
//	// t.Helper()
//	if ans := Add(c.A, c.B); ans != c.want {
//		t.Fatalf("%d * %d expected %d, but %d got",
//			c.A, c.B, c.want, ans)
//	}
//
//}
//
//func TestAdd2(t *testing.T) {
//	createAddTestCase(t, &addCase{1, 1, 2})
//	createAddTestCase(t, &addCase{2, -3, -1})
//	createAddTestCase(t, &addCase{0, -1, 0}) // wrong case
//}

type addCaseV1 struct {
	name string
	A, B int
	want int
}

func createAddTestCaseV1(c *addCaseV1, t *testing.T) {
	t.Helper()
	t.Run(c.name, func(t *testing.T) {
		if ans := main.Add(c.A, c.B); ans != c.want {
			t.Fatalf("%s: %d + %d expected %d, but %d got",
				c.name, c.A, c.B, c.want, ans)
		}
	})
}

func TestAddV1(t *testing.T) {
	createAddTestCaseV1(&addCaseV1{"case 1", 1, 1, 2}, t)
	createAddTestCaseV1(&addCaseV1{"case 2", 2, -3, -1}, t)
	createAddTestCaseV1(&addCaseV1{"case 3", 0, -1, 0}, t)
}
