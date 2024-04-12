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

/**
 * PS: EventSource 没有onclose事件.
 */
function connect(url) {
    disconnect();

    channel = new WebSocket(url);
    channel.binaryType = "arraybuffer"; // "arraybuffer" 或者 "blob"
    channel.onopen = function () {
        println("onopen");
    };
    channel.onmessage = async function (e) {
        let data = e.data;

        if (data instanceof ArrayBuffer) {
            // 方法1
            let decoder = new TextDecoder();
            let text = decoder.decode(data);
            println("on message(binary, ArrayBuffer): " + text);

            // 方法2
            // let blob = new Blob([data]);
            // let text = await blob.text();
            // println("on message(binary, ArrayBuffer): " + text);

            // 方法3
            // let blob = new Blob([data]);
            // let reader = new FileReader();
            // reader.readAsText(blob, "UTF-8");
            // reader.onload = () => {
            //     var text = reader.result;
            //     println("on message(binary, ArrayBuffer): " + text);
            // };
        } else if (data instanceof Blob) {
            // 方法1
            let text = await data.text();
            println("on message(binary, Blob): " + text);

            // 方法2
            // let reader = new FileReader();
            // reader.readAsText(data, "UTF-8");
            // reader.onload = () => {
            //     var text = reader.result;
            //     println("on message(binary, Blob): " + text);
            // };
        } else if (typeof data === "string") {
            let text = e.data;
            println("on message(text): " + text);
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
