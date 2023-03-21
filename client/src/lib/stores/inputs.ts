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

const inputs = writable<Inputs>(JSON.parse(browser ? localStorage.getItem('inputs') ?? defaultValue() : defaultValue()))

export default inputs

inputs.subscribe((value) => browser ? localStorage.setItem('inputs', JSON.stringify(value)) : undefined)
