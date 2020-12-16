import tensorflow as tf
import numpy as np

import base64
from PIL import Image
from io import BytesIO
from google.cloud import storage
import configparser

interpreter = None
labels = None
input_details = None
output_details = None

def base64_to_PIL(base64_data):
    """
    This might be useful when converting to run as a cloud function
    :param base64_data: base64 encoding of the image
    :return: PIL Image object
    """
    im = Image.open(BytesIO(base64.b64decode(base64_data)))
    return im


def predict(image):
    global interpreter
    """
    Predict the contents of the image
    :param image: PIL Image object
    :return: label (str), confidence (np.float32)
    """
    image = image.convert("RGB").resize((224, 224))
    img = np.array(image)
    img_preproc = img.reshape(1, 224, 224, 3)
    img_preproc = img_preproc.astype(np.uint8)
    interpreter.set_tensor(input_details[0]['index'], img_preproc)
    interpreter.invoke()
    prob = interpreter.get_tensor(output_details[0]['index'])
    print(prob)
    index = np.argmax(prob[0])
    label = labels[index]
    confidence = np.float32(prob[0, index]) / 255.
    return label, confidence


def handle(request):
    from flask import abort
    global interpreter
    global labels
    global input_details
    global output_details
    if request.content is None:
        return 400
    img = base64_to_PIL(request.content)
    if interpreter is None:
        config = configparser.ConfigParser()
        config.read('config')
        bucket_id = config['bucket_id']
        model_file = config['model_file']
        label_file = config['label_file']
        destination_path = config['destination_path']
        load_model(bucket_id, model_file, destination_path)
        interpreter = tf.lite.Interpreter(model_path=destination_path + model_file)
        interpreter.allocate_tensors()
        input_details = interpreter.get_input_details()
        output_details = interpreter.get_output_details()
        with open(destination_path + label_file, "r") as f:
            labels = [l.strip() for l in f.readlines()]
    if img is None:
        return abort(405)
    return predict(img)


def load_model(bucket_id, model_file, label_file, destination_path):
    storage_client = storage.Client()
    bucket = storage_client.bucket(bucket_id)
    blob = bucket.blob(model_file)
    blob.download_to_filename(destination_path + model_file)
    blob = bucket.blob(label_file)
    blob.download_to_filename(destination_path + label_file)


if __name__ == "__main__":
    img = Image.open("bee.jpg")
    print(predict(img))
