
package main
// priorityQueue is a heap-based priority queue of items
type priorityQueue []*item

// item represents an item in the priority queue
type item struct {
    value    string  // value of the item
    priority int     // priority of the item (i.e., its distance from the start vertex)
}

// Len returns the number of items in the priority queue
func (pq priorityQueue) Len() int {
    return len(pq)
}