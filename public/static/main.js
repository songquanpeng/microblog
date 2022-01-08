let mainElement = undefined;
let loading = false;
let offset = 0;

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

async function loadData(start) {
    let res = await fetch(`/api/nonsense?p=${start}`);
    let data = await res.json()
    if (data.status) {
        return data.data;
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

async function main() {
    mainElement = document.getElementById('main');
    await loadMore();
    window.onscroll = async function () {
        if (shouldLoad()) {
            await loadMore();
        }
    }

}

function shouldLoad() {
    return (window.innerHeight + window.scrollY + 5) >= document.body.offsetHeight
}