const path = require('path')

const resolve = dir => {
  return path.join(__dirname, dir)
}

// 项目部署基础
// 默认情况下，我们假设你的应用将被部署在域的根目录下,
// 例如：https://www.my-app.com/
// 默认：'/'
// 如果您的应用程序部署在子路径中，则需要在这指定子路径
// 例如：https://www.foobar.com/my-app/
// 需要将它改为'/my-app/'
// iview-admin线上演示打包路径： https://file.iviewui.com/admin-dist/
const BASE_URL = process.env.NODE_ENV === 'production'
  ? '/'
  : '/'

module.exports = {
  publicPath: BASE_URL,
  lintOnSave: false,
  chainWebpack: config => {
    config.resolve.alias
      .set('@', resolve('src')) // key,value自行定义，比如.set('@@', resolve('src/components'))
      .set('_c', resolve('src/components'))
      .set('_w', resolve('src/components/widgets'))
      .set('_v', resolve('src/views'))

    config.module
      .rule('svg')
      .test(() => false)
      .use('file-loader')
  },
  configureWebpack: {
    module: {
      rules: [
        {
          test: /\.(svg)(\?.*)?$/,
          use: [
            {
              loader: 'svg-inline-loader',
              options: {
                limit: 10000,
                name: 'assets/images/icons/[name].[hash:7].[ext]'
              }
            }
          ]
        },
      ],
    }
  },
  productionSourceMap: false
  // devServer: {
  //   proxy: 'localhost:3000'
  // }
}
