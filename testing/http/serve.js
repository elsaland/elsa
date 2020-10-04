function handler(req) {
    console.log(`Method: ${req.Method}\nPath: ${req.URL.Path}`)
}

Elsa.serve(":8080", handler);
