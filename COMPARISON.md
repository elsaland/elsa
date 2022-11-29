## Comparison

| Features                 | Elsa (unreleased) | Deno v1.3.3 | Node.js v14.4.0  |
| ------------------------ | ----------------- | ----------- | ---------------- |
| Language                 | Go                | Rust        | C++              |
| JS engine                | **Quickjs**       | V8          | V8               |
| Bundler                  | Yes               | Yes         | No               |
| HTTP imports             | Yes               | Yes         | No               |
| In-built package manager | No                | No          | Yes              |
| Compiling to executable  | **Yes**           | Yes         | No (third party) |
| Explicit imports         | Yes               | Yes         | No               |
| Secure by default        | Yes               | Yes         | No               |
| Binary size              | **~12mb** \*      | ~44mb \*    | ~68mb            |

\* Lacks data for `Intl` and `toLocaleString`

> The list is not complete and there is more to add, feel free to open a PR for the same.
