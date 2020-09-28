console.log("starting")

try { 
    Elsa.serve(":8080").then(x => console.log(1))
} catch(e) { throw new Error(e) }
