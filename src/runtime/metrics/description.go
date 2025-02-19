// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

// Description describes a runtime metric.
type Description struct {
	// Name is the full name of the metric which includes the unit.
	//
	// The format of the metric may be described by the following regular expression.
	//
	// 	^(?P<name>/[^:]+):(?P<unit>[^:*/]+(?:[*/][^:*/]+)*)$
	//
	// The format splits the name into two components, separated by a colon: a path which always
	// starts with a /, and a machine-parseable unit. The name may contain any valid Unicode
	// codepoint in between / characters, but by convention will try to stick to lowercase
	// characters and hyphens. An example of such a path might be "/memory/heap/free".
	//
	// The unit is by convention a series of lowercase English unit names (singular or plural)
	// without prefixes delimited by '*' or '/'. The unit names may contain any valid Unicode
	// codepoint that is not a delimiter.
	// Examples of units might be "seconds", "bytes", "bytes/second", "cpu-seconds",
	// "byte*cpu-seconds", and "bytes/second/second".
	//
	// For histograms, multiple units may apply. For instance, the units of the buckets and
	// the count. By convention, for histograms, the units of the count are always "samples"
	// with the type of sample evident by the metric's name, while the unit in the name
	// specifies the buckets' unit.
	//
	// A complete name might look like "/memory/heap/free:bytes".
	Name string

	// Description is an English language sentence describing the metric.
	Description string

	// Kind is the kind of value for this metric.
	//
	// The purpose of this field is to allow users to filter out metrics whose values are
	// types which their application may not understand.
	Kind ValueKind

	// Cumulative is whether or not the metric is cumulative. If a cumulative metric is just
	// a single number, then it increases monotonically. If the metric is a distribution,
	// then each bucket count increases monotonically.
	//
	// This flag thus indicates whether or not it's useful to compute a rate from this value.
	Cumulative bool
}

// The English language descriptions below must be kept in sync with the
// descriptions of each metric in doc.go.
var allDesc = []Description{
	{
		Name:        "/gc/cycles/automatic:gc-cycles",
		Description: "Count of completed GC cycles generated by the Go runtime.",
		Kind:        KindUint64,
		Cumulative:  true,
	},
	{
		Name:        "/gc/cycles/forced:gc-cycles",
		Description: "Count of completed GC cycles forced by the application.",
		Kind:        KindUint64,
		Cumulative:  true,
	},
	{
		Name:        "/gc/cycles/total:gc-cycles",
		Description: "Count of all completed GC cycles.",
		Kind:        KindUint64,
		Cumulative:  true,
	},
	{
		Name: "/gc/heap/allocs-by-size:bytes",
		Description: "Distribution of heap allocations by approximate size. " +
			"Note that this does not include tiny objects as defined by " +
			"/gc/heap/tiny/allocs:objects, only tiny blocks.",
		Kind:       KindFloat64Histogram,
		Cumulative: true,
	},
	{
		Name:        "/gc/heap/allocs:bytes",
		Description: "Cumulative sum of memory allocated to the heap by the application.",
		Kind:        KindUint64,
		Cumulative:  true,
	},
	{
		Name: "/gc/heap/allocs:objects",
		Description: "Cumulative count of heap allocations triggered by the application. " +
			"Note that this does not include tiny objects as defined by " +
			"/gc/heap/tiny/allocs:objects, only tiny blocks.",
		Kind:       KindUint64,
		Cumulative: true,
	},
	{
		Name: "/gc/heap/frees-by-size:bytes",
		Description: "Distribution of freed heap allocations by approximate size. " +
			"Note that this does not include tiny objects as defined by " +
			"/gc/heap/tiny/allocs:objects, only tiny blocks.",
		Kind:       KindFloat64Histogram,
		Cumulative: true,
	},
	{
		Name:        "/gc/heap/frees:bytes",
		Description: "Cumulative sum of heap memory freed by the garbage collector.",
		Kind:        KindUint64,
		Cumulative:  true,
	},
	{
		Name: "/gc/heap/frees:objects",
		Description: "Cumulative count of heap allocations whose storage was freed " +
			"by the garbage collector. " +
			"Note that this does not include tiny objects as defined by " +
			"/gc/heap/tiny/allocs:objects, only tiny blocks.",
		Kind:       KindUint64,
		Cumulative: true,
	},
	{
		Name:        "/gc/heap/goal:bytes",
		Description: "Heap size target for the end of the GC cycle.",
		Kind:        KindUint64,
	},
	{
		Name:        "/gc/heap/objects:objects",
		Description: "Number of objects, live or unswept, occupying heap memory.",
		Kind:        KindUint64,
	},
	{
		Name: "/gc/heap/tiny/allocs:objects",
		Description: "Count of small allocations that are packed together into blocks. " +
			"These allocations are counted separately from other allocations " +
			"because each individual allocation is not tracked by the runtime, " +
			"only their block. Each block is already accounted for in " +
			"allocs-by-size and frees-by-size.",
		Kind:       KindUint64,
		Cumulative: true,
	},
	{
		Name:        "/gc/pauses:seconds",
		Description: "Distribution individual GC-related stop-the-world pause latencies.",
		Kind:        KindFloat64Histogram,
		Cumulative:  true,
	},
	{
		Name:        "/gc/stack/starting-size:bytes",
		Description: "The stack size of new goroutines.",
		Kind:        KindUint64,
		Cumulative:  false,
	},
	{
		Name: "/memory/classes/heap/free:bytes",
		Description: "Memory that is completely free and eligible to be returned to the underlying system, " +
			"but has not been. This metric is the runtime's estimate of free address space that is backed by " +
			"physical memory.",
		Kind: KindUint64,
	},
	{
		Name:        "/memory/classes/heap/objects:bytes",
		Description: "Memory occupied by live objects and dead objects that have not yet been marked free by the garbage collector.",
		Kind:        KindUint64,
	},
	{
		Name: "/memory/classes/heap/released:bytes",
		Description: "Memory that is completely free and has been returned to the underlying system. This " +
			"metric is the runtime's estimate of free address space that is still mapped into the process, " +
			"but is not backed by physical memory.",
		Kind: KindUint64,
	},
	{
		Name:        "/memory/classes/heap/stacks:bytes",
		Description: "Memory allocated from the heap that is reserved for stack space, whether or not it is currently in-use.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/heap/unused:bytes",
		Description: "Memory that is reserved for heap objects but is not currently used to hold heap objects.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/metadata/mcache/free:bytes",
		Description: "Memory that is reserved for runtime mcache structures, but not in-use.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/metadata/mcache/inuse:bytes",
		Description: "Memory that is occupied by runtime mcache structures that are currently being used.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/metadata/mspan/free:bytes",
		Description: "Memory that is reserved for runtime mspan structures, but not in-use.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/metadata/mspan/inuse:bytes",
		Description: "Memory that is occupied by runtime mspan structures that are currently being used.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/metadata/other:bytes",
		Description: "Memory that is reserved for or used to hold runtime metadata.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/os-stacks:bytes",
		Description: "Stack memory allocated by the underlying operating system.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/other:bytes",
		Description: "Memory used by execution trace buffers, structures for debugging the runtime, finalizer and profiler specials, and more.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/profiling/buckets:bytes",
		Description: "Memory that is used by the stack trace hash map used for profiling.",
		Kind:        KindUint64,
	},
	{
		Name:        "/memory/classes/total:bytes",
		Description: "All memory mapped by the Go runtime into the current process as read-write. Note that this does not include memory mapped by code called via cgo or via the syscall package. Sum of all metrics in /memory/classes.",
		Kind:        KindUint64,
	},
	{
		Name:        "/sched/goroutines:goroutines",
		Description: "Count of live goroutines.",
		Kind:        KindUint64,
	},
	{
		Name:        "/sched/latencies:seconds",
		Description: "Distribution of the time goroutines have spent in the scheduler in a runnable state before actually running.",
		Kind:        KindFloat64Histogram,
	},
}

// All returns a slice of containing metric descriptions for all supported metrics.
func All() []Description {
	return allDesc
}
