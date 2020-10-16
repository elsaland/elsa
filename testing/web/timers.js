setTimeout(() => console.log('...and this after'), 1000)
//setTimeout(() => console.log('...and this after'), 3000)
//setTimeout(() => console.log('...and this after'), 3000)
console.log('this comes before')
globalThis.__sendAsync(40, console.log, null);