import click
import sys

from util import loadModelFromDisk, evaluateModel


@click.command()
@click.argument('c', type=click.STRING)
@click.argument('w', type=click.STRING)
@click.argument('o', type=click.STRING)
def run_standalone(c, w, o):

    config, weights = loadModelFromDisk(c, w)

    score = evaluateModel(config, weights)

    try:
        with open(o, "w") as file:
            file.write(str(score))
    except IOError as e:
        print("Could not write to", o, "Error:", e)
        sys.exit(1)
    print("Saved score to", o)


if __name__ == '__main__':
    run_standalone()
