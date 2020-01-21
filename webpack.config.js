const path = require('path');
const TerserPlugin = require('terser-webpack-plugin');
const { SourceMapDevToolPlugin } = require('webpack');
const VueLoaderPlugin = require('vue-loader/lib/plugin');

const entryPoints = {
  index: ['./web_src/js/index'],
  swagger: ['./web_src/js/swagger'],
  jquery: ['./web_src/js/jquery'],
};

const options = {
  mode: 'production',
  stats: 'minimal',
  devtool: false,
  output: {
    path: path.resolve(__dirname, 'public/js'),
    filename: '[name].js',
    chunkFilename: '[name].js',
  },
  optimization: {
    minimize: true,
    minimizer: [new TerserPlugin({
      sourceMap: true,
      extractComments: false,
      terserOptions: {
        output: {
          comments: false,
        },
      },
    })],
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        exclude: /node_modules/,
        loader: 'vue-loader'
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: [
          'thread-loader',
          {
            loader: 'babel-loader',
            options: {
              presets: [
                [
                  '@babel/preset-env',
                  {
                    useBuiltIns: 'usage',
                    corejs: 3,
                  }
                ]
              ],
              plugins: [
                [
                  '@babel/plugin-transform-runtime',
                  {
                    regenerator: true,
                  }
                ],
                '@babel/plugin-proposal-object-rest-spread',
              ],
            },
          },
        ],
      },
      {
        test: /\.css$/i,
        use: ['style-loader', 'css-loader'],
      },
    ]
  },
  plugins: [
    new VueLoaderPlugin(),
    new SourceMapDevToolPlugin({
      filename: '[name].js.map',
      exclude: [
        'swagger.js',
        'jquery.js',
      ],
    }),
  ],
  performance: {
    maxEntrypointSize: 512000,
    maxAssetSize: 512000,
    assetFilter: (filename) => {
      return !filename.endsWith('.map') && filename !== 'swagger.js';
    }
  },
};

module.exports = Object.keys(entryPoints).map((key) => {
  return { entry: { [key]: entryPoints[key] }, ...options };
});
