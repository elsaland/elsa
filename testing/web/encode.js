const encoder = new TextEncoder();
const decoder = new TextDecoder();

const text = "Elsa";

const encoded = encoder.encode(text);
console.log(`Encoded ${encoded}`);
const decoded = decoder.decode(encoded);
console.log(`Decoded ${decoded}`);

if (text !== decoded) throw new Error("Decoding assertion failed.");

// atob & btoa
console.log(`btoa: ${btoa("Hello, world!")}`); // 'SGVsbG8sIHdvcmxkIQ=='
console.log(`atob: ${atob("SGVsbG8sIHdvcmxkIQ==")}`); // 'Hello, world!'
