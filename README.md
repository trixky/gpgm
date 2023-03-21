# gpgm

GPGM (genetic process graph manager) is a homemade algorithm that find the best sequence of process execution to optimize focused resources production using pathfinding [graph traversal](https://en.wikipedia.org/wiki/Graph_traversal) and [genetic algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm).

<img src="https://github.com/trixky/gpgm/blob/main/.demo/screenshots.gif" alt="Demo gif" width="600"/>

## Usage

The program take in input a text that describe a scenario with tree parts:

- The __starting resources__ with which the program begins.
- The __processes__ the program can execute. (take consume and product resources with a delay)
- The __resources to optimize__.

```bash
# starting resources
name:quantity

# processes
name:(need1:qty1;need2:qty2;[...]):(result1:qty1;result2:qty2;[...]):delay

# resources to optimize
# time option can be added to get maximum results in minimum time (can block the production flow)
optimize:(time|stock1;time|stock2;...)
```

## Algorithm

The algorithm start by explore the graph in the graph traversal part and save the dependencies of each process.

Next the genetic part start, using a population that evolves over the generations by [mutation](https://en.wikipedia.org/wiki/Mutation_(genetic_algorithm)) and [crossing](https://en.wikipedia.org/wiki/Crossover_(genetic_algorithm)).

Each instance of the population is represented by a chromosome, made up of genes.

Each gene represents a process, with its dependencies to call or not in a certain order, according to previous processes executed using a history system.

## Prerequisites

- docker-compose

## Up

### Dev

```bash
export PORT=2178
docker compose up -d
# add --build argument if a package is added or environment change
# localhost:2178
```

### Production

```bash
export PORT=2178
docker compose -f docker-compose.prod.yml up -d
# add --build argument if a package is added or environment change
# localhost:2178
```

## Online

This project is [online](https://gpgm.trixky.com/)!

## Stack

- Svelte
- Go (wasm)
- Tailwind
- Chart.js
