$('.slider-yajirusi-right').on('click',function(){
    if($('.slider-content1').css('opacity') == '1'){
        
        $('.slider-content1').css({
            'opacity': '0',
            'transform': 'translateX(-120%)',
        });
        $('.slider-content2').css({
            'opacity': '1',
            'transform': 'translateX(0%)',
        });
        $('.slider-content3').css({
            'opacity': '0',
            'transform': 'translateX(120%)',
        });
    }else if($('.slider-content2').css('opacity') == '1'){
        $('.slider-content3').css({
            'opacity': '1',
            'transform': 'translateX(0%)',
        });
        $('.slider-content2').css({
            'opacity': '0',
            'transform': 'translateX(-120%)',
        });
        $('.slider-content1').css({
            'opacity': '0',
            'transform': 'translateX(120%)',
        });
    }else if($('.slider-content3').css('opacity') == '1'){
        $('.slider-content1').css({
            'opacity': '1',
            'transform': 'translateX(0%)',
        });
        $('.slider-content2').css({
            'opacity': '0',
            'transform': 'translateX(120%)',
        });
        $('.slider-content3').css({
            'opacity': '0',
            'transform': 'translateX(-120%)',
        });
    }
});

$('.slider-yajirusi-left').on('click',function(){
    if($('.slider-content1').css('opacity') == '1'){
        
        $('.slider-content1').css({
            'opacity': '0',
            'transform': 'translateX(120%)',
        });
        $('.slider-content3').css({
            'opacity': '1',
            'transform': 'translateX(0%)',
        });
        $('.slider-content2').css({
            'opacity': '0',
            'transform': 'translateX(-120%)',
        });
    }else if($('.slider-content2').css('opacity') == '1'){
        $('.slider-content1').css({
            'opacity': '1',
            'transform': 'translateX(0%)',
        });
        $('.slider-content2').css({
            'opacity': '0',
            'transform': 'translateX(120%)',
        });
        $('.slider-content3').css({
            'opacity': '0',
            'transform': 'translateX(-120%)',
        });
    }else if($('.slider-content3').css('opacity') == '1'){
        $('.slider-content2').css({
            'opacity': '1',
            'transform': 'translateX(0%)',
        });
        $('.slider-content1').css({
            'opacity': '0',
            'transform': 'translateX(-120%)',
        });
        $('.slider-content3').css({
            'opacity': '0',
            'transform': 'translateX(120%)',
        });
    }
});

var set;
var repeat = function() {
	// 10秒（1000ms）後に処理
	set = setTimeout(repeat, 10000);
    $('.slider-yajirusi-right').click();
}
repeat();