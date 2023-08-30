## TODO
(1) Go语言：识别全格式图片并执行相应的编码保存（支持webp） https://www.meiwen.com.cn/subject/gfiomctx.html
    image.Decode
(2) webp格式图片


## h2non/bimg
「GoCN酷Go推荐」Go 语言高性能图像处理神器 h2non/bimg
    https://mp.weixin.qq.com/s/kAFZohzJo2DiKkxjnVti6A
h2non/bimg 提供以下出片处理 API：
    调整大小
    放大
    裁剪（包括智能裁剪支持，libvips 8.5+）
    旋转（根据 EXIF 方向自动旋转）
    翻转（具有基于EXIF元数据的自动翻转）
    翻转
    缩略图
    获取大小
    水印（使用文本或图像）
    高斯模糊效果
    自定义输出颜色空间（RGB，灰度...）
    格式转换以及压缩处理
    EXIF元数据（大小，Alpha通道，配置文件，方向...）修改
    修剪（libvips 8.6+）

## problems
将透明背景的PNG转换为JPG（或JPEG），默认背景色为黑色
    https://www.zongscan.com/demo333/95729.html

## davidbyttow/govips
GO语言高性能图片处理库govips
    https://www.bilibili.com/video/BV1L14y1i7iG/
github: 
    https://github.com/davidbyttow/govips
服务器（linux || darwin）需要安装库：
(1) libvips
(2) gcc



