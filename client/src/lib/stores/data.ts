import InstanceStore from './instance'
import { derived } from 'svelte/store';

export default derived(
	InstanceStore,
	$InstanceStore => $InstanceStore.map((instance, index) => {
        const color = index === 0 ? 'rgb(220, 252, 231, 1)' : 'rgb(255, 255, 255, 0.7)'
        return{
            // label: '',
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
            data: instance.scores
    }})
);