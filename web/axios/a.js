var url = "http://127.0.0.1/test";

// 发送POST请求 && 设置请求头
axios.post(url, {
    name: 'world',
    key1: 'value1',
    key2: 'value2'
}, {
    headers: {
        'Authorization': 'Bearer yourToken'
    }
})
    .then(response => {
        console.log(response.data); // 处理响应数据
    })
    .catch(error => {
        console.error(error); // 处理错误情况
    });