<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { onMount } from 'svelte';
	import type { RunningSolver, WASMGenerationReturn } from '../types';
	import type { ScoredInstance } from '../types/population';
	import InstanceStore from '$lib/stores/instance';
	import Chart from '$lib/components/chart.svelte';
	import type GenerationModel from '$lib/models/generation';
	import { config } from '$lib/config';
	import args from '$lib/stores/arguments';
	import examples from '$lib/Examples';
	import { scale } from 'svelte/transition';
	import { parse_as } from '$lib/utils/parse';
	import { wasmReady } from '$lib/stores/ready';
	import { inputs } from '$lib/stores/inputs';

	export let data: { bytes: BufferSource };

	let start: number = -1;

	// ------------------------------ IO
	let output: ScoredInstance | null = null;
	let outputError = '';
	let outputFile = '';
	let lastError: string | null = null;

	// ------------------------------ State
	let running = false;
	let stop = false;
	let stopped = false;
	let finished = false;

	$: disabled_reset = !running || !stopped;

	// ------------------------------ chrono
	let chrono = 0;
	let stop_chrono = false;

	function remaining_chrono(): number {
		return start + $args.time_limit - new Date().getTime();
	}

	function recursive_chrono() {
		if (!stop_chrono) {
			setTimeout(() => {
				chrono = remaining_chrono();
				recursive_chrono();
			}, 1);
		} else {
			stop_chrono = false;
		}
	}

	function start_chrono() {
		recursive_chrono();
	}

	function finish_chrono() {
		stop_chrono = true;
	}

	// ------------------------------ Loop
	let generation = 0;
	let result_wasm_json: WASMGenerationReturn | undefined = undefined;

	function new_generation() {
		generation++;
		if (generation >= $args.max_generations) {
			handle_finish();
		}

		if (stop || stopped) {
			stop = false;
			stopped = true;
			return;
		}

		setTimeout(() => {
			// Recursive loop

			const result_wasm = WASM_run_generation(JSON.stringify(result_wasm_json!.running_solver));
			result_wasm_json = parse_as<WASMGenerationReturn>(result_wasm);

			const best = result_wasm_json.scored_population.instances[0];
			outputFile = WASM_generate_output(JSON.stringify(best.simulation));
			output = best;
			if (result_wasm_json != undefined) {
				const remaining = remaining_chrono();
				result_wasm_json.running_solver.time_limit_ms = remaining;
				InstanceStore.insert_population(<GenerationModel>{
					scores: result_wasm_json.scored_population.instances.map((instance) => instance.score)
				});
				if (remaining > 0) {
					new_generation();
				} else {
					handle_finish();
				}
			}
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
	function handle_finish() {
		finished = true;
		stopped = true;
		finish_chrono();
		handle_bottom();
	}

	function handle_run() {
		lastError = WASM_parse_input($inputs.current);

		if (!lastError && !running) {
			running = true;
			stop = false;
			stopped = false;
			handle_bottom();
			InstanceStore.reset();

			start = new Date().getTime();

			const raw_running_solver = WASM_initialize(
				JSON.stringify({
					...$args,
					text: $inputs.current
				})
			);

			if (raw_running_solver == undefined || raw_running_solver == null) {
				outputError = 'error';
			} else {
				generation = -1;
				const running_solver = parse_as<RunningSolver>(raw_running_solver);
				result_wasm_json = { running_solver, scored_population: { instances: [] } };
				start_chrono();
				InstanceStore.insert_population(<GenerationModel>{
					scores: result_wasm_json.scored_population.instances.map((instance) => instance.score)
				});
				new_generation();
			}
		}
	}

	function handle_reset() {
		if (running && stopped) {
			InstanceStore.reset();
			running = false;
			stop = false;
			stopped = false;
			generation = 0;
			finished = false;

			output = null;
			outputError = '';
		}
	}

	function handle_input(e: any) {
		$inputs.selectedExample = 0;
		$inputs.custom = e.target.value;
		if (lastError) {
			lastError = WASM_parse_input(e.target.value);
		}
	}

	function handle_input_change(e: any) {
		lastError = WASM_parse_input(e.target.value);
	}

	function download_output(e: Event) {
		e.preventDefault();
		if (outputFile) {
			const file = new Blob([outputFile], { type: 'plain/text' });

			// Create a new link
			const url = URL.createObjectURL(file);
			const anchor = document.createElement('a');
			anchor.href = url;
			anchor.download = `${examples[$inputs.selectedExample].name.toLocaleLowerCase()}.out`;

			// Append to the DOM
			document.body.appendChild(anchor);

			// Trigger `click` event
			anchor.click();

			// Remove element from DOM
			document.body.removeChild(anchor);
			URL.revokeObjectURL(url);
		}
	}

	// ------------------------------ Mascot
	const mascot_random_duration_secondes = 5
	const mascot_minimum_duration_secondes = 5
	const mascot_minimum_x_deplacement = 30
	const mascot_maximum_x_deplacement = 80
	const mascot_maximum_x = 140
	const second_ratio = 1000

	let mascot_x = 0
	let mascot_reverse = false
	let mascot_pause = false

	let information_readed = false

	function move_mascot() {
		setTimeout(() => {
			if (!mascot_pause) {
				while(true) {
					const new_mascot_x = Math.ceil(Math.random() * mascot_maximum_x)
					
					const mascot_deplacement = Math.abs(new_mascot_x - mascot_x)

					if (mascot_deplacement > mascot_minimum_x_deplacement && mascot_deplacement < mascot_maximum_x_deplacement) {
						mascot_reverse = new_mascot_x > mascot_x
						mascot_x = new_mascot_x
						break
					}
				}
			}
	
			move_mascot()
		}, ((Math.random() * mascot_random_duration_secondes) + mascot_minimum_duration_secondes) * second_ratio);
	}

	function mascot_mouse_in() {
		mascot_pause = true
	}
	
	function mascot_mouse_out() {
		mascot_pause = false
	}

	function information_mouse_in() {
		information_readed = true
	}

	onMount(() => {
		// @ts-expect-error
		// Go is loaded from the app.html (wasm)
		const goWasm = new Go();

		WebAssembly.instantiate(data.bytes, goWasm.importObject).then((result) => {
			goWasm.run(result.instance);
			$wasmReady = true;
		});

		move_mascot()
	});

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

	<div class="block-top">
		<div class="text-container">
			<h2>Input</h2>
			<div class="text-left relative">
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
			<div class="relative mt-4 w-full">
				<textarea
					placeholder=""
					class="scrollbar-custom"
					bind:value={$inputs.current}
					autocorrect="off"
					autocapitalize="off"
					spellcheck="false"
					on:input={handle_input}
					on:change={handle_input_change}
				/>
				<div class="mascot-container" style="transform: translateX(-{mascot_x}px)" on:mouseenter={mascot_mouse_in} on:mouseleave={mascot_mouse_out}>
					<img src="/mascot.png" alt="" class="mascot" class:reverse={mascot_reverse} title="GPGM mascot engineer"/>
					<img src="/information.svg" alt="" class="information" class:animate-pulse={!information_readed} on:mouseenter={information_mouse_in} title="GPGM is a solution that find the best sequence of process execution&#13to optimize focused resources production using pathfinding graph and genetic algorithms"/>
				</div>
			</div>
			{#if lastError}
				<div class="error mt-4">
					{lastError}
				</div>
			{/if}
		</div>
		<form on:submit|preventDefault={handle_run}>
			<div class="form-container">
				<div class="input-container" title="Number of generation to execute">
					<input
						type="number"
						min={config.io.max_generations.min}
						max={config.io.max_generations.max}
						bind:value={$args.max_generations}
						disabled={running}
					/>
					<p class="input-label">gen</p>
				</div>
				<div class="input-container" title="Population size of each generation">
					<input
						type="number"
						min={config.io.population_size.min}
						max={config.io.population_size.max}
						bind:value={$args.population_size}
						disabled={running}
					/>
					<p class="input-label">pop</p>
				</div>
				<div class="input-container" title="Maximum number of cycle">
					<input
						type="number"
						min={config.io.max_cycle.min}
						max={config.io.max_cycle.max}
						bind:value={$args.max_cycle}
						disabled={running}
					/>
					<p class="input-label">cyc</p>
				</div>
				<div class="input-container" title="Time out in millisecond">
					<input
						type="number"
						min={config.io.time_limit.min}
						max={config.io.time_limit.max}
						bind:value={$args.time_limit}
						disabled={running}
					/>
					<p class="input-label">ms</p>
				</div>
				<div class="input-container" title="Maximum depth of the graph explorer">
					<input
						type="number"
						min={config.io.max_depth.min}
						max={config.io.max_depth.max}
						bind:value={$args.max_depth}
						disabled={running}
					/>
					<p class="input-label">dep</p>
				</div>
				<div class="input-container" title="Number of preserved instances by elitism">
					<input
						type="number"
						min={config.io.elitism_amount.min}
						max={config.io.elitism_amount.max}
						bind:value={$args.elitism_amount}
						disabled={running}
					/>
					<p class="input-label">eli</p>
				</div>
				<div class="input-container" title="Number maximum of entry point/process by instance&#13(Set to 0 to disable)">
					<input
						type="number"
						min={config.io.max_cut.min}
						max={config.io.max_cut.max}
						bind:value={$args.max_cut}
						disabled={running}
					/>
					<p class="input-label">cut</p>
				</div>
				<div class="input-container" title="Cross-over strategy (genetic)">
					<select
						name="selection_method"
						id="selection_method"
						bind:value={$args.selection_method}
						disabled={running}
					>
						{#each config.io.selection_method.choices as choice}
							<option value={choice.value}>{choice.label}</option>
						{/each}
					</select>
				</div>
				<div class="input-container" title="Tournament size ???">
					<input
						type="number"
						min={config.io.tournament_size.min}
						max={config.io.tournament_size.max}
						bind:value={$args.tournament_size}
						disabled={running}
					/>
					<p class="input-label">tor</p>
				</div>
				<div class="input-container" title="Tournament probability ???">
					<input
						type="number"
						min={config.io.tournament_probability.min}
						max={config.io.tournament_probability.max}
						step="0.01"
						bind:value={$args.tournament_probability}
						disabled={running}
					/>
					<p class="input-label">pro</p>
				</div>
				<div class="input-container" title="Cross-over ???">
					<input
						type="number"
						min={config.io.crossover_new_instances.min}
						max={config.io.crossover_new_instances.max}
						bind:value={$args.crossover_new_instances}
						disabled={running}
					/>
					<p class="input-label">cro</p>
				</div>
				<div class="input-container" title="Mutation strategy&#13(The mutation rate decreases over time)">
					<select
						name="mutation_method"
						id="mutation_method"
						bind:value={$args.mutation_method}
						disabled={running}
					>
						{#each config.io.mutation_method.choices as choice}
							<option value={choice.value}>{choice.label}</option>
						{/each}
					</select>
				</div>
			</div>
			<div class="state-container">
				{#if running}
					<button class="side-button" on:click={handle_reset} disabled={disabled_reset}
						>Reset</button
					>
				{:else}
					<button class="play-button" disabled={lastError !== null} on:click={handle_run}
						>Run</button
					>
				{/if}
			</div>
		</form>
	</div>
	{#if output}
		<div class="block-bottom">
			<div transition:scale|local class="text-container">
				<h2>Output</h2>
				<div class="statistic-container shadow">
					<p class="statistic">
						<span class="statistic-label">generation</span>:
						<span class="statistic-value">{generation}/{$args.max_generations}</span>
					</p>
					<p class="statistic">
						<span class="statistic-value chrono">{chrono} ms</span>
					</p>
				</div>
				<div class="flex flex-row w-full mb-1 mt-3">
					{#if output}
						<span class="flex-shrink-0 inline-block w-40 text-left">
							Score: {output.score}
						</span>
						<span class="flex-shrink-0 inline-block w-40 text-left">
							Cycles: {output.cycle}
						</span>
					{/if}
					{#if outputError}
						{outputError}
					{/if}
				</div>
				<div class="text-left max-w-lg">
					<div>
						{#each Object.keys(output.simulation.stock) as product}
							<span
								class="badge mr-1 mb-1"
								class:highlight={output.simulation.initial_context.optimize[product] !== undefined}
							>
								{product}: {output.simulation.stock[product]}
							</span>
						{/each}
					</div>
				</div>
				<Chart />
				<div class="flex flex-col mt-3">
					<textarea class="scrollbar-custom" placeholder="" value={outputFile} readonly />
					<button class="download" on:click={download_output}>Download</button>
				</div>
			</div>
		</div>
	{/if}
</main>

<!-- <svelte:window use:wheel={{ scrollable }} /> -->

<!-- ---------------------------------------------- STYLE -->
<style lang="postcss">
	/* ----------------------- Input/Output block */
	.block-top,
	.block-bottom {
		@apply m-auto px-3 mt-6;
		width: 520px;
	}

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

	/* ----------------------- Mascot / Information */

	.mascot-container {
		@apply absolute -top-[0px] right-0 w-[100px] h-[1px] transition-all duration-[2000ms] -z-10;
		transition-timing-function: linear;
	}

	.mascot {
		@apply absolute -top-[86px] right-0;
	}

	.mascot.reverse {
		transform: scaleX(-1)
	}

	.information {
		@apply absolute w-6 -top-[118px] opacity-70 cursor-help;
	}

	/* ----------------------- Buttons */
	.state-container {
		@apply my-3;
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
		@apply relative w-fit flex m-auto flex-wrap justify-between;
	}

	.input-container {
		@apply relative;
		padding: -10px;
		margin: -4px;
	}

	.input-container > select {
		@apply mx-1 !important;
	}

	input,
	select {
		@apply mt-5 px-3 py-1 w-[116px];
	}

	select {
		@apply pl-1;
	}
	option {
		@apply py-1;
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
		@apply relative m-auto w-full;
	}

	.text-container > h2 {
		@apply mt-4 text-left;
	}

	textarea {
		@apply relative p-3 z-10 w-full h-56;
	}

	/* ----------------------- Textarea */
	.statistic-container {
		@apply flex w-fit py-2 mt-5 mb-4;
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

	.statistic-value.chrono {
		@apply w-24 mr-4 text-right;
	}

	.download {
		@apply ml-auto mr-0 mt-4;
	}

	@media screen and (max-width: 580px) {
		.block-top,
		.block-bottom {
			width: 420px;
		}
	}


	@media screen and (max-width: 440px) {
		.form-container {
			@apply justify-center;
		}

		.block-top,
		.block-bottom {
			width: 90%;
		}

		.input-container {
			padding: 0;
			margin: 0;
		}
	}
</style>
