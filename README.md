GO MEMORY MANAGEMENT

There are 3 places memory can be allocated: the stack, the heap, and the data segment.

Typically, a functions parameters and local variables are allocated on the stack.

Each goroutine has its own stack; thus, no synchronization (e.g., locking) is necessary.

Goroutine stacks are allocated on the heap. If the stack needs to grow beyond the amount allocated for it, then heap operations (allocate new, copy old to new, free old) will occur.

Heap
Some memory allocations are put on the heap. Unlike the stack, the heap does not have a single partition of allocated and free regions.

Unlike the stack, the heap is not owned by one goroutine, so manipulating the set of free regions in the heap requires synchronization (e.g., locking).

Data Segment
Memory can also be allocated in the data segment. This is where global variables are stored. The data segement is defined at compile time and therefore does not grow and shrink at runtime.
The Go Programming Language Specification does not define where items will be allocated. For example, a variable defined as var x int could be allocated on the stack or the heap and still follow the language spec. Likewise, the integer pointed to by p in p := new(int) could be allocated on the stack or the heap.
The size of the data segment cannot change at run time, and therefore cannot be used for data structures that change size.
 
Memory allocated by a function can escape that function if referenced by an item outside the scope of the function and therefore cannot be allocated on the stack (because it’s still being referenced), and neither can it be allocated in the data segment (because the data segment cannot grow at runtime), thus it must be allocated on the heap – although inlining can remove some of these heap allocations.
 
Escape Analysis
Escape analysis is used to determine whether an item can be allocated on the stack.
 
 
 
 
 
 
 


