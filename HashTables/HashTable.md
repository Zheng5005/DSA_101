## What is a Hash Table?
Is a data structure that stores data as key-values pairs using a hash function to convert a key into an array index, this makes for a incredibly fast data insertion, searching and delition.

## What's the complexity?
More often than not is Linear complexity (O(n))

## How it works?
- Data is stored as pairs, (e.g. {"apple": 1, "banana": 2})
- A mathematical function takes the key (e.g. "apple") and computes a numerical hash code.
- The hash code is used to find a specific index in an underlying array (or the "buckets")
- The value is stored at that calculated index
- To find a value, you hash the key again, get the index and directly retrive the item, bypassing slow sequential searches.

## Real world examples:
- Phone book lookups
- Database indexing
- Caching web content

This is a video where this concept is well explained:
https://www.youtube.com/watch?v=JCXLwfKMWOU
