## libvips

- [github](https://github.com/libvips/libvips)
- [官网](https://www.libvips.org/)

## .jpg 和 .jpeg

JPEG 和 JPG 实际上是同一种图像格式，本质上没有区别，只是文件扩展名不同。

## .webp

类似于 .png，也支持透明.

#### TODO: CGO_ENABLED=0 的情况下，将图片转换为 webp 格式

Richelieu: 目前找到的相关依赖，要么需求 libvips ，要么需求 libwebp ，都绕不开 C ，还是 vipKit 用吧.


