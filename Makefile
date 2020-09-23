build:
	go run ./bootstrap
	go build -o elsa .

benchmark:
	# console benchmarks
	hyperfine './elsa run ./testing/console.js' 'deno run ./testing/console.js' 'node ./testing/console.js' -s full -r 100 --warmup 50 --export-json ./benchmarks/console.json
	# bundle benchmarks
	hyperfine './elsa bundle ./testing/console.js' 'deno bundle ./testing/console.js' -s full -r 100 --warmup 50 --export-json ./benchmarks/bundle.json -i
	# readFile benchmarks
	hyperfine './elsa run ./testing/bench_scripts/fs.js --fs' 'deno run --allow-read ./testing/bench_scripts/fs_deno.js' 'node ./testing/bench_scripts/fs_node.js' --warmup 100 -s full -r 100 --export-json ./benchmarks/fs.json