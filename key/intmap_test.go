// (c) 2022-2022, LDC Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package key

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntMap(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	type Str string
	m := IntMap{
		1:  int(1),
		2:  int64(2),
		3:  int32(3),
		-1: int(-1),
		-2: int64(-2),
		-3: int32(-3),
		0:  math.MaxInt64,
		10: []byte{1, 2, 3, 4},
		11: ByteStr{1, 2, 3, 4},
		12: [4]byte{1, 2, 3, 4},
		13: "hello",
		14: []string{"hello"},
		15: Str("hello"),
	}

	smallInt, err := m.GetInt(1)
	assert.NoError(err)
	assert.Equal(1, smallInt)

	smallInt, err = m.GetInt(-1)
	assert.NoError(err)
	assert.Equal(-1, smallInt)

	smallInt, err = m.GetInt(0)
	assert.Error(err)
	assert.Equal(0, smallInt)

	smallInt, err = m.GetInt(-10)
	assert.NoError(err)
	assert.Equal(0, smallInt)

	smallInt, err = m.GetInt(10)
	assert.Error(err)
	assert.Equal(0, smallInt)

	vInt, err := m.GetInt64(1)
	assert.NoError(err)
	assert.Equal(int64(1), vInt)

	vInt, err = m.GetInt64(-1)
	assert.NoError(err)
	assert.Equal(int64(-1), vInt)

	vInt, err = m.GetInt64(0)
	assert.NoError(err)
	assert.Equal(int64(math.MaxInt64), vInt)

	vInt, err = m.GetInt64(10)
	assert.Error(err)
	assert.Equal(int64(0), vInt)

	vInt, err = m.GetInt64(-10)
	assert.NoError(err)
	assert.Equal(int64(0), vInt)

	vUint, err := m.GetUint64(1)
	assert.NoError(err)
	assert.Equal(uint64(1), vUint)

	vUint, err = m.GetUint64(-1)
	assert.Error(err)
	assert.Equal(uint64(0), vUint)

	vUint, err = m.GetUint64(-10)
	assert.NoError(err)
	assert.Equal(uint64(0), vUint)

	vUint, err = m.GetUint64(0)
	assert.NoError(err)
	assert.Equal(uint64(math.MaxInt64), vUint)

	vUint, err = m.GetUint64(10)
	assert.Error(err)
	assert.Equal(uint64(0), vUint)

	vb, err := m.GetBytes(1)
	assert.Error(err)
	assert.Nil(vb)

	vb, err = m.GetBytes(-1)
	assert.Error(err)
	assert.Nil(vb)

	vb, err = m.GetBytes(-10)
	assert.NoError(err)
	assert.Nil(vb)

	vb, err = m.GetBytes(10)
	assert.NoError(err)
	assert.Equal([]byte{1, 2, 3, 4}, vb)

	vb, err = m.GetBytes(11)
	assert.NoError(err)
	assert.Equal([]byte{1, 2, 3, 4}, vb)

	vb, err = m.GetBytes(12)
	assert.Error(err)
	assert.Nil(vb)

	vb, err = m.GetBytes(13)
	assert.Error(err)
	assert.Nil(vb)

	vb, err = m.GetBytes(14)
	assert.Error(err)
	assert.Nil(vb)

	vs, err := m.GetString(1)
	assert.Error(err)
	assert.Equal("", vs)

	vs, err = m.GetString(-1)
	assert.Error(err)
	assert.Equal("", vs)

	vs, err = m.GetString(-10)
	assert.NoError(err)
	assert.Equal("", vs)

	vs, err = m.GetString(13)
	assert.NoError(err)
	assert.Equal("hello", vs)

	vs, err = m.GetString(14)
	assert.Error(err)
	assert.Equal("", vs)

	vs, err = m.GetString(15)
	assert.NoError(err)
	assert.Equal("hello", vs)

	data, err := MarshalCBOR(m)
	require.NoError(err)
	// CBOR Diagnostic:
	// {0: 9223372036854775807, 1: 1, 2: 2, 3: 3, 10: h'01020304', 11: h'01020304', 12: h'01020304', 13: "hello", 14: ["hello"], 15: "hello", -1: -1, -2: -2, -3: -3}
	assert.Equal(`ad001b7fffffffffffffff0101020203030a44010203040b44010203040c44010203040d6568656c6c6f0e816568656c6c6f0f6568656c6c6f202021212222`, ByteStr(data).String())
}
