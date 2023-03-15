<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import type { RunningSolver, WASMGenerationReturn } from '../types';
	import Config from '$lib/config';
	import args from '$lib/stores/arguments';
	import examples from '$lib/Examples';
	import { scale } from 'svelte/transition';
	import { parse_as } from '$lib/utils/parse';
	import { wasmReady } from '$lib/stores/ready';
	import { inputs } from '$lib/stores/inputs';

	// ------------------------------ IO
	let output = '';
	let outputFile = '';
	let lastError: string | null = null;

	// ------------------------------ State
	let running = false;
	let stop = false;
	let stopped = false;
	let finished = false;

	$: disabled_reset = !running || !stopped;

	// ------------------------------ Loop
	let frame = 0; // Used to refresh the visualizator
	let generation = 0;
	let result_wasm_json: WASMGenerationReturn | undefined = undefined;

	function new_generation() {
		generation++;
		if (generation >= $args.max_generations) {
			finished = true;
			stopped = true;
		}

		if (stop || stopped) {
			stop = false;
			stopped = true;
			return;
		}

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
			outputFile = WASM_generate_output(JSON.stringify(best.simulation));
			output = `Cycles: ${best.cycle}\nScore: ${best.score}\n${JSON.stringify(
				best.simulation.stock,
				null,
				'\t'
			)}`;

			new_generation();
			handle_bottom();
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
			$inputs.current = $inputs.custom;
		} else if (index > 0 && index <= examples.length) {
			$inputs.current = examples[index - 1].text;
		}
		lastError = WASM_parse_input($inputs.current);
		handle_reset();
	}

	// -------- State
	function handle_run() {
		lastError = WASM_parse_input($inputs.current);

		if (!lastError && !running) {
			running = true;
			stop = false;
			stopped = false;
			handle_bottom();

			const raw_running_solver = WASM_initialize(
				JSON.stringify({
					...$args,
					text: $inputs.current
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

			frame++;
			output = '';
		}
	}

	// ------------------------------ Scrolling blocker
	// https://svelte.dev/repl/2bdbf66371a3418e9e3eda076df6e32d?version=3.18.1
	/* $: scrollable = !running || stopped;

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
	}; */

	// ------------------------------ cookie
	// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/cookies
	// https://developer.mozilla.org/en-US/docs/Glossary/Base64

	function handle_input(e: any) {
		$inputs.selectedExample = 0;
		$inputs.custom = e.target.value;
		lastError = WASM_parse_input(e.target.value);
	}

	wasmReady.subscribe((ready) => {
		if (ready) {
			lastError = WASM_parse_input($inputs.current);
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
			<select
				bind:value={$inputs.selectedExample}
				name="examples"
				id="examples"
				on:input={handle_select}
			>
				<option value={0}>Custom</option>
				{#each examples as example, index}
					<option value={index + 1}>{example.name}</option>
				{/each}
			</select>
		</div>
		<div class="relative mt-4">
			<textarea
				cols={Config.ui.input.cols}
				rows={Config.ui.input.row}
				placeholder=""
				bind:value={$inputs.current}
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
				min={Config.io.max_generations.min}
				max={Config.io.max_generations.max}
				value={$args.max_generations}
				disabled={running}
			/>
			<p class="input-label">gen</p>
		</div>
		<div class="input-container">
			<input
				type="number"
				min={Config.io.population_size.min}
				max={Config.io.population_size.max}
				value={$args.population_size}
				disabled={running}
			/>
			<p class="input-label">pop</p>
		</div>
		<div class="input-container">
			<input
				type="number"
				min={Config.io.max_cycle.min}
				max={Config.io.max_cycle.max}
				value={$args.max_cycle}
				disabled={running}
			/>
			<p class="input-label">cyc</p>
		</div>
		<div class="input-container">
			<input
				type="number"
				min={Config.io.time_limit.min}
				max={Config.io.time_limit.max}
				value={$args.time_limit}
				disabled={running}
			/>
			<p class="input-label">ms</p>
		</div>
	</div>
	<div class="state-container">
		{#if running}
			<button class="side-button" on:click={handle_bottom}> Bottom </button>
		{:else}
			<button class="play-button" disabled={lastError !== null} on:click={handle_run}> Run </button>
		{/if}
		<button class="play-button" on:click={handle_run} disabled={!$inputs.current.length || running}>
			Clear
		</button>
	</div>
	{#if output}
		<!-- <Visual {frame} /> -->
		<div class="statistic-container shadow">
			<p class="statistic">
				<span class="statistic-label">generation</span>:
				<span class="statistic-value">{generation}/{$args.max_generations}</span>
			</p>
			<!-- <p class="statistic">
				<span class="statistic-label">best score</span>:
				<span class="statistic-value">{$StatisticStore.scores.global.best}</span>
			</p> -->
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
				cols={Config.ui.output.cols}
				rows={Config.ui.output.row}
				placeholder=""
				value={output}
				readonly
			/>
			<textarea
				cols={Config.ui.output.cols}
				rows={Config.ui.output.row}
				placeholder=""
				value={outputFile}
				readonly
			/>
		</div>
	{/if}
</main>

<!-- <svelte:window use:wheel={{ scrollable }} /> -->

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
		@apply flex m-auto w-fit py-2 mt-4;
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
