const path = require('path');

module.exports = {
  mode: 'development',
  entry: './js-src/app.js',
  output: {
    filename: 'index.js',
    path: path.resolve(__dirname, 'app')
  },
  devtool: 'source-map',
  performance: { hints: false }
};
