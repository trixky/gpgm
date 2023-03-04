<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import Config from '../config'

	let delay = Config.io.delay.default;
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
		<textarea cols={Config.io.input.cols} rows={Config.io.input.row} placeholder="" bind:value={input} />
	</div>
	<input type="number" min={Config.io.delay.min} max={Config.io.delay.max} value={delay} />
	<button on:click={handle_run}>Run</button>
	<div class="text-container">
		<h2>Output</h2>
		<textarea cols={Config.io.output.cols} rows={Config.io.output.row} placeholder="" value={output} readonly />
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
