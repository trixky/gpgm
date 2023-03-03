<!-- ---------------------------------------------- SCRIPT -->
<script>
	import { onMount } from 'svelte';

	// delay in ms
	const delay_min = 500; // put in me in config PLEASE
	const delay_max = 60000; // put in me in config PLEASE
	const delay_default = 2000; // put in me in config PLEASE

	let delay = delay_default;
	let input = '';
	let output = '';

	onMount(() => {
		// @ts-ignore
		const goWasm = new Go();

		WebAssembly.instantiateStreaming(fetch('wasm/src/main.wasm'), goWasm.importObject).then(
			(result) => {
				goWasm.run(result.instance);
			}
		);
	});

	function handle_run() {
		// @ts-ignore
		output = Run(input, delay);
	}
</script>

<!-- ---------------------------------------------- CONTENT -->
<main>
	<h1>KRPSIM</h1>

	<div class="text-container">
		<h2>Input</h2>
		<textarea cols="42" rows="10" placeholder="" bind:value={input} />
	</div>
	<input type="number" min={delay_min} max={delay_max} value={delay} />
	<button on:click={handle_run}>Run</button>
	<div class="text-container">
		<h2>Output</h2>
		<textarea cols="42" rows="10" placeholder="" value={output} readonly />
	</div>
</main>

<!-- ---------------------------------------------- STYLE -->
<style>
	main {
		text-align: center;
	}

	h1 {
		margin-bottom: 70px;
	}

	button {
		margin-top: 20px;
		padding: 3px 10px;
	}

	.text-container {
		margin: auto;
		width: fit-content;
	}

	.text-container > h2 {
		margin-top: 0;
		text-align: left;
	}

	textarea {
		padding: 10px;
	}
</style>
