const path = require('path')
module.exports = {
    outputDir: process.env.VUE_APP_OUTPUT_DIR || 'dist',
    publicPath: process.env.VUE_APP_PUBLIC_PATH || './',
    productionSourceMap: false,
    configureWebpack: {
        resolve: {
            alias: {
                '@': path.resolve(__dirname, 'src')
            }
        }
    },
    devServer: {
        port: 8080,
        proxy: { '/nacos': { target: 'http://localhost:8848', changeOrigin: true } }
    }
}