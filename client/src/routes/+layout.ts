import type { PageLoad } from './$types';

export const prerender = true

export const load = (async ({ fetch }) => {
	const res = await fetch('wasm/src/main.wasm');
	const bytes = await res.arrayBuffer();

	return { bytes };
}) satisfies PageLoad;