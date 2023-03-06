import { writable } from 'svelte/store';
import type StatisticModel from '../models/statistic';
import type { Scores } from '../models/statistic';

function generate_default_statistic(): StatisticModel {
    return <StatisticModel>{
        started: false,
        scores: {
            referentiel: {
                offset: 0,
                diff: 0
            },
            current: {
                worst: 0,
                best: 0,
            },
            global: {
                worst: 0,
                best: 0,
            }
        }
    }
}

function create_statistic_store() {
    const { subscribe, update, set } = writable(generate_default_statistic());

    return {
        subscribe,
        reset: () => {
            set(generate_default_statistic())
        },
        // --------------------- insert
        set_insert_score: (scores: Scores) => {
            update(statistic => {
                if (!statistic.started) {
                    statistic.started = true
                    // update referentiel scores
                    statistic.scores.referentiel.offset = scores.worst
                    statistic.scores.referentiel.diff = scores.best - scores.worst
                }

                // update current scores
                statistic.scores.current = scores

                // update global scores
                if (statistic.scores.current.best > statistic.scores.global.best)
                    statistic.scores.global.best = statistic.scores.current.best
                // update current scores
                if (statistic.scores.current.worst > statistic.scores.global.worst)
                    statistic.scores.global.worst = statistic.scores.current.worst

                return statistic
            })
        },
    };
}

const statistic_store = create_statistic_store();

export default statistic_store;