<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { browser, dev } from '$app/environment';
	import { globalReady } from '$lib/stores/ready';
	import '../app.css';

	export let data: { bytes: BufferSource } | undefined;

	if (!dev && browser && 'serviceWorker' in navigator) {
		navigator.serviceWorker
			.register('/service-worker.js', {
				scope: '/'
			})
			.then((registration) => {
				let serviceWorker;
				if (registration.active) {
					serviceWorker = registration.active;
				}
				if (serviceWorker && serviceWorker.state === 'activated') {
					$globalReady = true;
				}
			});
	}
</script>

<!-- ---------------------------------------------- CONTENT -->
<svelte:head>
	<title>GPGM</title>
</svelte:head>

<main>
	<slot {data} />
</main>

<!-- ---------------------------------------------- STYLE -->
