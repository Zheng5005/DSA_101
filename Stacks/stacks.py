##There's several ways to implement stacks in python, each one with it's own shortcomings
## List implementation
print("========================List=============================")

stack = []
stack.append('Kelsier')
stack.append('Vin')
stack.append('Sazed')
print(f"Initial stack {stack}")

print('\nElements popped from stack:')
print(stack.pop())
print(stack.pop())

print('\nStack after elements are popped:')
print(stack)
## The downside of this implementation is that you can run into speed issues as it grows

## collections.deque
print("========================Deque=============================")
from collections import deque
stackD = deque()
stackD.append('Kelsier')
stackD.append('Vin')
stackD.append('Sazed')
print(f"Initial stack {stackD}")

print('\nElements popped from stack:')
print(stackD.pop())
print(stackD.pop())

print('\nStack after elements are popped:')
print(stackD)
## This is preferred over list in cases where we need quicker append and pop
## Also this implementations give us an O(1) complexity

print("=====================Queue module=========================")
from queue import LifoQueue
stackQ = LifoQueue(maxsize=3)
stackQ.put('Kelsier')
stackQ.put('Vin')
stackQ.put('Sazed')
print("Full: ", stackQ.full())
print("Size: ", stackQ.qsize())

print('\nElements popped from the stack')
print(stackQ.get())
print(stackQ.get())
print(stackQ.get())

print("\nEmpty: ", stackQ.empty())