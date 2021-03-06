// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

type shaderDataType int

const (
	stFloatMat2x3 shaderDataType = gl.FLOAT_MAT2x3
	stFloatMat2x4 shaderDataType = gl.FLOAT_MAT2x4
	stFloatMat2   shaderDataType = gl.FLOAT_MAT2
	stFloatMat3x2 shaderDataType = gl.FLOAT_MAT3x2
	stFloatMat3x4 shaderDataType = gl.FLOAT_MAT3x4
	stFloatMat3   shaderDataType = gl.FLOAT_MAT3
	stFloatMat4x2 shaderDataType = gl.FLOAT_MAT4x2
	stFloatMat4x3 shaderDataType = gl.FLOAT_MAT4x3
	stFloatMat4   shaderDataType = gl.FLOAT_MAT4
	stFloatVec1   shaderDataType = gl.FLOAT
	stFloatVec2   shaderDataType = gl.FLOAT_VEC2
	stFloatVec3   shaderDataType = gl.FLOAT_VEC3
	stFloatVec4   shaderDataType = gl.FLOAT_VEC4
	stSampler2d   shaderDataType = gl.SAMPLER_2D
)

func (s shaderDataType) String() string {
	switch s {
	case stFloatMat2x3:
		return "mat2x3"
	case stFloatMat2x4:
		return "mat2x4"
	case stFloatMat2:
		return "mat2"
	case stFloatMat3x2:
		return "mat3x2"
	case stFloatMat3x4:
		return "mat3x4"
	case stFloatMat3:
		return "mat3"
	case stFloatMat4x2:
		return "mat4x2"
	case stFloatMat4x3:
		return "mat4x3"
	case stFloatMat4:
		return "mat4"
	case stFloatVec1:
		return "float"
	case stFloatVec2:
		return "vec2"
	case stFloatVec3:
		return "vec3"
	case stFloatVec4:
		return "vec4"
	case stSampler2d:
		return "sampler2D"
	default:
		return "unknown"
	}
}

func (s shaderDataType) sizeInBytes() int {
	return s.vectorElementCount() * s.vectorElementType().sizeInBytes()
}

func (s shaderDataType) vectorElementCount() int {
	switch s {
	case stFloatMat2x3:
		return 2 * 3
	case stFloatMat2x4:
		return 2 * 4
	case stFloatMat2:
		return 2 * 2
	case stFloatMat3x2:
		return 3 * 2
	case stFloatMat3x4:
		return 3 * 4
	case stFloatMat3:
		return 3 * 3
	case stFloatMat4x2:
		return 4 * 2
	case stFloatMat4x3:
		return 4 * 3
	case stFloatMat4:
		return 4 * 4
	case stFloatVec1:
		return 1
	case stFloatVec2:
		return 2
	case stFloatVec3:
		return 3
	case stFloatVec4:
		return 4
	case stSampler2d:
		return 1
	default:
		panic(fmt.Errorf("Unknown shaderDataType 0x%.4x", s))
	}
}

func (s shaderDataType) vectorElementType() primitiveType {
	switch s {
	case stFloatMat2x3:
		return ptFloat
	case stFloatMat2x4:
		return ptFloat
	case stFloatMat2:
		return ptFloat
	case stFloatMat3x2:
		return ptFloat
	case stFloatMat3x4:
		return ptFloat
	case stFloatMat3:
		return ptFloat
	case stFloatMat4x2:
		return ptFloat
	case stFloatMat4x3:
		return ptFloat
	case stFloatMat4:
		return ptFloat
	case stFloatVec1:
		return ptFloat
	case stFloatVec2:
		return ptFloat
	case stFloatVec3:
		return ptFloat
	case stFloatVec4:
		return ptFloat
	case stSampler2d:
		return ptInt
	default:
		panic(fmt.Errorf("Unknown shaderDataType 0x%.4x", s))
	}
}
