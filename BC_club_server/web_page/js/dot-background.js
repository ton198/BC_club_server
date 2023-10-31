let seed = Math.random();
let ctx = null;
let mouseX = -1;
let mouseY = -1;
seededRandom = function() {
    seed = (seed * 9301 + 49297) % 233280;
    return seed / 233280.0;
};

function startDrawing() {
    let canvas = document.getElementById('drawing_board');
    canvas.style.backgroundColor = '#000'
    ctx = canvas.getContext('2d');
    ctx.lineWidth = 0.5;
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    let dots = new Array(200);
    for (let i = 0;i < dots.length;i+=1) {
        dots[i] = new Dot(canvas);
    }

    window.addEventListener('resize', function () {
        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;
    });

    window.setInterval(function () {
        canvas.width = canvas.width;
        canvas.height = canvas.height;

        for (let i = 0;i < dots.length;i+=1) {
            dots[i].draw();
            if (calcDistance(dots[i].x - mouseX, dots[i].y - mouseY) < 100) {
                ctx.moveTo(dots[i].x + 1, dots[i].y + 1);
                ctx.lineTo(mouseX, mouseY);
                ctx.strokeStyle = "#ffffff"
                ctx.stroke();
            }
        }
    }, 25);
}

function calcDistance(vertical, horizontal) {
    return (vertical ** 2 + horizontal ** 2) ** 0.5
}

function mouseMovement(event) {
    mouseX = event.clientX;
    mouseY = event.clientY;
}

class Color {
    r = 0;
    g = 0;
    b = 0;
    constructor(r, g, b) {
        this.r = parseInt(r);
        this.g = parseInt(g);
        this.b = parseInt(b);
    }

    getColor() {
        return "#" + this.r.toString(16).padStart(2, '0') + this.g.toString(16).padStart(2, '0') + this.b.toString(16).padStart(2, '0');
    }
}

class Dot {
    x = 0;
    y = 0;
    vx = 0;
    vy = 0;
    color = '';
    ctx = null;
    width = 0;
    height = 0;
    constructor(canvas) {
        this.color = new Color(seededRandom() * 255, seededRandom() * 255, seededRandom() * 255).getColor();
        this.x = seededRandom() * canvas.width;
        this.y = seededRandom() * canvas.height;
        this.vx = seededRandom() * 1 - 0.5;
        this.vy = seededRandom() * 1 - 0.5;
        this.width = canvas.width;
        this.height = canvas.height;
        this.ctx = canvas.getContext('2d');
    }
    draw() {
        let dotSize = 3;
        this.x += this.vx;
        this.y += this.vy;
        if (this.x > this.width - dotSize) {
            this.vx = -this.vx;
            this.x = this.width - dotSize;
        } else if (this.x < 0) {
            this.vx = -this.vx;
            this.x = 0;
        }
        if (this.y > this.height - dotSize) {
            this.vy = -this.vy;
            this.y = this.height - dotSize;
        } else if (this.y < 0) {
            this.vy = -this.vy;
            this.y = 0;
        }
        this.ctx.fillStyle = this.color;
        this.ctx.fillRect(this.x, this.y, dotSize, dotSize);
    }
}

window.addEventListener('load', startDrawing);

window.addEventListener('mousemove', mouseMovement);
