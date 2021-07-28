function domain_extractor(url) {
  let matched = url.split(/:{0,1}\/{1,2}/g);
  let arraymat = matched[1].match(/\w[^(www)\.]+/g);
  console.log(arraymat[0]);
}

// function domain_extractor(url) {
//   let matched = url.match(/(?=[^(www)|(https:)])(\w+-\w+|\w+)/g);
//   console.log(matched[0]);
// }

domain_extractor("http://github.com/carbonfive/raygun");
domain_extractor("http://www.zombie-bites.com");
