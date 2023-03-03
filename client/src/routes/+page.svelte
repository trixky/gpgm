<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	// delay in ms
	const delay_min = 500; // put in me in config PLEASE
	const delay_max = 60000; // put in me in config PLEASE
	const delay_default = 2000; // put in me in config PLEASE

	let delay = delay_default;
	let input = '';
	let output = '';

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
<style lang="postcss">
	main {
		@apply text-center;
	}

	h1 {
		@apply mb-12 mt-4;
	}

	button {
		@apply mt-5 px-3 py-1;
	}

	input {
		@apply mt-5 px-3 py-1;
	}

	.text-container {
		@apply m-auto w-fit;
	}

	.text-container > h2 {
		@apply mt-0 text-left;
	}

	textarea {
		@apply p-3;
	}
</style>
