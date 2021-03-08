const fs = require("fs");

fs.readFile("main.go", "utf8", function (err) {
  if (err) throw err;
});
