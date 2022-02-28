window.onload = function() {
    var input = document.getElementById('input');
    input.addEventListener('change', handleFiles);
}

function handleFiles(e) {
    var ctx = document.getElementById('canvas').getContext('2d');
    var url = URL.createObjectURL(e.target.files[0]);
    var img = new Image();
    img.onload = function() {
        ctx.drawImage(img, 0, 0, 600, 400);
    }
    img.src = url;
}

$('.canvas').click(function(e){
    var target = this.getBoundingClientRect();
    var x = e.clientX - target.left;
    var y = e.clientY - target.top;
    $('#coord-click').html(x + ', ' + y);
});