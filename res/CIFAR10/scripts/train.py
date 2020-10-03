import sys
import click

from util import trainModel, writeUpdatesToDisk, loadModelFromDisk, getData


@click.command()
@click.argument('c', type=click.STRING)
@click.argument('w', type=click.STRING)
@click.argument('o', type=click.STRING)
def run_standalone(c, w, o):

    config, weights = loadModelFromDisk(c, w)

    x_train, y_train = getData()

    if x_train is None or y_train is None:
        print("Data could not be loaded, abort.", file=sys.stderr)
        sys.exit(1)
    else:
        print("Data loaded")

    print("Starting training")
    new_weights = trainModel(config, weights, x_train, y_train)
    print("Finished training")

    writeUpdatesToDisk(new_weights, o)


if __name__ == '__main__':
    run_standalone()
