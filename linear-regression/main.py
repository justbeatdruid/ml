import numpy as np
import tensorflow as tf
#import matplotlib.pyplot as plt

def generate_dataset():
  x_batch = np.linspace(0, 2, 100)
  y_batch = 1.5 * x_batch + np.random.randn(*x_batch.shape) * 0.2 + 0.5
  #for i in range(len(x_batch)):
  #  print("{},{}".format(x_batch[i], y_batch[i]))
  #exit(0)
  return x_batch, y_batch

def load_dataset(filename):
  reader = tf.TextLineReader(skip_header_lines=0)
  _, csv_raw = reader.read(filename)
  record_defaults = [[0.0], [0.0]]
  x_batch, y_batch = tf.decode_csv(csv_raw, record_defaults=record_defaults)
  x_batch, y_batch
  return x_batch, y_batch

def linear_regression():
  x = tf.placeholder(tf.float32, shape=(None, ), name='x')
  y = tf.placeholder(tf.float32, shape=(None, ), name='y')

  with tf.variable_scope('lreg') as scope:
    w = tf.Variable(np.random.normal(), name='W')
    b = tf.Variable(np.random.normal(), name='b')
		
    y_pred = tf.add(tf.multiply(w, x), b)

    loss = tf.reduce_mean(tf.square(y_pred - y))

  print("x={},y={},y_pred={},loss={}".format(x, y, y_pred, loss))
  return x, y, y_pred, loss

def run():
  #x_batch, y_batch = generate_dataset()
  filenames = ["train_data.csv"]
  x_batch, y_batch = load_dataset(tf.train.string_input_producer(filenames, num_epochs=1, shuffle=False))
  x, y, y_pred, loss = linear_regression()

  optimizer = tf.train.GradientDescentOptimizer(0.1)
  train_op = optimizer.minimize(loss)

  with tf.Session() as session:
    session.run(tf.global_variables_initializer())
    feed_dict = {x: x_batch, y: y_batch}
		
    for i in range(30):
      session.run(train_op, feed_dict)
      print(i, "loss:", loss.eval(feed_dict))

    print('Predicting...')
    y_pred_batch = session.run(y_pred, {x : x_batch})

    print('Result: {}'.format(y_pred_batch))

  #plt.scatter(x_batch, y_batch)
  #plt.plot(x_batch, y_pred_batch, color='red')
  #plt.xlim(0, 2)
  #plt.ylim(0, 2)
  #plt.savefig('plot.png')

if __name__ == "__main__":
  run()
