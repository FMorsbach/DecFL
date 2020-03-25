import json
import os
import random
import sys

import numpy as np
import tensorflow as tf

from tensorflow import keras


def trainModel(config, weights, x_train, y_train):
    model = keras.models.model_from_json(json.dumps(config))
    model.set_weights(weights)
    model.compile(optimizer='adam',
                  loss='sparse_categorical_crossentropy',
                  metrics=['accuracy'])
    model.fit(x_train, y_train, epochs=1)
    weights = model.get_weights()

    return weights


def aggregateUpdates(updates):
    weights = [0*w for w in updates[0]]

    for update in updates:
        weights = [w1 + w2 for w1, w2 in zip(weights, update)]

    weights = [w/len(updates) for w in weights]
    weights = [np.array(w) for w in weights]

    return weights


def evaluateModel(config, weights):
    model = keras.models.model_from_json(json.dumps(config))
    model.set_weights(weights)
    model.compile(optimizer='adam',
                  loss='sparse_categorical_crossentropy',
                  metrics=['accuracy'])

    (_, _), (x_test, y_test) = tf.keras.datasets.mnist.load_data()

    evaluation = model.evaluate(x_test,  y_test, verbose=2)
    
    return evaluation


def getRandomSetOfData():
    partitions = 60

    mnist = tf.keras.datasets.mnist
    (x_train, y_train), (_, _) = mnist.load_data()
    partition = random.randint(0, partitions-1)
    x_train = [x for x in np.split(x_train, partitions)][partition]
    y_train = [y for y in np.split(y_train, partitions)][partition]
    x_train = x_train / 255.0
    return x_train, y_train


def loadModelFromDisk(c, w):
    config = loadConfigurationFromFile(c)
    weights = loadWeightsFromFile(w)

    return config, weights


def loadConfigurationFromFile(c):
    try:
        with open(c, "r") as config_file:
            config = config_file.read()
            config = json.loads(config)
    except IOError as e:
        print("Cant open", c, "Error:", e, file=sys.stderr)
        sys.exit(1)
    except json.decoder.JSONDecodeError as e:
        print("Cant parse", c, "to json", "Error:", e, file=sys.stderr)
        sys.exit(1)

    if config is None:
        print("Loaded config from", c, ", is empty", file=sys.stderr)
        sys.exit(1)
    elif not checkConfigurationFormat(config):
        print("Configuration does not match expected format", file=sys.stderr)
        sys.exit(1)
    else:
        print("Configuration loaded")
    return config


def loadWeightsFromFile(file):
    try:
        with open(file, "r") as weights_file:
            weights = weights_file.read()
            weights = json.loads(weights)
            weights = [np.array(w) for w in weights]
    except IOError as e:
        print("Cant open", file, "Error:", e, file=sys.stderr)
        sys.exit(1)
    except json.decoder.JSONDecodeError as e:
        print("Cant parse", file, "to json", "Error:", e, file=sys.stderr)
        sys.exit(1)

    if weights is None:
        print("Loaded weights from", file, "are empty", file=sys.stderr)
        sys.exit(1)
    elif not checkWeightsFormat(weights):
        print("Weights do not match expected format", file=sys.stderr)
        sys.exit(1)
    else:
        print("Weights loaded successfully from", file)

    return weights


def checkConfigurationFormat(config):
    try:
        keras.models.model_from_json(json.dumps(config))
    except ValueError:
        return False
    return True


def checkWeightsFormat(weights):
    if not isinstance(weights, list):
        return False
    if not len(weights) > 0:
        return False
    if not isinstance(weights[0], np.ndarray):
        return False
    if not len(weights[0]) > 0:
        return False
    return True


def writeUpdatesToDisk(new_weights, output_path):
    new_weights = [w.tolist() for w in new_weights]
    new_weights = json.dumps(new_weights)

    try:
        with open(output_path, "w") as output_file:
            output_file.write(new_weights)
    except IOError as e:
        print("Could not write to", output_path, "Error:", e)
        sys.exit(1)
    print("Saved updated weights to", output_path)


def absoluteFilePaths(directory):
    for dir_path, _, file_names in os.walk(directory):
        for f in file_names:
            yield os.path.abspath(os.path.join(dir_path, f))
