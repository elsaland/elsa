
function test() {
  return "test";
}

__console_write("Done 1");

__console_write("Done 2");

__console_write("Done 3");

const arrow = () => {
  return 1;
}

arrow(); // 1
test(); // "test"
