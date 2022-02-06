/*
 
16  -->  1 + 6 = 7
942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6
493193  -->  4 + 9 + 3 + 1 + 9 + 3 = 29  -->  2 + 9 = 11  -->  1 + 1 = 2

*/

// NOTE : HERE'S REAL LIFE CHEAT CODE : https://mathworld.wolfram.com/DigitalRoot.html

export const digitalRoot = (n: number): number => {
  let digitalRoot = 0;
  while (n > 0 || digitalRoot > 9) {
    if (n === 0) {
      n = digitalRoot;
      digitalRoot = 0;
    }
    digitalRoot += n % 10;
    // @ts-ignore
    n = parseInt(n / 10, 10);
    console.log("n : ", n, "root : ", digitalRoot);
  }
  return digitalRoot;
};

const ret = digitalRoot(132189);
console.log(ret);
