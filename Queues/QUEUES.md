## What is a Queue
Is a FIFO (First In, First Out) data structure, like a line of people, also is a linear data structure

There's:
- Normal Queues
- Priority Queues

## What is a Priority Queue
Is a Queue where each data or value has a certain priority, is an extension of the queue with the following properties:
- An element with high priority is dequeue before an element with low priority.
- If two elements have the same priority, they are served according to their order in the queue
You can think about these like queues where the order of the FIFO structure is based on priority instead of the normal order

## How you add or remove
- Add -> enqueue, you'll add it to the end of the Queue
- Remove -> dequeue, you'll remove it from the head of the Queue

## What's the complexity?
### Queue
It depends on the method but more often then not they are O(1) complexity, thou, if you're using list, dequeue an element would be O(n) complexity
### Priority Queue 
Again it depends on the method, but in general it would be O(log(n)), but depeding on the method, dequeue would be O(n)

## Where can I use them?
### Queues
- Print Job Scheduling
- Messaging Systems
### Priority queues
- Job Scheduling algorithms
- CPU and Disk Scheduling
- Navigations Systems (for example the Dijkstra’s Algorithm, if you have experience with Networking think about OSPF redistribution protocol)

## Drawbacks
### Queues
- There's no priority, it process elements strictly in arrival order
- If you want to search for an element, it's pretty inefficient (meaning O(n) complexity)
### Priority Queues
- Slower insertions compared to normal Queues (it takes O(log(n)) vs O(1) in normal Queues)
- Every enqueue and dequeue re-orders all the elements, meaning it adds time to process all the changes
- Finding an element takes O(n) time
