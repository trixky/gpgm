<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { onMount } from 'svelte';
	import { wasmReady } from '$lib/stores/ready';
	import '../app.css';

	export let data: { bytes: BufferSource };

	onMount(() => {
		// @ts-expect-error
		// Go is loaded from the app.html (wasm)
		const goWasm = new Go();

		WebAssembly.instantiate(data.bytes, goWasm.importObject).then((result) => {
			goWasm.run(result.instance);
			$wasmReady = true;
		});
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
