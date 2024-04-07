var url = "http://127.0.0.1/test";

var params = new URLSearchParams();
params.append('name', '张三');
params.append('param1', 'value1');
params.append('param2', 'value2');

a()

async function a() {
    try {
        var response = await axios.post(url, params, {
            headers: {
                'Authorization': 'Bearer yourToken'
            }
        });

        // 处理响应数据
        console.log(response.data);
    } catch (error) {
        console.log(error);
    }
}

// axios.post(url, params, {
//     headers: {
//         'Authorization': 'Bearer yourToken'
//     }
// })
//     .then(response => {
//         console.log(response.data); // 处理响应数据
//     })
//     .catch(error => {
//         if (error.response) {
//             // 请求已发出，但服务器响应的状态码不在 2xx 范围内
//             console.log(error.response.data);
//             console.log(error.response.status);
//             console.log(error.response.headers);
//         } else {
//             // Something happened in setting up the request that triggered an Error
//             console.log('Error', error.message);
//         }
//         console.log(error.config);
//     });