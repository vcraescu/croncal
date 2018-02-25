const
    webpack = require('webpack'),
    path = require('path'),
    merge = require('webpack-merge'),
    webPath = path.resolve(__dirname, '../public'),
    ExtractTextPlugin = require('extract-text-webpack-plugin'),
    CleanWebpackPlugin = require('clean-webpack-plugin'),

    extractSass = new ExtractTextPlugin({
        filename: '[name].min.css',
    })

const common = {
    context: path.resolve(__dirname),
    output: {
        filename: '[name].min.js',
        path: webPath,
        publicPath: '/public',
        hotUpdateChunkFilename: 'compiled/hot-update/hot-update.js',
        hotUpdateMainFilename: 'compiled/hot-update/hot-update.json',
    },
    devtool: '#eval-source-map',
    plugins: [
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: `"${process.env.NODE_ENV}"`,
            },
        }),
        extractSass,
    ],
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.esm.js',
        },
    },
    module: {
        rules: [
            {
                test: /\.js?$/,
                loader: 'babel-loader',
                exclude: [
                    path.resolve(__dirname, 'node_modules'),
                ],
            },
            {
                test: /\.(js|vue)?$/,
                enforce: 'pre',
                loader: 'eslint-loader',
                exclude: [
                    path.resolve(__dirname, 'node_modules'),
                ],
                options: {
                    failOnError: true,
                },
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader'],
            },
            {
                test: /\.scss$/,
                use: extractSass.extract({
                    use: [
                        {
                            loader: 'css-loader',
                            options: {
                                minimize: (process.env.NODE_ENV ===
                                    'production'),
                                sourceMap: (process.env.NODE_ENV ===
                                    'production'),
                            },
                        },
                        {
                            loader: 'resolve-url-loader',
                        },
                        {
                            loader: 'sass-loader',
                        },
                    ],
                }),
            },
            {
                test: /\.(png|jpg|jpeg|gif|svg)$/,
                use: {
                    loader: 'file-loader',
                    options: {
                        name: process.env.NODE_ENV === 'production'
                            ? '[path][name].[hash].[ext]'
                            : '[path][name].[ext]',
                        publicPath: '/',
                    },
                },
            },
            {
                test: /\.(eot|svg|ttf|woff|woff2)$/,
                use: {
                    loader: 'file-loader',
                    options: {
                        name: 'compiled/fonts/[name].[ext]',
                        publicPath: '/',
                    },
                },
            },
        ],
    },
    devServer: {
        compress: false,
        disableHostCheck: true,
        inline: false,
        publicPath: '/',
        port: 8000,
        host: '0.0.0.0',
        noInfo: true,
    },
}

let theme = {
    entry: {
        'compiled/style': path.resolve(__dirname, 'scss/style.scss'),
    },
    resolve: {
        alias: {
            'style': path.resolve(__dirname, 'assets/scss/style.scss'),
        },
    },
}

let apps = {
    entry: {
        'compiled/calendar': path.resolve(__dirname, 'js/calendar/app.js'),
    },
    plugins: [
        new webpack.optimize.CommonsChunkPlugin({
            name: 'vendor',
            chunks: ['compiled/calendar'],
            filename: 'compiled/vendor.min.js',
        }),
    ],
    module: {
        rules: [
            {
                test: /\.vue$/,
                loader: 'vue-loader',
            },
        ],
    },
}

module.exports = merge(common, apps, theme)

if (process.env.NODE_ENV === 'production') {
    module.exports.devtool = '#source-map'
    module.exports.plugins = (module.exports.plugins || []).concat([
        new webpack.optimize.OccurrenceOrderPlugin(),
        new webpack.optimize.UglifyJsPlugin({
            sourceMap: false,
            ecma: 8,
            compress: {
                warnings: false,
            },
            parallel: true,
            output: {
                comments: false,
            },
        }),
        new CleanWebpackPlugin(['hot-update'], {
            root: path.resolve(__dirname, '../public/compiled'),
        }),
    ])
}
