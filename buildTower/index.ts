/*
  Tower of 6 floors
[
  '     *     ', 
  '    ***    ', 
  '   *****   ', 
  '  *******  ', 
  ' ********* ', 
  '***********'
]

*/
export const towerBuilder = (n: number): string[] => {
  let resultArray = [];
  const lastRow = 2 * n - 1;
  while (n >= 1) {
    const stars = 2 * n - 1;
    const spaces = lastRow - stars;
    const halfSpaceString = " ".repeat(spaces / 2);
    const floor = halfSpaceString
      .concat("*".repeat(stars))
      .concat(halfSpaceString);
    resultArray.push(floor);
    n--;
  }
  resultArray = resultArray.reverse();
  return resultArray;
};

export const towerBuildSmall = (n: number): string[] => {
  return Array.from(
    { length: n },
    (_item, index) =>
      `${" ".repeat(n - index - 1)}${"*".repeat(2 * index + 1)}${" ".repeat(
        n - index - 1
      )}`
  );
};

function main() {
  console.log(towerBuildSmall(6));
}

main();
