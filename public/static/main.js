let mainElement = undefined;
let loading = false;
let offset = 0;
let num = 10;

function render(item) {
    let itemElement = `<div class="card item">
            <div class="card-content">
                <div class="content">
                    <p class="text">${item.content}</p>
                    <time>${item.time}</time>
                </div>
            </div>
        </div>`;
    mainElement.insertAdjacentHTML('beforeend', itemElement);
}

function loadData() {

}

function loadMore() {
    if (loading) return;
    loading = true;
    console.log(offset)
    for (let i = 0; i < num; i++) {
        render({content: "不，这只是你的错觉。在没有爱情的时候，你照样生活。", time: "2022-01-08 04:05"})
    }
    offset += num;
    loading = false;
}

function main() {
    mainElement = document.getElementById('main');

    loadMore();
    window.onscroll = function () {
        if (shouldLoad()) {
            loadMore();
        }
    }

}

function shouldLoad() {
    return (window.innerHeight + window.scrollY + 5) >= document.body.offsetHeight
}