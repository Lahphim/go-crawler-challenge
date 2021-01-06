const path = require('path');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const CopyPlugin = require("copy-webpack-plugin");

module.exports = {
  entry: {
    application: path.resolve(__dirname, '../..', 'assets', 'javascripts', 'application.js')
  },
  plugins: [
    new CleanWebpackPlugin(),
    new CopyPlugin({
      patterns: [
        {
          from: path.resolve(__dirname, '../..', 'assets', 'images'),
          to: path.resolve(__dirname, '../..', 'static', 'images')
        }
      ]
    })
  ],
  module: {
    rules: [
      {
        test: /\.(sa|sc|c)ss$/,
        use: [
          'style-loader',
          'css-loader',
          'postcss-loader',
          'sass-loader'
        ]
      }
    ]
  },
  output: {
    filename: 'javascripts/[name].js',
    path: path.resolve(__dirname, '../..', 'static')
  }
};
