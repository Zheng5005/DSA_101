## What is a DynamicArray?
Is an array with a dynamic size that can be change depending of our necesities

## What's the complexity?
Is a constant complexity O(1) when you access to an elements when access by index

Is a Linear complexity O(n) when you Append/Push at the end (when resizing is needed), Inser/Delete at an index (due to shifting elements) and when you search by value

## Real world applications
- Storing user input
- Handling variable-legth data structures

## Advantages
- You can access a random element of the array, this is done in O(1)
- Good locality of reference and data cache utilization
- Easy to insert/delete elements at the end

## Drawbacks
- Wastes more memory
- Shifting elements is time consuming O(n)
- Expand/Shrinking the array is time consuming O(n)
