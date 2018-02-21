const CopyWebpackPlugin = require('copy-webpack-plugin');

const config = {
    entry: './app/app.js',
    output: {
        filename: './dist/bundle.js'
      },
    plugins: [
        new CopyWebpackPlugin([{ from: 'views', to: 'dist' }])
    ]
  };
  
  module.exports = config;
