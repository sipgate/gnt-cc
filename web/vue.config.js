const path = require("path")

module.exports = {
  configureWebpack: {
    devServer: {
      proxy: {
        "/api": {
          target: "http://localhost:8000",
          pathRewrite: { "^/api": "" }
        }
      }
    }
  }
};
