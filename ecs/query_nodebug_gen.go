//go:build !debug

package ecs

// Code generated by go generate; DO NOT EDIT.

// Next advances the query's cursor to the next entity.
func (q *Query0) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query0) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Next advances the query's cursor to the next entity.
func (q *Query1[A]) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query1[A]) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
//
// ⚠️ Do not store the obtained pointers outside of the current context (i.e. the query loop)!
func (q *Query1[A]) Get() *A {
	return (*A)(q.columnA.Get(q.cursor.index))
}

// Next advances the query's cursor to the next entity.
func (q *Query2[A, B]) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query2[A, B]) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
//
// ⚠️ Do not store the obtained pointers outside of the current context (i.e. the query loop)!
func (q *Query2[A, B]) Get() (*A, *B) {
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index))
}

// Next advances the query's cursor to the next entity.
func (q *Query3[A, B, C]) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query3[A, B, C]) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
//
// ⚠️ Do not store the obtained pointers outside of the current context (i.e. the query loop)!
func (q *Query3[A, B, C]) Get() (*A, *B, *C) {
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index))
}

// Next advances the query's cursor to the next entity.
func (q *Query4[A, B, C, D]) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query4[A, B, C, D]) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
//
// ⚠️ Do not store the obtained pointers outside of the current context (i.e. the query loop)!
func (q *Query4[A, B, C, D]) Get() (*A, *B, *C, *D) {
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index))
}

// Next advances the query's cursor to the next entity.
func (q *Query5[A, B, C, D, E]) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query5[A, B, C, D, E]) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
//
// ⚠️ Do not store the obtained pointers outside of the current context (i.e. the query loop)!
func (q *Query5[A, B, C, D, E]) Get() (*A, *B, *C, *D, *E) {
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index)),
		(*E)(q.columnE.Get(q.cursor.index))
}

// Next advances the query's cursor to the next entity.
func (q *Query6[A, B, C, D, E, F]) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query6[A, B, C, D, E, F]) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
//
// ⚠️ Do not store the obtained pointers outside of the current context (i.e. the query loop)!
func (q *Query6[A, B, C, D, E, F]) Get() (*A, *B, *C, *D, *E, *F) {
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index)),
		(*E)(q.columnE.Get(q.cursor.index)),
		(*F)(q.columnF.Get(q.cursor.index))
}

// Next advances the query's cursor to the next entity.
func (q *Query7[A, B, C, D, E, F, G]) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query7[A, B, C, D, E, F, G]) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
//
// ⚠️ Do not store the obtained pointers outside of the current context (i.e. the query loop)!
func (q *Query7[A, B, C, D, E, F, G]) Get() (*A, *B, *C, *D, *E, *F, *G) {
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index)),
		(*E)(q.columnE.Get(q.cursor.index)),
		(*F)(q.columnF.Get(q.cursor.index)),
		(*G)(q.columnG.Get(q.cursor.index))
}

// Next advances the query's cursor to the next entity.
func (q *Query8[A, B, C, D, E, F, G, H]) Next() bool {
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query8[A, B, C, D, E, F, G, H]) Entity() Entity {
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
//
// ⚠️ Do not store the obtained pointers outside of the current context (i.e. the query loop)!
func (q *Query8[A, B, C, D, E, F, G, H]) Get() (*A, *B, *C, *D, *E, *F, *G, *H) {
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index)),
		(*E)(q.columnE.Get(q.cursor.index)),
		(*F)(q.columnF.Get(q.cursor.index)),
		(*G)(q.columnG.Get(q.cursor.index)),
		(*H)(q.columnH.Get(q.cursor.index))
}
