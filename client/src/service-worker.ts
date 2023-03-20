/// <reference types="@sveltejs/kit" />
/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />
/// <reference lib="webworker" />

import '/static/wasm/wasm_exec.js'

const sw = self as unknown as ServiceWorkerGlobalScope;

// @ts-expect-error Go is loaded with the wasm_exec import
const goWasm = new Go();

const wasmLoader = WebAssembly.instantiateStreaming(fetch('wasm/src/main.wasm'), goWasm.importObject).then((result) => {
	goWasm.run(result.instance);
});

// Set as ready when the WASM loaded
sw.addEventListener("install", () => {
	sw.skipWaiting()
});

// Directly claim clients to avoid page reload
sw.addEventListener("activate", (event) => {
	event.waitUntil(sw.clients.claim());
});

const routes = [
	'/sw/parse',
	'/sw/initialize',
	'/sw/generate',
	'/sw/output',
]

// Use fetch events to run a new generation
sw.addEventListener('fetch', (event) => {
	// Ignore everything except /sw/generate
	if (event.request.method !== 'POST') return;
	const url = new URL(event.request.url);
	if (!routes.includes(url.pathname)) return;

	// Execute the generation and return it's response
	async function respond() {
		await wasmLoader
		const route = url.pathname
		const input = await event.request.text()
		let body = ''

		if (route === '/sw/parse') {
			// @ts-expect-error WASM_parse_input is available in the global scope
			body = WASM_parse_input(input);
		} else if (route === '/sw/initialize') {
			// @ts-expect-error WASM_initialize is available in the global scope
			body = WASM_initialize(input);
		} else if (route === '/sw/generate') {
			// @ts-expect-error WASM_run_generation is available in the global scope
			body = WASM_run_generation(input);
		} else if (route === '/sw/output') {
			// @ts-expect-error WASM_generate_output is available in the global scope
			body = WASM_generate_output(input);
		}

		return new Response(body, {
			status: 200,
			headers: {
				"content-type": 'application/json'
			}
		})
	}

	event.respondWith(respond());
});
