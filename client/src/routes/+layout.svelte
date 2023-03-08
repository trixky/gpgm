<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { onMount } from 'svelte';
	import '../app.css';

	onMount(() => {
		// @ts-ignore
		// Go is loaded from the app.html (wasm)
		const goWasm = new Go();

		WebAssembly.instantiateStreaming(fetch('wasm/src/main.wasm'), goWasm.importObject).then(
			(result) => {
				goWasm.run(result.instance);
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
