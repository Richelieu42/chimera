var prefix = "ws_";

var channel = null;
var connectBtn = document.getElementById("connectBtn"),
    disconnectBtn = document.getElementById("disconnectBtn");

connectBtn.onclick = function () {
    var {url, err} = getFinalUrl();
    if (err) {
        alert(err);
        return;
    }
    if (!url.startsWith("ws://") && !url.startsWith("wss://")) {
        alert("Invalid url!!!");
        return;
    }

    println("[建立连接]");

    connect(url);
    println(`url: [${url}].`);
};

disconnectBtn.onclick = function () {
    println("[断开连接]");

    disconnect();
};

function connect(url) {
    disconnect();

    channel = new WebSocket(url);
    channel.binaryType = "arraybuffer"; // "arraybuffer" 或者 "blob"
    channel.onopen = function () {
        println("onopen");
    };
    channel.onmessage = function (e) {
        let data = e.data;

        if (data instanceof ArrayBuffer) {
            processArrayBuffer(data, "ArrayBuffer")
        } else if (data instanceof Blob) {
            let blob = data

            let reader = new FileReader();
            reader.onload = function (event) {
                let ab = event.target.result;
                processArrayBuffer(ab, "Blob");
            };
            reader.readAsArrayBuffer(blob);
        } else if (typeof data === "string") {
            let text = e.data;
            println("on message(text, no gzip): " + text);
        }
    };
    channel.onerror = function (e) {
        println("onerror");
        console.error(e);
    };
    channel.onclose = function (e) {
        println("onclose: code(" + e.code + "), reason(" + e.reason + "), wasClean(" + e.wasClean + ")");
        console.error(e);
    };
}

function disconnect() {
    if (channel == null) {
        return;
    }

    channel.close();
    channel = null;
}

/*
 * 作用: 判断数据是否为GZIP压缩，用于判断是否需要解压缩
 *
 * @param data 待判断的数据，ArrayBuffer类型
 */
function isGzipCompressed(data) {
    // 这里仅作为示例，实际判断逻辑需根据数据源和协议来确定
    // 例如，检查数据头部的GZIP标识符（两个字节：0x1f 0x8b）
    const view = new DataView(data);
    return view.getUint8(0) === 0x1f && view.getUint8(1) === 0x8b;
}

// 示例函数：检查数据是否为JSON格式（实际实现可能需要根据数据结构或协议来确定）
function isJsonData(binaryData) {
    // 这里仅作为示例，实际判断逻辑需根据数据源和协议来确定
    // 例如，检查数据头部是否有常见的JSON起始字符（如 '{' 或 '['）
    const firstByte = new Uint8Array(binaryData)[0];
    return firstByte === 0x7B || firstByte === 0x5B; // '{' or '[' in ASCII
}

function processArrayBuffer(data, typeStr) {
    if (isGzipCompressed(data)) {
        // 解压数据（ArrayBuffer => Uint8Array）
        let decompressedData = pako.inflate(new Uint8Array(data));

        let decoder = new TextDecoder();
        let text = decoder.decode(decompressedData);
        println(`on message(binary, ${typeStr}, gzip): ${text}`)
    } else {
        let decoder = new TextDecoder();
        let text = decoder.decode(data);
        println(`on message(binary, ${typeStr}, no gzip): ${text}`)
    }
}

{
    document.getElementById("sendBtn").onclick = function () {

        let text = document.getElementById("textToSend").value;
        if (!text) {
            alert("Message to send is empty.");
            return;
        }

        if (!channel) {
            alert("channel == null");
            return;
        }
        if (channel.readyState !== 1){
            alert(`Channel(readyState: ${channel.readyState}) isn't ready.`);
        }

        println(`[发送消息] text: ${text}`)
        channel.send(text);
    };
}
