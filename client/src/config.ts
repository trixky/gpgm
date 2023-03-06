export default {
    io: {
        input: {
            cols: 42,
            row: 10,
        },
        generations: {
            min: 50,
            max: 10000,
            default: 1000
        },
        population: {
            min: 10,
            max: 1000,
            default: 300
        },
        deep: {
            min: 10,
            max: 1000,
            default: 100
        },
        delay: {
            min: 500,
            max: 60000,
            default: 2000
        },
        output: {
            cols: 42,
            row: 10,
        }
    }
}