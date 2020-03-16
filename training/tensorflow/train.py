import json
import random
import sys

import click
import numpy as np

from tensorflow import keras
import tensorflow as tf


def trainModel(config, weights, x_train, y_train):
    model = keras.models.model_from_json(json.dumps(config))
    model.set_weights(weights)
    model.compile(optimizer='adam',
                  loss='sparse_categorical_crossentropy',
                  metrics=['accuracy'])
    model.fit(x_train, y_train, epochs=1)
    weights = model.get_weights()

    return weights


PARTITIONS = 60


def getData():
    mnist = tf.keras.datasets.mnist
    (x_train, y_train), (_, _) = mnist.load_data()
    partition = random.randint(0, PARTITIONS)
    x_train = [x for x in np.split(x_train, PARTITIONS)][partition]
    y_train = [y for y in np.split(y_train, PARTITIONS)][partition]
    x_train = x_train / 255.0
    return x_train, y_train

@click.command()
@click.argument('c', type=click.STRING)
@click.argument('w', type=click.STRING)
@click.argument('o', type=click.STRING)
def run_standalone(c, w, o):
    config = None

    with open(c, "r") as config_file:
        config = config_file.read()
        config = json.loads(config)

    if config is None:
        print("Config could not be loaded, abort.", file=sys.stderr)
        sys.exit(1)
    else:
        print("Configuration loaded")

    weights = None
    with open(w, "r") as weights_file:
        weights = weights_file.read()
        weights = json.loads(weights)
        weights = [np.array(w) for w in weights]

    if weights is None:
        print("Weights could not be loaded, abort.", file=sys.stderr)
        sys.exit(1)
    else:
        print("Weights loaded")

    x_train, y_train = getData()
    if x_train is None or y_train is None:
        print("Data could not be loaded, abort.", file=sys.stderr)
        sys.exit(1)
    else:
        print("Data loaded")

    print("Starting training")
    new_weights = trainModel(config, weights, x_train, y_train)
    print("Finished training")

    new_weights = [w.tolist() for w in new_weights]
    new_weights = json.dumps(new_weights)

    with open(o, "w") as output_file:
        output_file.write(new_weights)
        print("Saved updated weights to", o)

    #sys.exit(1)

if __name__ == '__main__':
    run_standalone()
