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
      },
      {
        test: /\.(woff2|ttf)$/,
        use: [
          'url-loader'
        ]
      }
    ]
  },

  output: {
    filename: 'javascripts/[name].js',
    path: path.resolve(__dirname, '../..', 'static')
  },

  resolve: {
    alias: {
      Components: path.resolve(__dirname, '../..', 'assets', 'javascripts', 'components'),
      Helpers: path.resolve(__dirname, '../..', 'assets', 'javascripts', 'helpers')
    },
    extensions: ['.js']
  }
};
