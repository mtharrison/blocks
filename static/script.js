function hexToRgbA(hex, a){
    var c;
    if(/^#([A-Fa-f0-9]{3}){1,2}$/.test(hex)){
        c= hex.substring(1).split('');
        if(c.length== 3){
            c= [c[0], c[0], c[1], c[1], c[2], c[2]];
        }
        c= '0x'+c.join('');
        return 'rgba('+[(c>>16)&255, (c>>8)&255, c&255].join(',')+',' + a + ')';
    }
    throw new Error('Bad Hex');
}

const colors = ['#e6194b', '#3cb44b', '#ffe119', '#4363d8', '#f58231', '#911eb4', '#46f0f0', '#f032e6', '#bcf60c', '#fabebe', '#008080', '#e6beff', '#9a6324', '#fffac8', '#800000', '#aaffc3', '#808000', '#ffd8b1', '#000075', '#808080', '#ffffff', '#000000'];

const sizeSlider = document.getElementById('size');
const coloursSlider = document.getElementById('colors');

let data;

function calculateData() {
    const colors = coloursSlider.value;
    data = new Array(sizeSlider.value * sizeSlider.value).fill(0).map(() => Math.floor(Math.random() * coloursSlider.value))
    render()
}

function render(highlight) {
    const canvas = document.getElementById('canvas');

    canvas.height = canvas.offsetHeight;
    canvas.width = canvas.offsetWidth;

    const ctx = canvas.getContext('2d');
    const width = canvas.width;
    const height = width;

    const size = width / sizeSlider.value

    ctx.clearRect(0, 0, width, height);

    data.forEach((color, i) => {

        ctx.fillStyle = hexToRgbA(colors[color], highlight ? (highlight.includes(i) ? 1 : 0.25) : 1);
        ctx.fillRect((i * size) % width, Math.floor((i * size) / width) * size, size, size);
    });
}

calculateData();

sizeSlider.addEventListener('input', calculateData);
coloursSlider.addEventListener('input', calculateData);
window.onresize = () => render();

document.getElementById('solve').addEventListener('click', async (e) => {

    e.preventDefault();

    const response = await fetch('/api/solve', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ data, size: parseInt(sizeSlider.value) })
    });

    const json = await response.json();

    render(json);
});
