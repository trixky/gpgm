<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { onMount } from 'svelte';
	import { wasmReady } from '$lib/stores/ready';
	import '../app.css';

	onMount(() => {
		// @ts-ignore
		// Go is loaded from the app.html (wasm)
		const goWasm = new Go();

		WebAssembly.instantiateStreaming(fetch('wasm/src/main.wasm'), goWasm.importObject).then(
			(result) => {
				goWasm.run(result.instance);
				$wasmReady = true;
			}
		);
	});
</script>

<!-- ---------------------------------------------- CONTENT -->
<svelte:head>
	<title>GPGM</title>
</svelte:head>

<main>
	<slot />
</main>

<!-- ---------------------------------------------- STYLE -->
