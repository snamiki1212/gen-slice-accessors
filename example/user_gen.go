// Code generated by "go-gen-slice-accessors"; DO NOT EDIT.
// Based on information from https://github.com/snamiki1212/go-gen-slice-accessors

package main

// UserIDs
func (xs Users) UserIDs() []string {
	sli := make([]string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].UserID)
	}
	return sli
}

// Ints
func (xs Users) Ints() []int {
	sli := make([]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Int)
	}
	return sli
}

// IntPtrs
func (xs Users) IntPtrs() []*int {
	sli := make([]*int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].IntPtr)
	}
	return sli
}

// IntSlices
func (xs Users) IntSlices() [][]int {
	sli := make([][]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].IntSlice)
	}
	return sli
}

// IntSlicePtrs
func (xs Users) IntSlicePtrs() [][]*int {
	sli := make([][]*int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].IntSlicePtr)
	}
	return sli
}

// Bools
func (xs Users) Bools() []bool {
	sli := make([]bool, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Bool)
	}
	return sli
}

// BoolPtrs
func (xs Users) BoolPtrs() []*bool {
	sli := make([]*bool, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].BoolPtr)
	}
	return sli
}

// BoolSlices
func (xs Users) BoolSlices() [][]bool {
	sli := make([][]bool, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].BoolSlice)
	}
	return sli
}

// BoolSlicePtrs
func (xs Users) BoolSlicePtrs() [][]*bool {
	sli := make([][]*bool, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].BoolSlicePtr)
	}
	return sli
}

// Strs
func (xs Users) Strs() []string {
	sli := make([]string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Str)
	}
	return sli
}

// StrPtrs
func (xs Users) StrPtrs() []*string {
	sli := make([]*string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StrPtr)
	}
	return sli
}

// StrSlices
func (xs Users) StrSlices() [][]string {
	sli := make([][]string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StrSlice)
	}
	return sli
}

// StrSlicePtrs
func (xs Users) StrSlicePtrs() [][]*string {
	sli := make([][]*string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StrSlicePtr)
	}
	return sli
}

// Fn1s
func (xs Users) Fn1s() []func() (string) {
	sli := make([]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Fn1)
	}
	return sli
}

// FnPtr1s
func (xs Users) FnPtr1s() []*func() (string) {
	sli := make([]*func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnPtr1)
	}
	return sli
}

// FnSlice1s
func (xs Users) FnSlice1s() [][]func() (string) {
	sli := make([][]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnSlice1)
	}
	return sli
}

// FnSlice1Ptrs
func (xs Users) FnSlice1Ptrs() [][]*func() (string) {
	sli := make([][]*func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnSlice1Ptr)
	}
	return sli
}

// Fn2s
func (xs Users) Fn2s() []func() () {
	sli := make([]func() (), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Fn2)
	}
	return sli
}

// FnPtr2s
func (xs Users) FnPtr2s() []*func() () {
	sli := make([]*func() (), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnPtr2)
	}
	return sli
}

// FnSlice2s
func (xs Users) FnSlice2s() [][]func() () {
	sli := make([][]func() (), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnSlice2)
	}
	return sli
}

// FnSlice2Ptrs
func (xs Users) FnSlice2Ptrs() [][]*func() () {
	sli := make([][]*func() (), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnSlice2Ptr)
	}
	return sli
}

// Fn3s
func (xs Users) Fn3s() []func(func() (string)) (func() (int)) {
	sli := make([]func(func() (string)) (func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Fn3)
	}
	return sli
}

// FnPtr3s
func (xs Users) FnPtr3s() []*func(func() (string)) (func() (int)) {
	sli := make([]*func(func() (string)) (func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnPtr3)
	}
	return sli
}

// FnSlice3s
func (xs Users) FnSlice3s() [][]func(func() (string)) (func() (int)) {
	sli := make([][]func(func() (string)) (func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnSlice3)
	}
	return sli
}

// FnSlice3Ptrs
func (xs Users) FnSlice3Ptrs() [][]*func(func() (string)) (func() (int)) {
	sli := make([][]*func(func() (string)) (func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnSlice3Ptr)
	}
	return sli
}

// Fn4s
func (xs Users) Fn4s() []func(*func() (string)) (*func() (int)) {
	sli := make([]func(*func() (string)) (*func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Fn4)
	}
	return sli
}

// FnPtr4s
func (xs Users) FnPtr4s() []*func(*func() (string)) (*func() (int)) {
	sli := make([]*func(*func() (string)) (*func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnPtr4)
	}
	return sli
}

// FnSlice4s
func (xs Users) FnSlice4s() [][]func(*func() (string)) (*func() (int)) {
	sli := make([][]func(*func() (string)) (*func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnSlice4)
	}
	return sli
}

// FnSlice4Ptrs
func (xs Users) FnSlice4Ptrs() [][]*func(*func() (string)) (*func() (int)) {
	sli := make([][]*func(*func() (string)) (*func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].FnSlice4Ptr)
	}
	return sli
}

// Struct0s
func (xs Users) Struct0s() []DefinedStruct0 {
	sli := make([]DefinedStruct0, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Struct0)
	}
	return sli
}

// StructPtr0s
func (xs Users) StructPtr0s() []*DefinedStruct0 {
	sli := make([]*DefinedStruct0, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StructPtr0)
	}
	return sli
}

// StructSlice0s
func (xs Users) StructSlice0s() [][]DefinedStruct0 {
	sli := make([][]DefinedStruct0, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StructSlice0)
	}
	return sli
}

// StructSlice0Ptrs
func (xs Users) StructSlice0Ptrs() [][]*DefinedStruct0 {
	sli := make([][]*DefinedStruct0, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StructSlice0Ptr)
	}
	return sli
}

// Struct1s
func (xs Users) Struct1s() []DefinedStruct1 {
	sli := make([]DefinedStruct1, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Struct1)
	}
	return sli
}

// StructPtr1s
func (xs Users) StructPtr1s() []*DefinedStruct1 {
	sli := make([]*DefinedStruct1, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StructPtr1)
	}
	return sli
}

// StructSlice1s
func (xs Users) StructSlice1s() [][]DefinedStruct1 {
	sli := make([][]DefinedStruct1, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StructSlice1)
	}
	return sli
}

// StructSlice1Ptrs
func (xs Users) StructSlice1Ptrs() [][]*DefinedStruct1 {
	sli := make([][]*DefinedStruct1, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StructSlice1Ptr)
	}
	return sli
}

// Map1s
func (xs Users) Map1s() []map[string]int {
	sli := make([]map[string]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Map1)
	}
	return sli
}

// MapPtr1s
func (xs Users) MapPtr1s() []*map[string]int {
	sli := make([]*map[string]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].MapPtr1)
	}
	return sli
}

// MapSlice1s
func (xs Users) MapSlice1s() [][]map[string]int {
	sli := make([][]map[string]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].MapSlice1)
	}
	return sli
}

// MapSlice1Ptrs
func (xs Users) MapSlice1Ptrs() [][]*map[string]int {
	sli := make([][]*map[string]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].MapSlice1Ptr)
	}
	return sli
}

// Map2s
func (xs Users) Map2s() []map[string]func() (string) {
	sli := make([]map[string]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Map2)
	}
	return sli
}

// MapPtr2s
func (xs Users) MapPtr2s() []*map[string]func() (string) {
	sli := make([]*map[string]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].MapPtr2)
	}
	return sli
}

// MapSlice2s
func (xs Users) MapSlice2s() [][]map[string]func() (string) {
	sli := make([][]map[string]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].MapSlice2)
	}
	return sli
}

// MapSlice2Ptrs
func (xs Users) MapSlice2Ptrs() [][]*map[string]func() (string) {
	sli := make([][]*map[string]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].MapSlice2Ptr)
	}
	return sli
}

// Chan0s
func (xs Users) Chan0s() []chan int {
	sli := make([]chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Chan0)
	}
	return sli
}

// ChanPtr0s
func (xs Users) ChanPtr0s() []*chan int {
	sli := make([]*chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanPtr0)
	}
	return sli
}

// ChanSlice0s
func (xs Users) ChanSlice0s() [][]chan int {
	sli := make([][]chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanSlice0)
	}
	return sli
}

// ChanSlicePtr0s
func (xs Users) ChanSlicePtr0s() [][]*chan int {
	sli := make([][]*chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanSlicePtr0)
	}
	return sli
}

// Chan1s
func (xs Users) Chan1s() []chan func() (string) {
	sli := make([]chan func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Chan1)
	}
	return sli
}

// ChanPtr1s
func (xs Users) ChanPtr1s() []*chan func() (string) {
	sli := make([]*chan func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanPtr1)
	}
	return sli
}

// ChanSlice1s
func (xs Users) ChanSlice1s() [][]chan func() (string) {
	sli := make([][]chan func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanSlice1)
	}
	return sli
}

// ChanSlicePtr1s
func (xs Users) ChanSlicePtr1s() [][]*chan func() (string) {
	sli := make([][]*chan func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanSlicePtr1)
	}
	return sli
}

// ChanSendPtr0s
func (xs Users) ChanSendPtr0s() []*chan<- int {
	sli := make([]*chan<- int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanSendPtr0)
	}
	return sli
}

// ChanSendSlice0s
func (xs Users) ChanSendSlice0s() [][]chan<- int {
	sli := make([][]chan<- int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanSendSlice0)
	}
	return sli
}

// ChanSendSlicePtr0s
func (xs Users) ChanSendSlicePtr0s() [][]*chan<- int {
	sli := make([][]*chan<- int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanSendSlicePtr0)
	}
	return sli
}

// ChanRecv0s
func (xs Users) ChanRecv0s() []<-chan int {
	sli := make([]<-chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanRecv0)
	}
	return sli
}

// ChanRecvPtr0s
func (xs Users) ChanRecvPtr0s() []*<-chan int {
	sli := make([]*<-chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanRecvPtr0)
	}
	return sli
}

// ChanRecvSlice0s
func (xs Users) ChanRecvSlice0s() [][]<-chan int {
	sli := make([][]<-chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanRecvSlice0)
	}
	return sli
}

// ChanRecvSlicePtr0s
func (xs Users) ChanRecvSlicePtr0s() [][]*<-chan int {
	sli := make([][]*<-chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanRecvSlicePtr0)
	}
	return sli
}
