const path = require('path');

module.exports = {
  mode: 'development',
  entry: './glue.js',
  output: {
    filename: 'index.js',
    path: path.resolve(__dirname, 'build')
  },
  devtool: 'source-map',
  performance: {
    hints: false
  }
};
