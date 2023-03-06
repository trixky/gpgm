<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import Config from '../config';
	import Visual from '../components/visual/visual.svelte';
	import GenerationStore from '../stores/generation';
	import ArgumentStore from '../stores/arguments';

	// ------------------------------ IO
	let input = '';
	let output = '';

	// ------------------------------ State
	let running = false;
	let stop = false;
	let stopped = false;

	$: disabled_reset = !running || !stopped;

	// ------------------------------ Loop
	let frame = 0; // Used to refresh the visualizator

	function new_generation() {
		if (stop || stopped) {
			stop = false;
			stopped = true;
			return;
		}
		GenerationStore.push_random($ArgumentStore.population);
		frame++;
		setTimeout(() => {
			// Recursive loop
			new_generation();
		}, 3);
	}

	// ------------------------------ Handlers
	// -------- Navigation
	function handle_top() {
		window.scrollTo({
			top: 0,
			behavior: 'smooth'
		});
	}

	function handle_bottom() {
		window.scrollTo({
			top: document.body.scrollHeight,
			behavior: 'smooth'
		});
	}

	// -------- State
	function handle_run() {
		if (!running) {
			running = true;
			stop = false;
			stopped = false;

			// @ts-ignore
			// Run is loaded from the layout (wasm)
			output = Run(input, $ArgumentStore.delay);

			new_generation();
		}
	}

	function handle_stop() {
		if (running) {
			stop = true;
			stopped = true;
		}
	}

	function handle_continue() {
		stop = false;
		stopped = false;

		new_generation();
	}

	function handle_reset() {
		if (running && stopped) {
			running = false;
			stop = false;
			stopped = false;

			GenerationStore.reset();
			frame++;
		}
	}
	// -------- Inputs
	function handle_generations(e: any) {
		ArgumentStore.update_generations(+e.target.value);
	}

	function handle_population(e: any) {
		ArgumentStore.update_population(+e.target.value);
	}

	function handle_deep(e: any) {
		ArgumentStore.update_deep(+e.target.value);
	}

	function handle_delay(e: any) {
		ArgumentStore.update_deep(+e.target.value);
	}

	// ------------------------------ Scrolling blocker
	// https://svelte.dev/repl/2bdbf66371a3418e9e3eda076df6e32d?version=3.18.1
	$: scrollable = !running || stopped;

	const wheel = (node: any, options: any) => {
		let { scrollable } = options;

		const handler = (e: any) => {
			if (!scrollable) e.preventDefault();
		};

		node.addEventListener('wheel', handler, { passive: false });

		return {
			update(options: any) {
				scrollable = options.scrollable;
			},
			destroy() {
				node.removeEventListener('wheel', handler, { passive: false });
			}
		};
	};
</script>

<!-- ---------------------------------------------- CONTENT -->
<main>
	<header>
		<h1>GPGM</h1>
		<p class="opacity-30">genetic process graph manager</p>
		<p>gen: {$ArgumentStore.generations}</p>
		<p>pop: {$ArgumentStore.population}</p>
		<p>dp: {$ArgumentStore.deep}</p>
		<p>delay: {$ArgumentStore.delay}</p>
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
	<div class="form-container">
		<div class="input-container">
			<input
				type="number"
				min={Config.io.generations.min}
				max={Config.io.generations.max}
				value={$ArgumentStore.generations}
				disabled={running}
				on:input={handle_generations}
			/>
			<p class="input-label">gen</p>
		</div>
		<div class="input-container">
			<input
				type="number"
				min={Config.io.population.min}
				max={Config.io.population.max}
				value={$ArgumentStore.population}
				disabled={running}
				on:input={handle_population}
			/>
			<p class="input-label">pop</p>
		</div>
		<div class="input-container">
			<input
				type="number"
				min={Config.io.deep.min}
				max={Config.io.deep.max}
				value={$ArgumentStore.deep}
				disabled={running}
				on:input={handle_deep}
			/>
			<p class="input-label">dp</p>
		</div>
		<div class="input-container">
			<input
				type="number"
				min={Config.io.delay.min}
				max={Config.io.delay.max}
				value={$ArgumentStore.delay}
				disabled={running}
				on:input={handle_delay}
			/>
			<p class="input-label">ms</p>
		</div>
	</div>
	<div class="state-container">
		<button class="side-button" on:click={handle_bottom}>Bottom</button>
		{#if !running}
			<button class="play-button" on:click={handle_run}>Run</button>
		{:else if !stopped}
			<button class="play-button" on:click={handle_stop}>Stop</button>
		{:else}
			<button class="play-button" on:click={handle_continue}>Continue</button>
		{/if}
		<button class="side-button" on:click={handle_reset} disabled={disabled_reset}>Reset</button>
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
	<Visual {frame} />
	<div class="state-container">
		<button class="side-button" on:click={handle_top} disabled={running && !stopped}>Top</button>
		{#if !running}
			<button class="play-button" on:click={handle_run}>Run</button>
		{:else if !stopped}
			<button class="play-button" on:click={handle_stop}>Stop</button>
		{:else}
			<button class="play-button" on:click={handle_continue}>Continue</button>
		{/if}
		<button class="side-button" on:click={handle_reset} disabled={disabled_reset}>Reset</button>
	</div>
</main>

<svelte:window use:wheel={{ scrollable }} />

<!-- ---------------------------------------------- STYLE -->
<style lang="postcss">
	/* ----------------------- Global */
	header {
		@apply relative w-fit m-auto mb-12 mt-8;
	}

	main {
		@apply text-center mb-16;
	}

	img {
		@apply absolute -top-[43px] -right-[0px];
		width: 100px;
		height: 100px;
	}

	/* ----------------------- Buttons */
	.state-container {
		@apply mb-4;
	}

	button {
		@apply mt-5 px-3 py-1;
	}

	.play-button {
		@apply w-24;
	}

	.side-button {
		@apply w-24;
	}

	/* ----------------------- Form/Inputs */
	.form-container {
		@apply relative w-fit flex m-auto;
	}

	.input-container {
		@apply relative;
	}

	input {
		@apply mt-5 px-3 py-1 w-[116px];
	}

	.input-label {
		@apply absolute top-[25px] opacity-20 right-[16px] text-right w-8;
	}

	input::-webkit-outer-spin-button,
	input::-webkit-inner-spin-button {
		-webkit-appearance: none;
		-moz-appearance: textfield;
		margin: 0;
	}

	/* ----------------------- Textarea */
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
