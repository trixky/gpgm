# gpgm

GPGM (genetic process graph manager) is a homemade algorithm used to solve [RCPSP](https://fr.wikipedia.org/wiki/Probl%C3%A8me_de_gestion_de_projet_%C3%A0_contraintes_de_ressources) (resource-constrained project scheduling problem) using [graph traversal](https://en.wikipedia.org/wiki/Graph_traversal) and [genetic algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm).

<img src="https://github.com/trixky/gpgm/blob/main/.demo/screenshots.gif" alt="Demo gif" width="400"/>

## Usage

The program takes in input a text that describes a scenario with tree parts:

- The __starting resources__ with which the program begins.
- The __processes__ the program can execute. (consume and product resources with a delay)
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

The algorithm starts by the graph traversal part, exploring the dependencies of each process.

Next the genetic part uses a population that evolves over the generations by [mutation](https://en.wikipedia.org/wiki/Mutation_(genetic_algorithm)) and [crossing](https://en.wikipedia.org/wiki/Crossover_(genetic_algorithm)).

Each instance of the population is represented by a chromosome, made up of genes.

Each gene represents a process, with its dependencies to call or not in a certain order, according to previous processes executed using a history system.

## Prerequisites

- docker-compose

## Up

### Dev

```bash
cp .env.example .env
docker compose up -d
# add the --build argument if a package is added or the environment changes
# localhost:7777
```

### Production

```bash
cp .env.example .env
docker compose -f docker-compose.prod.yml up -d
# add the --build argument if a package is added or the environment changes
# localhost:7777
```

## Online

This project is [online](https://gpgm.trixky.com/)!

## Stack

- Svelte
- Go (wasm)
- Tailwind
- Chart.js
