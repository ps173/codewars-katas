//try 1
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
domain_extractor("https://www.cnet.com");
domain_extractor("https://www.hyphen-site.com");
domain_extractor("www.xakep.ru");
