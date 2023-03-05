import { derived } from 'svelte/store';
import GenerationStore from './generation'
import type GenerationModel from '../models/generation';

export default derived(
    GenerationStore,
    $generationStore => $generationStore.map(generation => {
        return <GenerationModel>{
            instances: generation.instances.filter(_ => true).sort((a, b) => b.score - a.score)
        }
    })
);