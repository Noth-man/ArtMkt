$('.menu-open').on('click',function(){
    $(".menu-open").css('display','none');
    $(".menu-close").css('display','block');
    $(".menu-content-2").addClass('active');
});

$('.menu-close').on('click',function(){
    $(".menu-open").css('display','block');
    $(".menu-close").css('display','none');
    $(".menu-content-2").removeClass('active');

});

$(document).click(function(event) {
    if(!$(event.target).closest('.menu-open').length) {
        var winW = $(window).width();
        var devW = 799;
        if (winW <= devW) {
            $(".menu-open").css('display','block');
            $(".menu-close").css('display','none');
            $(".menu-content-2").removeClass('active');
        }
        
    }
  });