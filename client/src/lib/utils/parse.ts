export function parse_as<T extends object>(value: string): T {
	return JSON.parse(value)
}
