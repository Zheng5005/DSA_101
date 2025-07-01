
## What is a LinkedList?
They have a big advantage, they can add or remove elements much faster then an array for example, they have a "data" and an "address" part, the data bbeing the information stored, and the address points to the next set of "data" and "address" of the list

## What is a doubly LinkedList
I'ts a type of LinkedList, normal LinkedList are singly, meaning they have 1 "address" part, but a doubly LinkedList has 2 "address" parts, one for the previous set of data, and the other for the next set of data, but the doubly LinkedList uses double of the memory compare with the singly LinkedList.

## How you add or remove
You add an element by just changing the address part of the previous element, the address should point to the new element, and the "address" part of the new element should point to the next element

And to remove, you change the "address" part of the previous element, the new address should point to the next element of the recent removed element.

## What's the complexity?
Linear complexity or O(n) for insertion at the end, at n position, deletion at the end, delition at position or searching

And Constant complexity or O(1) for insertion at the beginning, deletion at the beginning.

## Where can I use them?
LinkedList can be used to:
    -When you need O(1) time insertion/deletion (like real-time computing)
    -Don't know how many items will be in the list
    -Don't need random access to any element
    -Want to be able to insert items in the middle of the list

## Real world applications
    -Song in a music player that are linked to the previous and next songs
    -previous and next web page URLs (doubly LinkedList)
    -Implementing Queues and Deques because of fast delition or insertions from the front of the LinkedList

## Drawbacks
    -They are bad at searching an element (linear complexity)
    -Pointers or references make them more difficult to debug and maintain
    -Higher overhead because they have to store the reference to the next node, requiring more memory
    -Cache inefficiency because the memory is not contiguous. This means that when you traverse a linked list, you are not likely to get the data you need in the cache, leading to cache misses and slow performance.
