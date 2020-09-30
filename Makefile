build:
	go run ./bootstrap
	go build -o elsa .

benchmark:
	# console benchmarks
	hyperfine './elsa run ./testing/web/console.js' 'deno run ./testing/console.js' 'node ./testing/console.js' -s full -r 100 --warmup 50 --export-json ./benchmarks/console.json
	# bundle benchmarks
	hyperfine './elsa bundle ./testing/web/console.js' 'deno bundle ./testing/web/console.js' -s full -r 100 --warmup 50 --export-json ./benchmarks/bundle.json -i
	# readFile benchmarks
	hyperfine './elsa run ./testing/bench/fs.js --fs' 'deno run --allow-read ./testing/bench/fs_deno.js' 'node ./testing/bench/fs_node.js' --warmup 100 -s full -r 100 --export-json ./benchmarks/fs.json
	# PI benchmarks
	hyperfine 'deno run testing/bench/pi.js' './elsa run testing/bench/pi.js' 'node testing/bench/pi.js' -s full -r 100 --warmup 50 --export-json ./benchmarks/pi.json -i

test:
	go test

test-create-out:
	./elsa bundle testing/bundle/local_imports.js >> testing/bundle/local_imports.js.out
	./elsa bundle testing/bundle/hello.ts >> testing/bundle/hello.ts.out
	./elsa bundle testing/bundle/basics.js >> testing/bundle/basics.js.out

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
