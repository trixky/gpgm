/// <reference types="@sveltejs/kit" />
/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />
/// <reference lib="webworker" />

import '/static/wasm/wasm_exec.js'

const sw = self as unknown as ServiceWorkerGlobalScope;

// @ts-expect-error Go is loaded with the wasm_exec import
const goWasm = new Go();

WebAssembly.instantiateStreaming(fetch('wasm/src/main.wasm'), goWasm.importObject).then((result) => {
	goWasm.run(result.instance);
});

// Use fetch events to run a new generation
sw.addEventListener('fetch', (event) => {
	// Ignore everything except /sw/generate
	if (event.request.method !== 'POST') return;
	const url = new URL(event.request.url);
	if (url.pathname !== '/sw/generate') return;

	// Execute the generation and return it's response
	async function respond() {
		const solver = await event.request.text()
		// @ts-expect-error WASM_run_generation is available in the global scope
		const result_wasm = WASM_run_generation(solver);
		return new Response(result_wasm, {
			status: 200,
		})
	}

	event.respondWith(respond());
});
