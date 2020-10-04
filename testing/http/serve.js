function handler(req) {
  console.log(`Method: ${req.Method}\nPath: ${req.URL.Path}`);
  return "Hello from Elsa :)"
}

Elsa.serve(":8080", handler);
