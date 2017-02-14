var webpack = require("webpack");
var vue = require('vue-loader');
var path = require('path');

module.exports = {
    entry: {
        'bundle': path.join(__dirname, 'assets/js/app.js'),
    },
    output: {
        path: path.join(__dirname, 'assets/dist'),
        filename: "[name].js",
    },
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.common.js'
        },
        modules: [
            path.join(__dirname, '/assets/js'),
            path.join(__dirname, '/node_modules')
        ],
        extensions: ['.js', '.vue', '.css']
    },
    cache: true,
    module: {
        loaders: [{
            test: /\.css$/, loader: 'style-loader!css-loader',
        }, {
            test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loader: 'url-loader?limit=10000&mimetype=application/font-woff',
        }, {
            test: /\.(jpg|png)$/, loader: 'url-loader',
        }, {
            test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loader: 'file-loader',
        }, {
            test: /\.vue$/,
            loader: 'vue-loader',
            exclude: /node_modules/,
        }, {
            test: /\.js$/,
            exclude: /node_modules/,
            loader: 'babel-loader',
        }],
    },
    devtool: '#source-map',
    plugins: [
        new webpack.ProvidePlugin({
            jQuery: 'jquery',
            $: 'jquery',
            Tether: 'tether',
            Vue: 'vue',
            VueResource: 'vue-resource',
            VueRouter: 'vue-router'
        }),
    ],
}
