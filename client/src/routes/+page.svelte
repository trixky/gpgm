<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import type { RunningSolver, WASMGenerationReturn } from '../types';
	import type { ScoredInstance } from '../types/population';
	import Config from '$lib/config';
	import Visual from '$lib/components/visual/visual.svelte';
	import GenerationStore from '$lib/stores/generation';
	import ArgumentStore from '$lib/stores/arguments';
	import StatisticStore from '$lib/stores/statistic';
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import examples from '$lib/Examples';
	import { scale } from 'svelte/transition';
	import { parse_as } from '$lib/utils/parse';

	// ------------------------------ IO
	let selectedExample = 0;
	let customInput = '';
	let input = '';
	let output = '';
	let outputFile = '';
	let lastError: string | null = null;

	// ------------------------------ State
	let running = false;
	let stop = false;
	let stopped = false;
	let finished = false;
	let allTimeBest: ScoredInstance | null = null;

	$: disabled_reset = !running || !stopped;

	// ------------------------------ Loop
	let frame = 0; // Used to refresh the visualizator
	let generation = 0;
	let result_wasm_json: WASMGenerationReturn | undefined = undefined;

	function new_generation() {
		generation++;
		if (generation >= $ArgumentStore.generations) {
			finished = true;
			stopped = true;
		}

		if (stop || stopped) {
			stop = false;
			stopped = true;
			return;
		}

		GenerationStore.push_random($ArgumentStore.population);
		frame++;
		setTimeout(() => {
			// Recursive loop

			const result_wasm = WASM_run_generation(JSON.stringify(result_wasm_json!.running_solver));
			result_wasm_json = parse_as<WASMGenerationReturn>(result_wasm);

			console.log(
				result_wasm_json.scored_population.instances[0].instance.chromosome.entry_gene.Process_ids
			);

			// const processes = result_wasm_json.scored_population?.instances[0]?.simulation?.history
			// 	?.map((process: any) => {
			// 		return `cycle:${process.cycle}\t\t${process.process.name}\t(${process.amount})`;
			// 	})
			// 	.join('\n');

			const best = result_wasm_json.scored_population.instances[0];
			if (!allTimeBest || allTimeBest.score < best.score) {
				allTimeBest = best;
			}

			outputFile = WASM_generate_output(JSON.stringify(allTimeBest.simulation));
			output = `Cycles: ${allTimeBest.cycle}\nScore: ${allTimeBest.score}\n${JSON.stringify(
				allTimeBest.simulation.stock,
				null,
				'\t'
			)}`;

			new_generation();
		}, 1);
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

	// -------- Example
	function handle_select(e: any) {
		const index = e.target.value as number;
		if (index == 0) {
			input = customInput;
		} else if (index > 0 && index <= examples.length) {
			input = examples[index - 1].text;
		}
		save_input_state(input, index);
		lastError = WASM_parse_input(input);
	}

	// -------- State
	function handle_run() {
		if (!running) {
			running = true;
			stop = false;
			stopped = false;

			const raw_running_solver = WASM_initialize(
				JSON.stringify({
					text: input,
					generations: $ArgumentStore.generations,
					deep: $ArgumentStore.deep,
					population: $ArgumentStore.population
				})
			);

			if (raw_running_solver == undefined || raw_running_solver == null) {
				output = 'error';
			} else {
				generation = -1;
				const running_solver = parse_as<RunningSolver>(raw_running_solver);
				result_wasm_json = { running_solver, scored_population: { instances: [] } };
				new_generation();
			}
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
			generation = 0;
			finished = false;
			allTimeBest = null;

			GenerationStore.reset();
			frame++;
			output = '';
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

	// ------------------------------ cookie
	// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/cookies
	// https://developer.mozilla.org/en-US/docs/Glossary/Base64

	const COOKIE_KEY_INPUT = 'input';
	const COOKIE_KEY_SELECT = 'select';

	function save_input_state(input: string, select: number) {
		if (browser) {
			// encode
			const input_64 = btoa(input);

			// save in cookies
			document.cookie = `${COOKIE_KEY_INPUT}=${input_64}; path=/`;
			document.cookie = `${COOKIE_KEY_SELECT}=${select}; path=/`;
		}
	}

	function handle_input(e: any) {
		selectedExample = 0;
		customInput = e.target.value;
		save_input_state(e.target.value, 0);
		lastError = WASM_parse_input(e.target.value);
	}

	onMount(async () => {
		if (browser) {
			// extract input from cookies
			const input_64 = document.cookie
				.match('(^|;)\\s*' + COOKIE_KEY_INPUT + '\\s*=\\s*([^;]+)')
				?.pop();

			if (input_64 != undefined) {
				// decode
				const input_text = atob(input_64);

				input = input_text;
				customInput = input_text;
			}

			// extract select value
			const rawSelect = document.cookie
				.match('(^|;)\\s*' + COOKIE_KEY_SELECT + '\\s*=\\s*([^;]+)')
				?.pop();

			if (rawSelect != undefined) {
				const select = Number(rawSelect);
				if (!isNaN(select) && select >= 0 && select < examples.length) {
					selectedExample = select;
				}
			}

			// huh
			setTimeout(() => {
				lastError = WASM_parse_input(input);
			}, 50);
		}
	});
</script>

<!-- ---------------------------------------------- CONTENT -->
<main>
	<header>
		<h1>GPGM</h1>
		<p class="opacity-30">genetic process graph manager</p>
	</header>

	<div class="text-container">
		<h2>Input</h2>
		<div class="text-left">
			<select bind:value={selectedExample} name="examples" id="examples" on:input={handle_select}>
				<option value={0}>Custom</option>
				{#each examples as example, index}
					<option value={index + 1}>{example.name}</option>
				{/each}
			</select>
		</div>
		<div class="relative mt-4">
			<textarea
				cols={Config.io.input.cols}
				rows={Config.io.input.row}
				placeholder=""
				bind:value={input}
				autocorrect="off"
				autocapitalize="off"
				spellcheck="false"
				on:input={handle_input}
			/>
			<img src="/mascot.png" alt="" class="absolute -translate-y-[44%]" />
		</div>
		{#if lastError}
			<div class="error mt-4">
				{lastError}
			</div>
		{/if}
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
		{#if running}
			<button class="side-button" on:click={handle_bottom}> Bottom </button>
		{:else}
			<button class="play-button" on:click={handle_run}> Run </button>
		{/if}
		<button class="play-button" on:click={handle_run} disabled={!input.length || running}>
			Clear
		</button>
	</div>
	{#if output}
		<Visual {frame} />
		<div class="statistic-container shadow">
			<p class="statistic">
				<span class="statistic-label">generation</span>:
				<span class="statistic-value">{generation}</span>
			</p>
			<p class="statistic">
				<span class="statistic-label">best score</span>:
				<span class="statistic-value">{$StatisticStore.scores.global.best}</span>
			</p>
		</div>
		<div class="state-container">
			<button class="side-button" on:click={handle_top} disabled={running && !stopped}>Top</button>
			{#if !running}
				<button class="play-button" on:click={handle_run}>Run</button>
			{:else if !stopped}
				<button class="play-button" on:click={handle_stop}>Stop</button>
			{:else}
				<button class="play-button" on:click={handle_continue} disabled={finished}>Continue</button>
			{/if}
			<button class="side-button" on:click={handle_reset} disabled={disabled_reset}>Reset</button>
		</div>
		<div transition:scale|local class="text-container">
			<h2>Output</h2>
			<textarea
				cols={Config.io.output.cols}
				rows={Config.io.output.row}
				placeholder=""
				value={output}
				readonly
			/>
			<textarea
				cols={Config.io.output.cols}
				rows={Config.io.output.row}
				placeholder=""
				value={outputFile}
				readonly
			/>
		</div>
	{/if}
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

	/* ----------------------- Textarea */
	.statistic-container {
		@apply flex m-auto w-fit py-2;
	}

	.statistic > span {
		@apply inline-block;
	}

	.statistic-label {
		@apply text-right w-24;
	}

	.statistic-value {
		@apply text-left w-16;
	}
</style>
