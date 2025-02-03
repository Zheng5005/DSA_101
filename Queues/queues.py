##There's several ways to implement stacks in python, each one with it's own shortcomings
## List implementation
print("========================List=============================")

queue = []
queue.append('Kelsier')
queue.append('Vin')
queue.append('Sazed')
print(f"Initial queue {queue}")


print('\nElements popped from queue:')
print(queue.pop(0))
print(queue.pop(0))

print('\nQueue after elements are popped:')
print(queue)
## The downside of this implementation is that you can run into speed issues as it grows
## It also has a O(n) time complexity

## collections.deque
print("========================Deque=============================")
from collections import deque
queueD = deque()
queueD.append('Kelsier')
queueD.append('Vin')
queueD.append('Sazed')
print(f"Initial queue {queueD}")

print('\nElements popped from queue:')
print(queueD.popleft())
print(queueD.popleft())

print('\nQueue after elements are popped:')
print(queueD)
## This is preferred over list in cases where we need quicker append and pop
## Also this implementations give us an O(1) complexity

print("=====================Queue module=========================")
from queue import Queue
queueQ = Queue(maxsize=3)
queueQ.put('Kelsier')
queueQ.put('Vin')
queueQ.put('Sazed')
print("Full: ", queueQ.full())
print("Size: ", queueQ.qsize())

print('\nElements popped from the queue')
print(queueQ.get())
print(queueQ.get())
print(queueQ.get())

print("\nEmpty: ", queueQ.empty())
