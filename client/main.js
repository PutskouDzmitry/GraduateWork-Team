const canvas = document.getElementById('canvas');
const ctx = canvas.getContext('2d');

// Рисуем фон
ctx.fillStyle = '#FD0';
ctx.fillRect(0, 0, 75, 75);
ctx.fillStyle = '#6C0';
ctx.fillRect(75, 0, 75, 75);
ctx.fillStyle = '#09F';
ctx.fillRect(0, 75, 75, 75);
ctx.fillStyle = '#F30';
ctx.fillRect(75, 75, 75, 75);
ctx.fillStyle = '#FFF';

// Устанавливаем уровень прозрачности
ctx.globalAlpha = 0.8;

// Рисуем круги
ctx.fillStyle = 'blue';
ctx.fillRect(70, 70, 100, 100);

// for (let i = 0; i < 7; i++) {
//     ctx.beginPath();
//     ctx.arc(75, 75, 10 + 10 * i, 0, Math.PI * 2, true);
//     ctx.fill();
// }