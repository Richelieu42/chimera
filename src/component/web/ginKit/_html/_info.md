## 全局安装 html-minifier（以便后续使用 html-minifier 命令）

npm install -g html-minifier

## 命令（压缩404.html）

html-minifier --collapse-whitespace --remove-comments --remove-optional-tags --remove-redundant-attributes --remove-script-type-attributes --remove-tag-whitespace --use-short-doctype --minify-css true --minify-js true -o 404.min.html 404.html
