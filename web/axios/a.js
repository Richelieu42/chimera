console.log(axios);
console.log(666);

var url = "https://blog.csdn.net/iijik55/article/details/123250856";

// 发送GET请求
axios.get(url)
    .then(response => {
        console.log(response.data); // 处理响应数据
    })
    .catch(error => {
        console.error(error); // 处理错误情况
    });

// 发送POST请求 && 设置请求头
axios.post(url, {
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