import { dev } from '$app/environment';
import { writable } from 'svelte/store';

export const useWorker = writable(!dev);
