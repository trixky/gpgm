<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { onMount, tick } from 'svelte';
	import GenerationSortedStore from '../../stores/generation_sorted';
	import { get_color_from_percentage } from '../../utils/color';
	import type GenerationModel from '../../models/generation';

	const PIXEL_SIZE = 4;
	const COLOR_R = 0;
	const COLOR_G = 1;
	const COLOR_B = 2;
	const COLOR_X = 3;

	let mounted = false;

	let previous_frame: number;
	export let frame: number;

    let width = 0
    let height = 0
    
	$: width = $GenerationSortedStore.length ? $GenerationSortedStore[0].instances.length : 0;
	$: height = $GenerationSortedStore.length;

	let canvas: any = undefined;
	let ctx: any = undefined;

    function handle_bottom() {
		window.scrollTo({
			top: document.body.scrollHeight,
		});
	}

	const draw = (generations: Array<GenerationModel>) => {
		// https://stackoverflow.com/questions/4899799/whats-the-best-way-to-set-a-single-pixel-in-an-html5-canvas

		if (width) {
			var image = ctx.getImageData(0, 0, width, height);
			var pixels = image.data;

			let off = 0;

			generations.forEach((generation) => {
				generation.instances.forEach((instance) => {
					const rgb = get_color_from_percentage(instance.score);

					pixels[off + COLOR_R] = rgb.red;
					pixels[off + COLOR_G] = rgb.green;
					pixels[off + COLOR_B] = rgb.blue;

					pixels[off + COLOR_X] = 255;

					off += PIXEL_SIZE;
				});
			});

			ctx.putImageData(image, 0, 0);
		}
	};

	onMount(async () => {
		mounted = true;
	});

	$: if (mounted && previous_frame != frame && canvas != undefined) {
		(async () => {
			previous_frame = frame;

			ctx = canvas.getContext('2d');
			canvas.width = width;
			canvas.height = height;
            handle_bottom()
			await tick();
			draw($GenerationSortedStore);
            handle_bottom()
		})();
	}
</script>

<!-- ---------------------------------------------- CONTENT -->
<div class="visual shadow" class:invisible={!width}>
    <canvas bind:this={canvas} />
</div>

<!-- ---------------------------------------------- STYLE -->
<style lang="postcss">
	.visual {
		@apply m-auto w-fit mt-8 mb-3;
		border: solid 1px black;
	}
</style>
