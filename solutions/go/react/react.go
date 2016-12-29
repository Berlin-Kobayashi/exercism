// Package react implements an API for reactive programming.
package react

import "errors"

const testVersion = 4

// New initializes a new Reactor.
func New() Reactor {
	return basicReactor{store: make(map[int]cellStore)}
}

type basicReactor struct {
	store map[int]cellStore
}

func (reactor basicReactor) CreateInput(value int) InputCell {
	cell := basicInputCell{indexedCell{len(reactor.store), &reactor}}
	reactor.store[cell.id] = cellStore{cell, value, make(map[int]func()), []func(){}}

	return cell
}

func (reactor basicReactor) CreateCompute1(cell Cell, callback func(int) int) ComputeCell {
	idxCell := getIndexedCell(cell)

	computeCell := basicComputeCell{indexedCell{len(reactor.store), &reactor}}
	reactor.store[computeCell.id] = cellStore{computeCell, 0, make(map[int]func()), []func(){}}

	boundCallback := func() {
		store := reactor.store[computeCell.id]
		newValue := callback(reactor.store[idxCell.id].value)
		if newValue != store.value {
			store.value = newValue
			reactor.store[computeCell.id] = store
		}
	}
	boundCallback()
	store := reactor.store[idxCell.id]
	store.outputCells[computeCell.id] = boundCallback
	reactor.store[idxCell.id] = store

	return computeCell
}

func (reactor basicReactor) CreateCompute2(cell1, cell2 Cell, callback func(int, int) int) ComputeCell {
	idxCell1 := getIndexedCell(cell1)
	idxCell2 := getIndexedCell(cell2)

	computeCell := basicComputeCell{indexedCell{len(reactor.store), &reactor}}
	reactor.store[computeCell.id] = cellStore{computeCell, 0, make(map[int]func()), []func(){}}

	boundCallback := func() {
		store := reactor.store[computeCell.id]
		newValue := callback(reactor.store[idxCell1.id].value, reactor.store[idxCell2.id].value)
		if newValue != store.value {
			store.value = newValue
			reactor.store[computeCell.id] = store
		}
	}
	boundCallback()
	store1 := reactor.store[idxCell1.id]
	store1.outputCells[computeCell.id] = boundCallback
	reactor.store[idxCell1.id] = store1
	store2 := reactor.store[idxCell2.id]
	store2.outputCells[computeCell.id] = boundCallback
	reactor.store[idxCell2.id] = store2

	return computeCell
}

func (reactor *basicReactor) cascadeChanges(cellID int) {
	cellToInitialValue := make(map[int]int)
	reactor.getValuesOfAllDependentCells(cellID, cellToInitialValue)

	var updateCells func(int)
	updateCells = func(cellID int) {
		for i, update := range reactor.store[cellID].outputCells {
			update()
			updateCells(i)
		}
	}
	updateCells(cellID)

	cellToUpdatedValue := make(map[int]int)
	reactor.getValuesOfAllDependentCells(cellID, cellToUpdatedValue)

	for i, initialValue := range cellToInitialValue {
		if initialValue != cellToUpdatedValue[i] {
			for _, callback := range reactor.store[i].callbacks {
				callback()
			}
		}
	}
}

func (reactor *basicReactor) getValuesOfAllDependentCells(cellID int, cellToValue map[int]int) {
	for i := range reactor.store[cellID].outputCells {
		cellToValue[i] = reactor.store[i].value
		reactor.getValuesOfAllDependentCells(i, cellToValue)
	}
}

type cellStore struct {
	cell        Cell
	value       int
	outputCells map[int]func()
	callbacks   []func()
}

type basicInputCell struct {
	indexedCell
}

func (cell basicInputCell) SetValue(value int) {
	store := cell.reactor.store[cell.id]
	store.value = value
	cell.reactor.store[cell.id] = store

	cell.reactor.cascadeChanges(cell.id)
}

type basicComputeCell struct {
	indexedCell
}

func (cell basicComputeCell) AddCallback(callback func(int)) CallbackHandle {
	store := cell.reactor.store[cell.id]
	boundCallback := func() {
		callback(cell.reactor.store[cell.id].value)
	}
	callbackHandle := basicCallbackHandle(len(store.callbacks))
	store.callbacks = append(store.callbacks, boundCallback)
	cell.reactor.store[cell.id] = store

	return callbackHandle
}

func (cell basicComputeCell) RemoveCallback(callbackHandle CallbackHandle) {
	handle, ok := callbackHandle.(basicCallbackHandle)
	if ok {
		store := cell.reactor.store[cell.id]
		store.callbacks[handle] = func() {}
		cell.reactor.store[cell.id] = store
	}
}

type indexedCell struct {
	id      int
	reactor *basicReactor
}

func (cell indexedCell) Value() int {
	return cell.reactor.store[cell.id].value
}

func getIndexedCell(cell Cell) indexedCell {
	var idxCell indexedCell
	switch c := cell.(type) {
	case basicInputCell:
		idxCell = c.indexedCell
	case basicComputeCell:
		idxCell = c.indexedCell
	default:
		panic(errors.New("Cannot create compute cell for unknown cell"))
	}

	return idxCell
}

type basicCallbackHandle int
