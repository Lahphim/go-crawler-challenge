const path = require('path');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');

module.exports = {
  entry: {
    application: path.resolve(__dirname, '../..', 'assets', 'javascripts', 'application.js')
  },
  plugins: [
    new CleanWebpackPlugin()
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
