<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import { Line } from 'svelte-chartjs';
	import {
		Chart as ChartJS,
		LineElement,
		Title,
		Tooltip,
		Legend,
		ArcElement,
		CategoryScale,
		LinearScale,
		PointElement,
		type ChartData,
		type Point
	} from 'chart.js';
	import DataStore from '$lib/stores/data';
	import LabelStore from '$lib/stores/label';

	ChartJS.register(
		LineElement,
		Title,
		Tooltip,
		Legend,
		ArcElement,
		CategoryScale,
		LinearScale,
		PointElement
	);

	let dataLine: ChartData<'line', (number | Point)[]> = {
		labels: [],
		datasets: $DataStore
	};

	$: dataLine = {
		labels: $LabelStore,
		datasets: $DataStore
	};

	const options = {
		events: [], // disable mouse hover events
		plugins: {
			legend: {
				display: false, // remove line legends
				labels: {
					color: 'white'
				}
			}
		},
		responsive: true,
		scales: {
			y: {
				title: {
					display: false,
					text: 'score',
					color: 'white'
				},
				grid: {
					color: 'black'
				},
				ticks: {
					color: 'white'
				}
			},
			x: {
				title: {
					display: false,
					text: 'generation',
					color: 'white'
				},
				grid: {
					color: 'black'
				},
				ticks: {
					color: 'white'
				}
			}
		}
	};
</script>

<!-- ---------------------------------------------- CONTENT -->
<div class="chart-container">
	<Line data={dataLine} {options} />
</div>

<!-- ---------------------------------------------- STYLE -->
<style lang="postcss">
	.chart-container {
		@apply w-[500px] pl-3 pr-8 m-auto my-6;
	}

	@media screen and (max-width: 580px) {
		.chart-container {
			@apply w-[400px];
		}
	}

	@media screen and (max-width: 440px) {
		.chart-container {
			@apply w-[320px] pl-2 pr-5;
		}
	}

	@media screen and (max-width: 390px) {
		.chart-container {
			@apply w-[240px] pl-0 pr-3;
		}
	}

	@media screen and (max-width: 280px) {
		.chart-container {
			@apply hidden;
		}
	}
</style>
