var urlInput = document.getElementById("url"),
    groupInput = document.getElementById("group"),
    userInput = document.getElementById("user"),
    bsidInput = document.getElementById("bsid");

init();

function init(pre) {
    var url = localStorage[prefix + "url"] || "",
        group = localStorage[prefix + "group"] || "",
        user = localStorage[prefix + "user"] || "",
        bsid = localStorage[prefix + "bsid"] || "";

    urlInput.value = url;
    groupInput.value = group;
    userInput.value = user;
    bsidInput.value = bsid;
}

function setToLocalStorage(url, group, user, bsid) {
    localStorage[prefix + "url"] = url;
    localStorage[prefix + "group"] = group;
    localStorage[prefix + "user"] = user;
    localStorage[prefix + "bsid"] = bsid;
}

function getFinalUrl() {
    var rst = "",
        err = "";

    var url = urlInput.value,
        group = groupInput.value,
        user = userInput.value,
        bsid = bsidInput.value;

    if (!url) {
        err = "Please enter url!!!";
        return {
            url: rst,
            err: err,
        };
    }

    var queryString;
    if (!group && !user && !bsid) {
        queryString = "";
    } else {
        queryString = `?group=${encodeURIComponent(group)}&user=${encodeURIComponent(user)}&bsid=${encodeURIComponent(bsid)}`
    }

    if (queryString) {
        var index = url.indexOf("?");
        if (index !== -1) {
            url = url.substring(0, index);
        }
        urlInput.value = url;

        rst = url + queryString;
    } else {
        rst = url;
    }

    setToLocalStorage(url, group, user, bsid);

    rst = url + queryString
    return {
        url: rst,
        err: err,
    };
}

var output = document.getElementById("output"),
    clearBtn = document.getElementById("clearBtn");

clearBtn.onclick = clearOutput;

function clearOutput() {
    output.value = "";
}

function println(text) {
    var str = getTimeString() + " " + text

    output.value += str + "\n";
    console.info(str);
}

function getTimeString() {
    let d = new Date();
    return `${d.getFullYear()}-${d.getMonth() + 1}-${d.getDate()} ${d.getHours()}:${d.getMinutes()}:${d.getSeconds()}.${d.getMilliseconds()}`
}

