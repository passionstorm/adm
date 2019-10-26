const VueLoaderPlugin = require('vue-loader/lib/plugin');
const webpack = require("webpack");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const ExtractTextPlugin = require("extract-text-webpack-plugin");
const fs = require('fs');
function requireAll(folder, filter) {
    fs.readdirSync(folder).forEach(function (filename) {
        var stat = fs.lstatSync(filename);
        if (stat.isDirectory()){
            requireAll(folder,filter); //recurse
        }else if (filename.indexOf(filter)>=0) {
            console.log('-- found: ',filename);
        };
    });
}
requireAll('./assets/js/', '.js');
const t = {}
console.log("a");
// const t = {
//     mode: 'development',
//     entry: {
//         admin: "./assets/js/admin.js",
//         // frontend: "./assets/js/frontend.js",
//         // "frontend/home/home": "./assets/js/frontend/home/home.js",
//         // "admin-style": "./assets/scss/admin.scss",
//         // "frontend-style": "./assets/scss/frontend.scss",
//     },
//     // entry: '',
//     resolve: {
//         alias: {
//             'vue$': 'vue/dist/vue.esm.js'
//         },
//         extensions: ['*', '.js', '.vue', '.json']
//     },
//     output: {
//         filename: "[name].js",
//         path: __dirname + "/public/assets"
//     },
//     plugins: [
//         // new webpack.HotModuleReplacementPlugin(),
//         // new FriendlyErrorsPlugin(),
//         new webpack.ProvidePlugin({
//             $: 'jquery',
//             jQuery: 'jquery',
//             'window.jQuery': 'jquery',
//             Popper: ['popper.js', 'default'],
//             Collapse: "exports-loader?Collapse!bootstrap/js/dist/collapse",
//             Alert: "exports-loader?Alert!bootstrap/js/dist/alert",
//             Dropdown: "exports-loader?Dropdown!bootstrap/js/dist/dropdown",
//             Tooltip: "exports-loader?Tooltip!bootstrap/js/dist/tooltip",
//             Tab: "exports-loader?Tab!bootstrap/js/dist/tab",
//         }),
//         // new webpack.optimize.CommonsChunkPlugin('vendor'),
//
//         new ExtractTextPlugin({
//             filename: "[name].css"
//         }),
//         // new MiniCssExtractPlugin(),
//         // new VueLoaderPlugin(),
//         new CopyWebpackPlugin(
//             [
//                 {
//                     from: "./assets",
//                     to: ""
//                 }
//             ],
//             {
//                 ignore: ["scss/**/*", "js/**/*"]
//             }
//         )
//     ],
//     module: {
//         rules: [
//             {
//                 test: /\.scss$/,
//                 use: ExtractTextPlugin.extract({
//                     fallback: "style-loader",
//                     use: [
//                         {
//                             loader: "css-loader",
//                             options: {
//                                 sourceMap: false
//                             }
//                         },
//                         {
//                             loader: "sass-loader",
//                             options: {
//                                 sourceMap: false
//                             }
//                         }
//                     ]
//                 })
//             },
//             {
//                 test: /\.css$/,
//                 use: ["style-loader", "css-loader"]
//             },
//             {
//                 test: /\.vue$/,
//                 loader: 'vue-loader',
//             },
//             {
//                 test: /\.woff(\?v=\d+\.\d+\.\d+)?$/,
//                 use: "url-loader?limit=10000&mimetype=application/font-woff"
//             },
//             {
//                 test: /\.woff2(\?v=\d+\.\d+\.\d+)?$/,
//                 use: "url-loader?limit=10000&mimetype=application/font-woff"
//             },
//             {
//                 test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
//                 use: "url-loader?limit=10000&mimetype=application/octet-stream"
//             },
//             {
//                 test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
//                 use: "file-loader"
//             },
//             {
//                 test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
//                 use: "url-loader?limit=10000&mimetype=image/svg+xml"
//             },
//             {
//                 test: require.resolve("jquery"),
//                 use: "expose-loader?jQuery!expose-loader?$"
//             }
//         ]
//     }
// };
module.exports = t;
