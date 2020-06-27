// Code generated by execgen; DO NOT EDIT.
// Copyright 2020 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
)

// Remove unused warning.
var _ = colexecerror.InternalError

func newBoolAndHashAggAlloc(
	allocator *colmem.Allocator, allocSize int64,
) aggregateFuncAlloc {
	return &boolAndHashAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type boolAndHashAgg struct {
	sawNonNull bool
	vec        []bool
	nulls      *coldata.Nulls
	curIdx     int
	curAgg     bool
}

var _ aggregateFunc = &boolAndHashAgg{}

const sizeOfBoolAndHashAgg = int64(unsafe.Sizeof(boolAndHashAgg{}))

func (b *boolAndHashAgg) Init(groups []bool, vec coldata.Vec) {
	b.vec = vec.Bool()
	b.nulls = vec.Nulls()
	b.Reset()
}

func (b *boolAndHashAgg) Reset() {
	b.curIdx = 0
	b.nulls.UnsetNulls()
	// true indicates whether we are doing an AND aggregate or OR aggregate.
	// For bool_and the true is true and for bool_or the true is false.
	b.curAgg = true
}

func (b *boolAndHashAgg) CurrentOutputIndex() int {
	return b.curIdx
}

func (b *boolAndHashAgg) SetOutputIndex(idx int) {
	b.curIdx = idx
}

func (b *boolAndHashAgg) Compute(batch coldata.Batch, inputIdxs []uint32) {
	inputLen := batch.Length()
	vec, sel := batch.ColVec(int(inputIdxs[0])), batch.Selection()
	col, nulls := vec.Bool(), vec.Nulls()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {

			// TODO(yuzefovich): template out has nulls vs no nulls cases.
			isNull := nulls.NullAt(i)
			if !isNull {
				b.curAgg = b.curAgg && col[i]
				b.sawNonNull = true
			}

		}
	} else {
		col = col[:inputLen]
		for i := range col {

			// TODO(yuzefovich): template out has nulls vs no nulls cases.
			isNull := nulls.NullAt(i)
			if !isNull {
				b.curAgg = b.curAgg && col[i]
				b.sawNonNull = true
			}

		}
	}
}

func (b *boolAndHashAgg) Flush() {
	if !b.sawNonNull {
		b.nulls.SetNull(b.curIdx)
	} else {
		b.vec[b.curIdx] = b.curAgg
	}
	b.curIdx++
}

func (b *boolAndHashAgg) HandleEmptyInputScalar() {
	b.nulls.SetNull(0)
}

type boolAndHashAggAlloc struct {
	aggAllocBase
	aggFuncs []boolAndHashAgg
}

var _ aggregateFuncAlloc = &boolAndHashAggAlloc{}

func (a *boolAndHashAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfBoolAndHashAgg * a.allocSize)
		a.aggFuncs = make([]boolAndHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

func newBoolOrHashAggAlloc(
	allocator *colmem.Allocator, allocSize int64,
) aggregateFuncAlloc {
	return &boolOrHashAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type boolOrHashAgg struct {
	sawNonNull bool
	vec        []bool
	nulls      *coldata.Nulls
	curIdx     int
	curAgg     bool
}

var _ aggregateFunc = &boolOrHashAgg{}

const sizeOfBoolOrHashAgg = int64(unsafe.Sizeof(boolOrHashAgg{}))

func (b *boolOrHashAgg) Init(groups []bool, vec coldata.Vec) {
	b.vec = vec.Bool()
	b.nulls = vec.Nulls()
	b.Reset()
}

func (b *boolOrHashAgg) Reset() {
	b.curIdx = 0
	b.nulls.UnsetNulls()
	// false indicates whether we are doing an AND aggregate or OR aggregate.
	// For bool_and the false is true and for bool_or the false is false.
	b.curAgg = false
}

func (b *boolOrHashAgg) CurrentOutputIndex() int {
	return b.curIdx
}

func (b *boolOrHashAgg) SetOutputIndex(idx int) {
	b.curIdx = idx
}

func (b *boolOrHashAgg) Compute(batch coldata.Batch, inputIdxs []uint32) {
	inputLen := batch.Length()
	vec, sel := batch.ColVec(int(inputIdxs[0])), batch.Selection()
	col, nulls := vec.Bool(), vec.Nulls()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {

			// TODO(yuzefovich): template out has nulls vs no nulls cases.
			isNull := nulls.NullAt(i)
			if !isNull {
				b.curAgg = b.curAgg || col[i]
				b.sawNonNull = true
			}

		}
	} else {
		col = col[:inputLen]
		for i := range col {

			// TODO(yuzefovich): template out has nulls vs no nulls cases.
			isNull := nulls.NullAt(i)
			if !isNull {
				b.curAgg = b.curAgg || col[i]
				b.sawNonNull = true
			}

		}
	}
}

func (b *boolOrHashAgg) Flush() {
	if !b.sawNonNull {
		b.nulls.SetNull(b.curIdx)
	} else {
		b.vec[b.curIdx] = b.curAgg
	}
	b.curIdx++
}

func (b *boolOrHashAgg) HandleEmptyInputScalar() {
	b.nulls.SetNull(0)
}

type boolOrHashAggAlloc struct {
	aggAllocBase
	aggFuncs []boolOrHashAgg
}

var _ aggregateFuncAlloc = &boolOrHashAggAlloc{}

func (a *boolOrHashAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfBoolOrHashAgg * a.allocSize)
		a.aggFuncs = make([]boolOrHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}