const shouldExist = Elsa.exists("testing/sample.txt");
const shouldNotExist = Elsa.exists("testing/idonotexist.txt");

console.log(shouldExist);
console.log(shouldNotExist);
