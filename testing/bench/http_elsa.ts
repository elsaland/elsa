const addr = Elsa.args[0] || "127.0.0.1:4501";
const body = "Hello World";

function handler(req) {
  const res = {
    body,
    status: 200,
  };
  return res;
}

console.log(`http://${addr}/`);

Elsa.serve(addr, handler);
