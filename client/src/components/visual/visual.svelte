<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { onMount, tick } from 'svelte';
	import GenerationStore from '../../stores/generation';
	import StatisticStore from '../../stores/statistic';
	import { get_color_from_percentage } from '../../utils/color';
	import type GenerationModel from '../../models/generation';
	import Config from '../../config';
	import type {Referentiel as ReferentielModel} from '../../models/statistic'

	const PIXEL_SIZE = 4;
	const COLOR_R = 0;
	const COLOR_G = 1;
	const COLOR_B = 2;
	const COLOR_X = 3;

	let previous_frame: number;
	export let frame: number;

	let mounted = false;

	let width = 0;
	let height = 0;

	let canvas: any = undefined;
	let ctx: any = undefined;

	function handle_bottom() {
		window.scrollTo({
			top: document.body.scrollHeight
		});
	}

	function normalize_score(score: number, referentiel: ReferentielModel): number {
		return Math.ceil(((score - referentiel.offset) / referentiel.diff) * 100)
	}

	const draw = (generations: Array<GenerationModel>) => {
		// https://stackoverflow.com/questions/4899799/whats-the-best-way-to-set-a-single-pixel-in-an-html5-canvas

		const local_height = $GenerationStore.length;

		if (local_height) {
			const last_generation = generations[generations.length - 1];
			const local_width = last_generation.instances.length;

			var image = ctx.getImageData(0, 0, local_width, local_height);
			var pixels = image.data;

			let pixel = (local_height - 3) * local_width * 4;

			last_generation.instances.forEach((instance) => {
				const normalized_score = normalize_score(instance.score, $StatisticStore.scores.referentiel)

				const rgb = get_color_from_percentage(normalized_score);

				pixels[pixel + COLOR_R] = rgb.red;
				pixels[pixel + COLOR_G] = rgb.green;
				pixels[pixel + COLOR_B] = rgb.blue;

				pixels[pixel + COLOR_X] = 255;

				pixel += PIXEL_SIZE;
			});

			ctx.putImageData(image, 0, 0);
			width = local_width;
			height = generations.length;
		} else {
			// reset
			height = 0;
		}
	};

	onMount(async () => {
		mounted = true;

		ctx = canvas.getContext('2d', {
			willReadFrequently: true // browser optimization
		});
		canvas.height = Config.io.generations.max;
	});

	$: if (mounted && previous_frame != frame && canvas != undefined) {
		(async () => {
			previous_frame = frame;

			ctx = canvas.getContext('2d', {
				willReadFrequently: true
			});

			if (canvas.width != width) {
				canvas.width = width;
			}

			handle_bottom();
			await tick();
			draw($GenerationStore);
			handle_bottom();
		})();
	}
</script>

<!-- ---------------------------------------------- CONTENT -->
<div class="visual shadow" class:invisible={!height} style="width: {width}px; height: {height}px;">
	<canvas bind:this={canvas} />
</div>

<!-- ---------------------------------------------- STYLE -->
<style lang="postcss">
	.visual {
		@apply m-auto w-fit mt-8 mb-5 overflow-hidden;
		border: solid 1px black;
	}
</style>
