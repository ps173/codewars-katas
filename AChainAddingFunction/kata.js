function add(x) {
 // This is black magic
 let f = function(y) {
  return add(x + y);
 };

 f.valueOf = function() {
  return x;
 };

 return f;
}
console.log(add(1)(2)(3)(4).valueOf());
