const path = require('path');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');

module.exports = {
  mode: 'development',
  entry: {
    application: path.resolve(__dirname, '../..', 'assets', 'javascripts', 'application.js')
  },
  devtool: 'inline-source-map',
  plugins: [
    new CleanWebpackPlugin()
  ],
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, '../..', 'static', 'js')
  }
};
