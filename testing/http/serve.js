function handler(req) {
  console.log(req);
  console.log(`Method: ${req.Method}\nPath: ${req.Path}`);
  return { body: "Hello from Elsa :)", status: 200 };
}

Elsa.serve(":8080", handler);
