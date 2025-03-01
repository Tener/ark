package ecs

import "reflect"

// ComponentID returns the [ID] for a component type via generics.
// Registers the type if it is not already registered.
//
// The number of unique component types per [World] is limited to 256 ([MaskTotalBits]).
//
// Panics if called on a locked world and the type is not registered yet.
//
// Note that type aliases are not considered separate component types.
// Type re-definitions, however, are separate types.
//
// ⚠️ Warning: Using IDs that are outside of the range of registered IDs anywhere in [World] or other places will result in undefined behavior!
func ComponentID[T any](w *World) ID {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	return w.componentID(tp)
}

// ComponentIDs returns a list of all registered component IDs.
func ComponentIDs(w *World) []ID {
	intIds := w.storage.registry.IDs
	ids := make([]ID, len(intIds))
	for i, iid := range intIds {
		ids[i] = id8(iid)
	}
	return ids
}

// ComponentInfo returns the [CompInfo] for a component [ID], and whether the ID is assigned.
func ComponentInfo(w *World, id ID) (CompInfo, bool) {
	tp, ok := w.storage.registry.ComponentType(id.id)
	if !ok {
		return CompInfo{}, false
	}

	return CompInfo{
		ID:         id,
		Type:       tp,
		IsRelation: w.storage.registry.IsRelation[id.id],
	}, true
}

// TypeID returns the [ID] for a component type.
// Registers the type if it is not already registered.
//
// The number of unique component types per [World] is limited to [MaskTotalBits].
func TypeID(w *World, tp reflect.Type) ID {
	return w.componentID(tp)
}

// Comp is a helper to pass component types to functions and methods.
// Use function [C] to create one.
type Comp struct {
	tp reflect.Type
}

// C creates a [Comp] instance for the given type.
func C[T any]() Comp {
	return Comp{typeOf[T]()}
}

// Type returns the reflect.Type of the component.
func (c Comp) Type() reflect.Type {
	return c.tp
}

// ResourceID returns the [ResID] for a resource type via generics.
// Registers the type if it is not already registered.
//
// The number of resources per [World] is limited to [MaskTotalBits].
func ResourceID[T any](w *World) ResID {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	return w.resourceID(tp)
}

// ResourceIDs returns a list of all registered resource IDs.
func ResourceIDs(w *World) []ResID {
	intIds := w.resources.registry.IDs
	ids := make([]ResID, len(intIds))
	for i, iid := range intIds {
		ids[i] = ResID{id: iid}
	}
	return ids
}

// ResourceType returns the reflect.Type for a resource [ResID], and whether the ID is assigned.
func ResourceType(w *World, id ResID) (reflect.Type, bool) {
	return w.resources.registry.ComponentType(id.id)
}
