package mapping

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testTagName = "key"

type Foo struct {
	Str                 string
	StrWithTag          string `key:"stringwithtag"`
	StrWithTagAndOption string `key:"stringwithtag,string"`
}

func TestDeferInt(t *testing.T) {
	i := 1
	s := "hello"
	number := struct {
		f float64
	}{
		f: 6.4,
	}
	cases := []struct {
		t      reflect.Type
		expect reflect.Kind
	}{
		{
			t:      reflect.TypeOf(i),
			expect: reflect.Int,
		},
		{
			t:      reflect.TypeOf(&i),
			expect: reflect.Int,
		},
		{
			t:      reflect.TypeOf(s),
			expect: reflect.String,
		},
		{
			t:      reflect.TypeOf(&s),
			expect: reflect.String,
		},
		{
			t:      reflect.TypeOf(number.f),
			expect: reflect.Float64,
		},
		{
			t:      reflect.TypeOf(&number.f),
			expect: reflect.Float64,
		},
	}

	for _, each := range cases {
		t.Run(each.t.String(), func(t *testing.T) {
			assert.Equal(t, each.expect, Deref(each.t).Kind())
		})
	}
}

func TestParseKeyAndOptionWithoutTag(t *testing.T) {
	var foo Foo
	rte := reflect.TypeOf(&foo).Elem()
	field, _ := rte.FieldByName("Str")
	key, options, err := parseKeyAndOptions(testTagName, field)
	assert.Nil(t, err)
	assert.Equal(t, "Str", key)
	assert.Nil(t, options)
}

func TestParseKeyAndOptionWithTagWithoutOption(t *testing.T) {
	var foo Foo
	rte := reflect.TypeOf(&foo).Elem()
	field, _ := rte.FieldByName("StrWithTag")
	key, options, err := parseKeyAndOptions(testTagName, field)
	assert.Nil(t, err)
	assert.Equal(t, "stringwithtag", key)
	assert.Nil(t, options)
}

func TestParseKeyAndOptionWithTagAndOption(t *testing.T) {
	var foo Foo
	rte := reflect.TypeOf(&foo).Elem()
	field, _ := rte.FieldByName("StrWithTagAndOption")
	key, options, err := parseKeyAndOptions(testTagName, field)
	assert.Nil(t, err)
	assert.Equal(t, "stringwithtag", key)
	assert.True(t, options.FromString)
}

func TestValidatePtrWithNonPtr(t *testing.T) {
	var foo string
	rve := reflect.ValueOf(foo)
	assert.NotNil(t, ValidatePtr(&rve))
}

func TestValidatePtrWithPtr(t *testing.T) {
	var foo string
	rve := reflect.ValueOf(&foo)
	assert.Nil(t, ValidatePtr(&rve))
}

func TestValidatePtrWithNilPtr(t *testing.T) {
	var foo *string
	rve := reflect.ValueOf(foo)
	assert.NotNil(t, ValidatePtr(&rve))
}

func TestValidatePtrWithZeroValue(t *testing.T) {
	var s string
	e := reflect.Zero(reflect.TypeOf(s))
	assert.NotNil(t, ValidatePtr(&e))
}

func TestSetValueNotSettable(t *testing.T) {
	var i int
	assert.NotNil(t, setValue(reflect.Int, reflect.ValueOf(i), "1"))
}

func TestParseKeyAndOptionsErrors(t *testing.T) {
	type Bar struct {
		OptionsValue string `key:",options=a=b"`
		DefaultValue string `key:",default=a=b"`
	}

	var bar Bar
	_, _, err := parseKeyAndOptions("key", reflect.TypeOf(&bar).Elem().Field(0))
	assert.NotNil(t, err)
	_, _, err = parseKeyAndOptions("key", reflect.TypeOf(&bar).Elem().Field(1))
	assert.NotNil(t, err)
}

func TestSetValueFormatErrors(t *testing.T) {
	type Bar struct {
		IntValue   int
		UintValue  uint
		FloatValue float32
		MapValue   map[string]interface{}
	}

	var bar Bar
	tests := []struct {
		kind   reflect.Kind
		target reflect.Value
		value  string
	}{
		{
			kind:   reflect.Int,
			target: reflect.ValueOf(&bar.IntValue).Elem(),
			value:  "a",
		},
		{
			kind:   reflect.Uint,
			target: reflect.ValueOf(&bar.UintValue).Elem(),
			value:  "a",
		},
		{
			kind:   reflect.Float32,
			target: reflect.ValueOf(&bar.FloatValue).Elem(),
			value:  "a",
		},
		{
			kind:   reflect.Map,
			target: reflect.ValueOf(&bar.MapValue).Elem(),
		},
	}

	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			err := setValue(test.kind, test.target, test.value)
			assert.NotEqual(t, errValueNotSettable, err)
			assert.NotNil(t, err)
		})
	}
}

func TestRepr(t *testing.T) {
	var (
		f32 float32 = 1.1
		f64         = 2.2
		i8  int8    = 1
		i16 int16   = 2
		i32 int32   = 3
		i64 int64   = 4
		u8  uint8   = 5
		u16 uint16  = 6
		u32 uint32  = 7
		u64 uint64  = 8
	)
	tests := []struct {
		v      interface{}
		expect string
	}{
		{
			nil,
			"",
		},
		{
			mockStringable{},
			"mocked",
		},
		{
			new(mockStringable),
			"mocked",
		},
		{
			newMockPtr(),
			"mockptr",
		},
		{
			true,
			"true",
		},
		{
			false,
			"false",
		},
		{
			f32,
			"1.1",
		},
		{
			f64,
			"2.2",
		},
		{
			i8,
			"1",
		},
		{
			i16,
			"2",
		},
		{
			i32,
			"3",
		},
		{
			i64,
			"4",
		},
		{
			u8,
			"5",
		},
		{
			u16,
			"6",
		},
		{
			u32,
			"7",
		},
		{
			u64,
			"8",
		},
		{
			[]byte(`abcd`),
			"abcd",
		},
		{
			mockOpacity{val: 1},
			"{1}",
		},
	}

	for _, test := range tests {
		t.Run(test.expect, func(t *testing.T) {
			assert.Equal(t, test.expect, Repr(test.v))
		})
	}
}

type mockStringable struct{}

func (m mockStringable) String() string {
	return "mocked"
}

type mockPtr struct{}

func newMockPtr() *mockPtr {
	return new(mockPtr)
}

func (m *mockPtr) String() string {
	return "mockptr"
}

type mockOpacity struct {
	val int
}
