const VueLoaderPlugin = require('vue-loader/lib/plugin');
const webpack = require("webpack");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const ExtractTextPlugin = require("extract-text-webpack-plugin");
const fs = require('fs');
const path = require('path')
const devMode = process.env.NODE_ENV !== 'production';
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

function getFileNameNoExt(filename){
    return filename.split('.').slice(0, -1).join('.')
}

function requireAll(folder, filter, subf_, files_) {
    files_ = files_ || [];
    subf_ = subf_ || '';
    fs.readdirSync(folder).forEach(function (filename) {
        const file = path.join(folder, filename);
        const stat = fs.lstatSync(file);
        if (stat.isDirectory()){
            requireAll(file, filter, path.join(subf_, path.basename(file)), files_); //recurse
        }else if (filename.indexOf(filter)>=0) {
            files_.push(path.join(subf_, filename));
        }
    });
    return files_;
}
const t = {
    stats: {
        // One of the two if I remember right
        entrypoints: false,
        children: false
    },
    mode: process.env.NODE_ENV,
    optimization: {
        runtimeChunk: 'single',
        splitChunks: {
            chunks: 'all',
            maxInitialRequests: Infinity,
            minSize: 0,
            cacheGroups: {
                vendor: {
                    test: /[\\/]node_modules[\\/]/,
                    chunks: "initial",
                    name: "vendor",
                    priority: 10,
                    enforce: true
                },
                commons: {
                    chunks: "initial",
                    minChunks: 2,
                    maxInitialRequests: 5, // The default limit is too small to showcase the effect
                    minSize: 0 // This is example is too small to create commons chunks
                },
            },
        },
    },
    entry: {
    },
    // entry: '',
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.esm.js'
        },
        extensions: ['*', '.js', '.vue', '.json']
    },
    output: {
        filename: "[name].js",
        path: __dirname + "/public/dist"
    },
    plugins: [
        // new webpack.HotModuleReplacementPlugin(),
        // new FriendlyErrorsPlugin(),
        new webpack.ProvidePlugin({
            $: 'jquery',
            jQuery: 'jquery',
            'window.jQuery': 'jquery',
            Popper: ['popper.js', 'default'],
            Collapse: "exports-loader?Collapse!bootstrap/js/dist/collapse",
            Alert: "exports-loader?Alert!bootstrap/js/dist/alert",
            Dropdown: "exports-loader?Dropdown!bootstrap/js/dist/dropdown",
            Tooltip: "exports-loader?Tooltip!bootstrap/js/dist/tooltip",
            Tab: "exports-loader?Tab!bootstrap/js/dist/tab",
        }),
        // new webpack.optimize.CommonsChunkPlugin('vendor'),

        // new ExtractTextPlugin({
        //     filename: "[name].css"
        // }),
        new MiniCssExtractPlugin({
            // Options similar to the same options in webpackOptions.output
            // both options are optional
            filename: devMode ? '[name].css' : '[name].[hash].css',
            chunkFilename: devMode ? '[id].css' : '[id].[hash].css',
            // filename: '[name].[hash].css',
        }),
        // new MiniCssExtractPlugin(),
        // new VueLoaderPlugin(),
        new CopyWebpackPlugin(
            [
                {
                    from: "./assets",
                    to: ""
                }
            ],
            {
                ignore: ["scss/**/*", "js/**/*"]
            }
        )
    ],
    module: {
        rules: [
            {
                test: /\.(sa|sc|c)ss$/,
                use: [
                    {
                        loader: MiniCssExtractPlugin.loader,
                        options: {
                            hmr: !devMode,
                        },
                    },
                    'css-loader',
                    'sass-loader',
                ],
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader',
            },
            {
                test: /\.woff(\?v=\d+\.\d+\.\d+)?$/,
                use: "url-loader?limit=10000&mimetype=application/font-woff"
            },
            {
                test: /\.woff2(\?v=\d+\.\d+\.\d+)?$/,
                use: "url-loader?limit=10000&mimetype=application/font-woff"
            },
            {
                test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
                use: "url-loader?limit=10000&mimetype=application/octet-stream"
            },
            {
                test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
                use: "file-loader"
            },
            {
                test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
                use: "url-loader?limit=10000&mimetype=image/svg+xml"
            },
            {
                test: require.resolve("jquery"),
                use: "expose-loader?jQuery!expose-loader?$"
            }
        ]
    }
};

const loadJs = requireAll(__dirname + '/assets/js/export/', '.js');
const loadScss = requireAll(__dirname + '/assets/scss/export/', '.scss');
loadJs.forEach(name => { //[ 'bs.js', 'home/home.js', 'home/home2.js' ]
    //"frontend/home/home": "./assets/js/frontend/home/home.js",
    t.entry['js/' + getFileNameNoExt(name)] = './assets/js/export/' + name
});
loadScss.forEach(name => {
    t.entry['css/' + getFileNameNoExt(name)] = './assets/scss/export/' + name
});

module.exports = t;
