export interface Scores {
    worst: number,
    best: number,
}

export interface Referentiel {
    offset: number,
    diff: number
}

export default interface Statistic {
    started: boolean,
    scores: {
        referentiel: Referentiel,
        global: Scores,
        current: Scores
    }
}
