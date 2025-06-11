let queue = [];

// Enqueue (add elements to the end)
queue.push(1);
queue.push(2);
queue.push(3);

console.log(queue); // Output: [1, 2, 3]

// Dequeue (remove elements from the front)
let firstElement = queue.shift();
console.log(firstElement); // Output: 1

let secondElement = queue.shift();
console.log(secondElement); // Output: 2

console.log(queue); // Output: [3]
