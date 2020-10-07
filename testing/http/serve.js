function handler(req) {
  console.log(`Method: ${req.Method}\nPath: ${req.URL.Path}`);
  return { body: "Hello from Elsa :)", status: 200 };
}

Elsa.serve(":8080", handler);
