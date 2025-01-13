## What is a Stack?
It's a LIFO(Last In, First Out) data structure that stores objects into a sort of "vertical tower"

## How you add or remove
To add to the top of the list use the .push() method (or the equivalent of this function in your favorite language)

To remove from the top use the .pop() method

## What's the complexity?
These .push() and .pop() methods only have a O(1) complexity

## Where can I use them?
Stacks can be used to:
    -Implement undo/redo functions in applications.
    -Moving back/fordward through browser history
    -Backtracking algorithms (maze, file directories)
    -Calling functions (call stack)

## Are there any drawbacks?
    -Restriction of size in Stack is a drawback and if they are full, you cannot add any more elements to the stack.
    -Stacks do not provide fast access to elements other than the top element.
    -Stacks do not support efficient searching, as you have to pop elements one by one until you find the element you are looking for.