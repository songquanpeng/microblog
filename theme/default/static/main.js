let mainElement = undefined;
let loading = false;
let offset = 0;
let token = localStorage.getItem('token');
let status = {
    authed: false,
    version: "v0.0.0",
    author: ""
};

let colorList = [
    "#0074D9", "#7FDBFF", "#39CCCC", "#B10DC9", "#F012BE",
    "#FF4136", "#FF851B", "#2ECC40", "#01FF70"
];

async function main() {
    mainElement = document.getElementById('main');
    await loadMore();
    status = await loadStatus();
    window.onscroll = async function () {
        if (shouldLoad()) {
            await loadMore();
        }
    }
}

function onNewBtnClicked() {
    if (status.authed) {
        showModal('newModal');
    } else {
        showModal('authModal');
    }
}

function text2color(text) {
    let color = "#111111";
    if (text.length !== 0) {
        let firstChar = text[0];
        let n = firstChar.charCodeAt(0);
        if (!isNaN(n)) {
            let idx = n % colorList.length;
            color = colorList[idx];
        }
    }
    return color;
}


function timestamp2time(timestamp) {
    let time = new Date(timestamp);
    let year = time.getFullYear().toString();
    let month = (time.getMonth() + 1).toString();
    let day = time.getDate().toString();
    let hour = time.getHours().toString();
    let minute = time.getMinutes().toString();
    let second = time.getSeconds().toString();
    if (month.length === 1) {
        month = "0" + month;
    }
    if (day.length === 1) {
        day = "0" + day;
    }
    if (hour.length === 1) {
        hour = "0" + hour;
    }
    if (minute.length === 1) {
        minute = "0" + minute;
    }
    if (second.length === 1) {
        second = "0" + second;
    }
    return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
}

function render(item, insertEnd = true) {
    item.content = marked.parse(item.content);
    let itemElement = `<div class="card item" id="item-${item.id}">
            <div class="card-content">
                <div class="content">
                    <div class="markdown">${item.content}</div>
                    <time>${timestamp2time(item.timestamp * 1000)}</time>
                    <p class="id-tag">#${item.id}</p>
                </div>
            </div>
        </div>`;
    if (insertEnd) {
        mainElement.insertAdjacentHTML('beforeend', itemElement);
    } else {
        mainElement.insertAdjacentHTML('afterbegin', itemElement);
    }
}

function parseLink(text) {
    let urlPattern = /(\b(https?|ftp):\/\/[-A-Z0-9+&@#\/%?=~_|!:,.;]*[-A-Z0-9+&@#\/%=~_|])/gim;
    text = text.replace(urlPattern, '<a href="$1" target="_blank">$1</a>');
    return text
}

function showModal(id) {
    document.getElementById(id).className = "modal is-active";
}

function closeModal(id) {
    document.getElementById(id).className = "modal";
}

async function deletePost(id) {
    let res = await fetch(`/api/post/${id}`, {
        method: 'DELETE',
    });
    return await res.json();
}

async function onSubmitBtnClicked() {
    let content = document.getElementById('editor').value;
    if (content.startsWith("delete #")) {
        let t = content.split('#');
        let id = t[t.length - 1];
        let data = await deletePost(id);
        if (data.success) {
            document.getElementById(`item-${id}`).style.display = 'none';
            closeModal("newModal");
        } else {
            document.getElementById('newModalTitle').textContent = "删除失败：" + data.message;
            console.error(data);
        }
        return;
    }
    let res = await fetch(`/api/post`, {
        method: 'POST',
        body: JSON.stringify({
            'content': content,
        })
    });
    let data = await res.json();
    let id = data.data;
    if (data.success) {
        closeModal("newModal");
        document.getElementById('editor').value = "";
        let res = await fetch(`/api/post/${id}`);
        let data = await res.json();
        if (data.success) {
            render(data.data, false);
            offset += 1;
            window.scrollTo(0, 0);
        }
    } else {
        document.getElementById('newModalTitle').textContent = "发布失败：" + data.message;
        console.error(data);
    }
}

async function loadStatus() {
    let res = await fetch(`/api/status`);
    let data = await res.json()
    if (data.success) {
        return data.data;
    }
}

async function loadData(start) {
    let res = await fetch(`/api/post/?p=${start}`);
    let data = await res.json()
    if (data.success) {
        return data.data;
    } else {
        return [];
    }
}

async function loadMore() {
    if (loading) return;
    loading = true;
    let items = await loadData(offset)
    offset += items.length;
    items.forEach((item) => {
        render(item);
    })
    loading = false;
}

function shouldLoad() {
    return (window.innerHeight + window.scrollY + 5) >= document.body.offsetHeight
}

async function login() {
    let username = document.getElementById('usernameInput').value.trim();
    let password = document.getElementById('passwordInput').value.trim();
    if (username === "" || password === "") {
        return
    }
    let res = await fetch(`/api/login`, {
        method: 'POST',
        body: JSON.stringify({
            username,
            password
        })
    });
    let data = await res.json();
    if (data.success) {
        status.authed = true;
        closeModal('authModal');
        showModal('newModal');
        document.getElementById('authModalTitle').innerText = "用户登录";
    } else {
        document.getElementById('authModalTitle').innerText = "登录失败：" + data.message;
    }
}
