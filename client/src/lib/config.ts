export default {
    io: {
        input: {
            cols: 84,
            row: 30,
        },
        generations: {
            min: 1,
            max: 10000,
            default: 1
        },
        population: {
            min: 1,
            max: 1000,
            default: 1
        },
        deep: {
            min: 10,
            max: 1000,
            default: 100
        },
        delay: {
            min: 500,
            max: 60000,
            default: 5000
        },
        output: {
            cols: 42,
            row: 10,
        }
    }
}