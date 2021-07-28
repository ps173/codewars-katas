function generateHashtag(str) {
  newStr = str.split(/\s+/g);
  str = str.split(/\s+/g).join("");
  if (str.length <= 140 && str.length > 0) {
    let genratedArray = [];
    newStr.forEach((word, index) => {
      let w = word[0].toUpperCase();
      genratedArray.push(w);
      genratedArray.push(word.substring(1));
    });
    genratedArray.unshift("#");
    let finalStr = genratedArray.join("");
    if (finalStr.length <= 140) {
      return finalStr;
    }
    return false;
  }
  return false;
}
console.log(generateHashtag(""));
console.log(generateHashtag("Codewarz is nice"));
console.log(
  generateHashtag("code" + " ".repeat(140) + "wars"),
);

console.log(
  generateHashtag(
    "Codewarz is nice and so coooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooolllllllllllllllllllllllllllllllllllllll",
  ),
  "should be false",
);

console.log(
  generateHashtag("a".repeat(139)),
  // "#A" + "a".repeat(138),
  "Should work",
);

console.log(
  generateHashtag("a".repeat(140)),
  // "#A" + "a".repeat(138),
  "Should not work",
);

console.log(generateHashtag("Do We have A Hashtag"), "should work");
