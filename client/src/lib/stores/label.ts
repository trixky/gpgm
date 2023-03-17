import InstanceStore from './instance'
import { derived } from 'svelte/store';

export default derived(
	InstanceStore,
	$InstanceStore => $InstanceStore.length > 0 ?
    $InstanceStore[0].scores.map((_, index) => (index + 1).toString())
    : []
);