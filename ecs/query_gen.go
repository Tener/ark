package ecs

// Code generated by go generate; DO NOT EDIT.

type cursor struct {
	archetype int
	table     int
	index     uintptr
	maxIndex  int64
}

// Query0 is a query for 0 components.
// Use a [NewFilter0] to create one.
type Query0 struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	hasWithout bool
}

func newQuery0(world *World, mask Mask, without Mask, relations []RelationID) Query0 {
	components := make([]*componentStorage, 0)

	return Query0{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query0) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query0) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query0) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.world.unlock(q.lock)
}

func (q *Query0) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query0) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query0) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query0) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}

// Query1 is a query for 1 components.
// Use a [NewFilter1] to create one.
type Query1[A any] struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	columnA    *column
	hasWithout bool
}

func newQuery1[A any](world *World, mask Mask, without Mask, ids []ID, relations []RelationID) Query1[A] {
	components := make([]*componentStorage, 1)
	for i := range 1 {
		components[i] = &world.storage.components[ids[i].id]
	}

	return Query1[A]{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query1[A]) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query1[A]) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
func (q *Query1[A]) Get() *A {
	q.world.checkQueryGet(&q.cursor)
	return (*A)(q.columnA.Get(q.cursor.index))
}

// GetRelation returns the entity relation target of the component at the given index.
func (q *Query1[A]) GetRelation(index int) Entity {
	return q.components[index].columns[q.table.id].target
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query1[A]) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.columnA = nil
	q.world.unlock(q.lock)
}

func (q *Query1[A]) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query1[A]) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query1[A]) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query1[A]) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.columnA = q.components[0].columns[q.table.id]
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}

// Query2 is a query for 2 components.
// Use a [NewFilter2] to create one.
type Query2[A any, B any] struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	columnA    *column
	columnB    *column
	hasWithout bool
}

func newQuery2[A any, B any](world *World, mask Mask, without Mask, ids []ID, relations []RelationID) Query2[A, B] {
	components := make([]*componentStorage, 2)
	for i := range 2 {
		components[i] = &world.storage.components[ids[i].id]
	}

	return Query2[A, B]{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query2[A, B]) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query2[A, B]) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
func (q *Query2[A, B]) Get() (*A, *B) {
	q.world.checkQueryGet(&q.cursor)
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index))
}

// GetRelation returns the entity relation target of the component at the given index.
func (q *Query2[A, B]) GetRelation(index int) Entity {
	return q.components[index].columns[q.table.id].target
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query2[A, B]) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.columnA = nil
	q.columnB = nil
	q.world.unlock(q.lock)
}

func (q *Query2[A, B]) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query2[A, B]) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query2[A, B]) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query2[A, B]) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.columnA = q.components[0].columns[q.table.id]
	q.columnB = q.components[1].columns[q.table.id]
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}

// Query3 is a query for 3 components.
// Use a [NewFilter3] to create one.
type Query3[A any, B any, C any] struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	columnA    *column
	columnB    *column
	columnC    *column
	hasWithout bool
}

func newQuery3[A any, B any, C any](world *World, mask Mask, without Mask, ids []ID, relations []RelationID) Query3[A, B, C] {
	components := make([]*componentStorage, 3)
	for i := range 3 {
		components[i] = &world.storage.components[ids[i].id]
	}

	return Query3[A, B, C]{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query3[A, B, C]) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query3[A, B, C]) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
func (q *Query3[A, B, C]) Get() (*A, *B, *C) {
	q.world.checkQueryGet(&q.cursor)
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index))
}

// GetRelation returns the entity relation target of the component at the given index.
func (q *Query3[A, B, C]) GetRelation(index int) Entity {
	return q.components[index].columns[q.table.id].target
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query3[A, B, C]) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.columnA = nil
	q.columnB = nil
	q.columnC = nil
	q.world.unlock(q.lock)
}

func (q *Query3[A, B, C]) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query3[A, B, C]) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query3[A, B, C]) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query3[A, B, C]) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.columnA = q.components[0].columns[q.table.id]
	q.columnB = q.components[1].columns[q.table.id]
	q.columnC = q.components[2].columns[q.table.id]
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}

// Query4 is a query for 4 components.
// Use a [NewFilter4] to create one.
type Query4[A any, B any, C any, D any] struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	columnA    *column
	columnB    *column
	columnC    *column
	columnD    *column
	hasWithout bool
}

func newQuery4[A any, B any, C any, D any](world *World, mask Mask, without Mask, ids []ID, relations []RelationID) Query4[A, B, C, D] {
	components := make([]*componentStorage, 4)
	for i := range 4 {
		components[i] = &world.storage.components[ids[i].id]
	}

	return Query4[A, B, C, D]{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query4[A, B, C, D]) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query4[A, B, C, D]) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
func (q *Query4[A, B, C, D]) Get() (*A, *B, *C, *D) {
	q.world.checkQueryGet(&q.cursor)
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index))
}

// GetRelation returns the entity relation target of the component at the given index.
func (q *Query4[A, B, C, D]) GetRelation(index int) Entity {
	return q.components[index].columns[q.table.id].target
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query4[A, B, C, D]) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.columnA = nil
	q.columnB = nil
	q.columnC = nil
	q.columnD = nil
	q.world.unlock(q.lock)
}

func (q *Query4[A, B, C, D]) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query4[A, B, C, D]) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query4[A, B, C, D]) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query4[A, B, C, D]) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.columnA = q.components[0].columns[q.table.id]
	q.columnB = q.components[1].columns[q.table.id]
	q.columnC = q.components[2].columns[q.table.id]
	q.columnD = q.components[3].columns[q.table.id]
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}

// Query5 is a query for 5 components.
// Use a [NewFilter5] to create one.
type Query5[A any, B any, C any, D any, E any] struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	columnA    *column
	columnB    *column
	columnC    *column
	columnD    *column
	columnE    *column
	hasWithout bool
}

func newQuery5[A any, B any, C any, D any, E any](world *World, mask Mask, without Mask, ids []ID, relations []RelationID) Query5[A, B, C, D, E] {
	components := make([]*componentStorage, 5)
	for i := range 5 {
		components[i] = &world.storage.components[ids[i].id]
	}

	return Query5[A, B, C, D, E]{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query5[A, B, C, D, E]) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query5[A, B, C, D, E]) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
func (q *Query5[A, B, C, D, E]) Get() (*A, *B, *C, *D, *E) {
	q.world.checkQueryGet(&q.cursor)
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index)),
		(*E)(q.columnE.Get(q.cursor.index))
}

// GetRelation returns the entity relation target of the component at the given index.
func (q *Query5[A, B, C, D, E]) GetRelation(index int) Entity {
	return q.components[index].columns[q.table.id].target
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query5[A, B, C, D, E]) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.columnA = nil
	q.columnB = nil
	q.columnC = nil
	q.columnD = nil
	q.columnE = nil
	q.world.unlock(q.lock)
}

func (q *Query5[A, B, C, D, E]) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query5[A, B, C, D, E]) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query5[A, B, C, D, E]) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query5[A, B, C, D, E]) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.columnA = q.components[0].columns[q.table.id]
	q.columnB = q.components[1].columns[q.table.id]
	q.columnC = q.components[2].columns[q.table.id]
	q.columnD = q.components[3].columns[q.table.id]
	q.columnE = q.components[4].columns[q.table.id]
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}

// Query6 is a query for 6 components.
// Use a [NewFilter6] to create one.
type Query6[A any, B any, C any, D any, E any, F any] struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	columnA    *column
	columnB    *column
	columnC    *column
	columnD    *column
	columnE    *column
	columnF    *column
	hasWithout bool
}

func newQuery6[A any, B any, C any, D any, E any, F any](world *World, mask Mask, without Mask, ids []ID, relations []RelationID) Query6[A, B, C, D, E, F] {
	components := make([]*componentStorage, 6)
	for i := range 6 {
		components[i] = &world.storage.components[ids[i].id]
	}

	return Query6[A, B, C, D, E, F]{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query6[A, B, C, D, E, F]) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query6[A, B, C, D, E, F]) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
func (q *Query6[A, B, C, D, E, F]) Get() (*A, *B, *C, *D, *E, *F) {
	q.world.checkQueryGet(&q.cursor)
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index)),
		(*E)(q.columnE.Get(q.cursor.index)),
		(*F)(q.columnF.Get(q.cursor.index))
}

// GetRelation returns the entity relation target of the component at the given index.
func (q *Query6[A, B, C, D, E, F]) GetRelation(index int) Entity {
	return q.components[index].columns[q.table.id].target
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query6[A, B, C, D, E, F]) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.columnA = nil
	q.columnB = nil
	q.columnC = nil
	q.columnD = nil
	q.columnE = nil
	q.columnF = nil
	q.world.unlock(q.lock)
}

func (q *Query6[A, B, C, D, E, F]) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query6[A, B, C, D, E, F]) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query6[A, B, C, D, E, F]) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query6[A, B, C, D, E, F]) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.columnA = q.components[0].columns[q.table.id]
	q.columnB = q.components[1].columns[q.table.id]
	q.columnC = q.components[2].columns[q.table.id]
	q.columnD = q.components[3].columns[q.table.id]
	q.columnE = q.components[4].columns[q.table.id]
	q.columnF = q.components[5].columns[q.table.id]
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}

// Query7 is a query for 7 components.
// Use a [NewFilter7] to create one.
type Query7[A any, B any, C any, D any, E any, F any, G any] struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	columnA    *column
	columnB    *column
	columnC    *column
	columnD    *column
	columnE    *column
	columnF    *column
	columnG    *column
	hasWithout bool
}

func newQuery7[A any, B any, C any, D any, E any, F any, G any](world *World, mask Mask, without Mask, ids []ID, relations []RelationID) Query7[A, B, C, D, E, F, G] {
	components := make([]*componentStorage, 7)
	for i := range 7 {
		components[i] = &world.storage.components[ids[i].id]
	}

	return Query7[A, B, C, D, E, F, G]{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query7[A, B, C, D, E, F, G]) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query7[A, B, C, D, E, F, G]) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
func (q *Query7[A, B, C, D, E, F, G]) Get() (*A, *B, *C, *D, *E, *F, *G) {
	q.world.checkQueryGet(&q.cursor)
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index)),
		(*E)(q.columnE.Get(q.cursor.index)),
		(*F)(q.columnF.Get(q.cursor.index)),
		(*G)(q.columnG.Get(q.cursor.index))
}

// GetRelation returns the entity relation target of the component at the given index.
func (q *Query7[A, B, C, D, E, F, G]) GetRelation(index int) Entity {
	return q.components[index].columns[q.table.id].target
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query7[A, B, C, D, E, F, G]) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.columnA = nil
	q.columnB = nil
	q.columnC = nil
	q.columnD = nil
	q.columnE = nil
	q.columnF = nil
	q.columnG = nil
	q.world.unlock(q.lock)
}

func (q *Query7[A, B, C, D, E, F, G]) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query7[A, B, C, D, E, F, G]) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query7[A, B, C, D, E, F, G]) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query7[A, B, C, D, E, F, G]) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.columnA = q.components[0].columns[q.table.id]
	q.columnB = q.components[1].columns[q.table.id]
	q.columnC = q.components[2].columns[q.table.id]
	q.columnD = q.components[3].columns[q.table.id]
	q.columnE = q.components[4].columns[q.table.id]
	q.columnF = q.components[5].columns[q.table.id]
	q.columnG = q.components[6].columns[q.table.id]
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}

// Query8 is a query for 8 components.
// Use a [NewFilter8] to create one.
type Query8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	world      *World
	mask       Mask
	without    Mask
	relations  []RelationID
	lock       uint8
	cursor     cursor
	tables     []tableID
	table      *table
	components []*componentStorage
	columnA    *column
	columnB    *column
	columnC    *column
	columnD    *column
	columnE    *column
	columnF    *column
	columnG    *column
	columnH    *column
	hasWithout bool
}

func newQuery8[A any, B any, C any, D any, E any, F any, G any, H any](world *World, mask Mask, without Mask, ids []ID, relations []RelationID) Query8[A, B, C, D, E, F, G, H] {
	components := make([]*componentStorage, 8)
	for i := range 8 {
		components[i] = &world.storage.components[ids[i].id]
	}

	return Query8[A, B, C, D, E, F, G, H]{
		world:      world,
		mask:       mask,
		without:    without,
		relations:  relations,
		hasWithout: !without.IsZero(),
		lock:       world.lock(),
		components: components,
		cursor: cursor{
			archetype: -1,
			table:     -1,
			index:     0,
			maxIndex:  -1,
		},
	}
}

// Next advances the query's cursor to the next entity.
func (q *Query8[A, B, C, D, E, F, G, H]) Next() bool {
	q.world.checkQueryNext(&q.cursor)
	if int64(q.cursor.index) < q.cursor.maxIndex {
		q.cursor.index++
		return true
	}
	return q.nextTableOrArchetype()
}

// Entity returns the current entity.
func (q *Query8[A, B, C, D, E, F, G, H]) Entity() Entity {
	q.world.checkQueryGet(&q.cursor)
	return q.table.GetEntity(q.cursor.index)
}

// Get returns the queried components of the current entity.
func (q *Query8[A, B, C, D, E, F, G, H]) Get() (*A, *B, *C, *D, *E, *F, *G, *H) {
	q.world.checkQueryGet(&q.cursor)
	return (*A)(q.columnA.Get(q.cursor.index)),
		(*B)(q.columnB.Get(q.cursor.index)),
		(*C)(q.columnC.Get(q.cursor.index)),
		(*D)(q.columnD.Get(q.cursor.index)),
		(*E)(q.columnE.Get(q.cursor.index)),
		(*F)(q.columnF.Get(q.cursor.index)),
		(*G)(q.columnG.Get(q.cursor.index)),
		(*H)(q.columnH.Get(q.cursor.index))
}

// GetRelation returns the entity relation target of the component at the given index.
func (q *Query8[A, B, C, D, E, F, G, H]) GetRelation(index int) Entity {
	return q.components[index].columns[q.table.id].target
}

// Close closes the Query and unlocks the world.
//
// Automatically called when iteration finishes.
// Needs to be called only if breaking out of the query iteration or not iterating at all.
func (q *Query8[A, B, C, D, E, F, G, H]) Close() {
	q.cursor.archetype = -2
	q.cursor.table = -2
	q.tables = nil
	q.table = nil
	q.columnA = nil
	q.columnB = nil
	q.columnC = nil
	q.columnD = nil
	q.columnE = nil
	q.columnF = nil
	q.columnG = nil
	q.columnH = nil
	q.world.unlock(q.lock)
}

func (q *Query8[A, B, C, D, E, F, G, H]) nextTableOrArchetype() bool {
	if q.cursor.archetype >= 0 && q.nextTable() {
		return true
	}
	return q.nextArchetype()
}

func (q *Query8[A, B, C, D, E, F, G, H]) nextArchetype() bool {
	maxArchIndex := len(q.world.storage.archetypes) - 1
	for q.cursor.archetype < maxArchIndex {
		q.cursor.archetype++
		archetype := &q.world.storage.archetypes[q.cursor.archetype]
		if !archetype.mask.Contains(&q.mask) ||
			archetype.mask.ContainsAny(&q.without) {
			continue
		}

		if !archetype.HasRelations() {
			table := &q.world.storage.tables[archetype.tables[0]]
			if table.Len() > 0 {
				q.setTable(0, table)
				return true
			}
			continue
		}

		q.tables = archetype.GetTables(q.relations)
		q.cursor.table = -1
		if q.nextTable() {
			return true
		}
	}
	q.Close()
	return false
}

func (q *Query8[A, B, C, D, E, F, G, H]) nextTable() bool {
	maxTableIndex := len(q.tables) - 1
	for q.cursor.table < maxTableIndex {
		q.cursor.table++
		table := &q.world.storage.tables[q.tables[q.cursor.table]]
		if table.Len() == 0 {
			continue
		}
		if !table.Matches(q.relations) {
			continue
		}
		q.setTable(q.cursor.table, table)
		return true
	}
	return false
}

func (q *Query8[A, B, C, D, E, F, G, H]) setTable(index int, table *table) {
	q.cursor.table = index
	q.table = table
	q.columnA = q.components[0].columns[q.table.id]
	q.columnB = q.components[1].columns[q.table.id]
	q.columnC = q.components[2].columns[q.table.id]
	q.columnD = q.components[3].columns[q.table.id]
	q.columnE = q.components[4].columns[q.table.id]
	q.columnF = q.components[5].columns[q.table.id]
	q.columnG = q.components[6].columns[q.table.id]
	q.columnH = q.components[7].columns[q.table.id]
	q.cursor.index = 0
	q.cursor.maxIndex = int64(q.table.entities.Len() - 1)
}
