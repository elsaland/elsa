const fs = require("fs");

fs.readFile("main.go", "utf8", function (err, data) {
  if (err) throw err;
});
