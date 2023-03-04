<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import Config from '../config';

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
	<header>
		<h1>GPGM</h1>
		<p class="opacity-30">genetic process graph manager</p>
	</header>

	<div class="text-container">
		<h2>Input</h2>
		<textarea
			cols={Config.io.input.cols}
			rows={Config.io.input.row}
			placeholder=""
			bind:value={input}
			autocorrect="off"
			autocapitalize="off"
			spellcheck="false"
		/>
		<img src="/mascot.png" alt="" />
	</div>
	<div class="run-container">
		<input type="number" min={Config.io.delay.min} max={Config.io.delay.max} value={delay} />
		<p class="ms">ms</p>
		<button on:click={handle_run}>Run</button>
	</div>
	<div class="text-container">
		<h2>Output</h2>
		<textarea
			cols={Config.io.output.cols}
			rows={Config.io.output.row}
			placeholder=""
			value={output}
			readonly
		/>
	</div>
</main>

<!-- ---------------------------------------------- STYLE -->
<style lang="postcss">
	header {
		@apply relative w-fit m-auto mb-12 mt-8;
	}

	main {
		@apply text-center;
	}

	img {
		@apply absolute -top-[43px] -right-[0px];
		width: 100px;
		height: 100px;
		filter: drop-shadow(12px 12px 2px rgba(0, 0, 0, 0.2));
	}

	button {
		@apply mt-5 px-3 py-1 transition-all duration-150;
	}

	button:hover {
		@apply bg-black;
	}

	input {
		@apply mt-5 px-3 py-1 w-[102px];
	}

	input::-webkit-outer-spin-button,
	input::-webkit-inner-spin-button {
		-webkit-appearance: none;
		-moz-appearance: textfield;
		margin: 0;
	}

	.run-container {
		@apply relative w-fit m-auto;
	}

	.ms {
		@apply absolute left-[70px] top-[25px];
	}

	.text-container {
		@apply relative m-auto w-fit;
	}

	.text-container > h2 {
		@apply mt-4 mb-2 text-left;
	}

	textarea {
		@apply relative p-3 z-10;
	}
</style>
