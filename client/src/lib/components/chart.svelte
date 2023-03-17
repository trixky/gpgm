<!-- ---------------------------------------------- SCRIPT -->
<script lang="ts">
	import InstanceStore from '$lib/stores/instance';
	import ArgumentStore from '$lib/stores/arguments';
	import { browser } from '$app/environment';

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
		PointElement
	} from 'chart.js';

	let chart: any = undefined;

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

	let dataLine: any = {
		labels: $ArgumentStore.max_generations,
		datasets: []
	};

	let last_data_length = -1;

	$: if ($InstanceStore.length > 0 && $InstanceStore[0].length > last_data_length && browser) {
		if ($InstanceStore[0].length == 1) {
			dataLine.labels = ['1'];
			dataLine.datasets = $InstanceStore.map((instance, index) => {
				const color = index === 0 ? 'rgb(220, 252, 231, 1)' : 'rgb(255, 255, 255, 0.7)'

				return {
					lineWidth: 40,
					width: 40,
					weight: 40,
					lineTension: 0,
					backgroundColor: color,
					borderColor: color,
					borderCapStyle: 'butt',
					borderDash: [],
					borderDashOffset: 0.0,
					borderJoinStyle: 'miter',
					pointBorderColor: color,
					pointBackgroundColor: color,
					borderWidth: 2,
					pointBorderWidth: 4,
					pointHoverRadius: 5,
					pointHoverBackgroundColor: color,
					pointHoverBorderColor: color,
					pointHoverBorderWidth: 2,
					pointRadius: 1,
					pointHitRadius: 10,
					data: instance.map((score: number) => score)
				};
			});
		} else {
			dataLine.labels.push($InstanceStore[0].length.toString());
			$InstanceStore.forEach((instance, index) => {
				dataLine.datasets[index].data.push(instance[instance.length - 1]);
			});
		}
		chart?.update();
	}

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
	<Line bind:chart data={dataLine} {options} />
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
