import click

from util import aggregateUpdates, absoluteFilePaths, writeUpdatesToDisk, loadWeightsFromFile


@click.command()
@click.argument('w', type=click.STRING)
@click.argument('o', type=click.STRING)
def run_standalone(w, o):

    updates = []

    for file in absoluteFilePaths(w):

        if not file.endswith("_trainingWeights.in"):
            continue

        weights = loadWeightsFromFile(file)
        updates.append(weights)

    new_weights = aggregateUpdates(updates)

    writeUpdatesToDisk(new_weights, o)


if __name__ == '__main__':
    run_standalone()
