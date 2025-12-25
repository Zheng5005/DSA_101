## What is a Linear Search?
Is search where you iterate through a collection one element at the time

## What's the complexity?
O(n)

## How it works?
1- Start at the first element (index 0) of the list
2- Compare the current element with the target value you are searching for
3- If they match, the search is successful, and the index of the element is returned
4- If they do not match, move to the next element
5- if the algorithm reaches the end of the list and no match is found, the search terminates unsuccessfully, often indicated by returning a value like -1

## When to use Linear Search
- When the data set is small
- When tha data is unsorted and sorting it first would be more time-consuming than the search itself
- When searching data structures that only support sequential access, such as linked list.

## Advantages
- Fast for searches of small to medium data sets
- Does not need to sort
- Useful for data structures that do not have random acces (Linked List)

## Drawbacks
- Slow for large data sets
- Time taken increases proportionally with data size
- Slower that algorithms like binary search for large, sorted data
