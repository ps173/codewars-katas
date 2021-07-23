function countBits(n) {
  binary = (n >>> 0).toString(2);
  return (binary.match(/1/g) || []).length;
}

function tobin(n) {
  binary = (n >>> 0).toString(2);
  return binary;
}

console.log(countBits(1234), 5);
console.log(countBits(0), 0);
console.log(countBits(4), 1);
console.log(countBits(7), 3);
console.log(countBits(9), 2);
console.log(countBits(10), 2);
console.log(tobin(6795709331), 16);
