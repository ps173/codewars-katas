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
