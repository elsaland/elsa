build:
	go run ./bootstrap
	go build -o elsa .

benchmark:
	# console benchmarks
	hyperfine './elsa run ./testing/console.js' 'deno run ./testing/console.js' 'node ./testing/console.js' -s full -r 100 --warmup 50 --export-json ./benchmarks/console.json
	# bundler benchmarks
	hyperfine './elsa bundle ./testing/console.js' 'deno bundle ./testing/console.js' -s full -r 100 --warmup 50 --export-json ./benchmarks/bundle.json -i
	# readFile benchmarks
	hyperfine './elsa run ./testing/bench_scripts/fs.js --fs' 'deno run --allow-read ./testing/bench_scripts/fs_deno.js' 'node ./testing/bench_scripts/fs_node.js' --warmup 100 -s full -r 100 --export-json ./benchmarks/fs.json
	# PI benchmarks
	hyperfine 'deno run testing/pi.js' './elsa run testing/pi.js' 'node testing/pi.js' -s full -r 100 --warmup 50 --export-json ./benchmarks/pi.json -i

test:
	go test

clean-cache:
	rm -rf /tmp/x.nest.land/
	rm -rf /tmp/deno.land/

fmt:
	gofmt -w .
	prettier --write .

check-fmt:
	gofmt -l .
	prettier --check .

.PHONY: build benchmark clean-cache fmt check-fmt
