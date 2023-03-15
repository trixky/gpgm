import { browser } from '$app/environment'
import { writable } from 'svelte/store'

interface Inputs {
	current: string,
	custom: string,
	selectedExample: number
}

const defaultValue = () => JSON.stringify({
	current: '',
	custom: '',
	selectedExample: 0,
})

export const inputs = writable<Inputs>(JSON.parse(browser ? localStorage.getItem('inputs') ?? defaultValue() : defaultValue()))

inputs.subscribe((value) => browser ? localStorage.setItem('inputs', JSON.stringify(value)) : undefined)
