<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import '../app.css';

	export let data: { bytes: BufferSource };
</script>

<!-- ---------------------------------------------- CONTENT -->
<svelte:head>
	<title>GPGM</title>
</svelte:head>

<main>
	<slot {data} />
</main>

<!-- ---------------------------------------------- STYLE -->
