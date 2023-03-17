<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { browser, dev } from '$app/environment';
	import { globalReady } from '$lib/stores/globalReady';
	import { useWorker } from '$lib/stores/useWorker';
	import { workerReady } from '$lib/stores/workerReady';
	import '../app.css';

	export let data: { bytes: BufferSource };

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
					if (navigator.serviceWorker.controller) {
						$workerReady = true;
						$globalReady = true;
					}
				}
			})
			.catch(() => {
				$useWorker = false;
			});
	} else {
		$useWorker = false;
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
