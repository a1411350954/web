"use strict";
(function () {
    if(document.documentElement.clientWidth > 760){
        var particlesDiv = document.createElement("div");
        particlesDiv.setAttribute("id","particles-js");
        particlesDiv.style.width='100%';
        particlesDiv.style.height='100%';
        document.body.appendChild(particlesDiv);
        particlesJS.load('particles-js', '/static/js/bg/particles.json', function() {
        });
    }

})();
function setSrcQuery(e, q) {
    var src  = e.src;
    var p = src.indexOf('?');
    if (p >= 0) {
        src = src.substr(0, p);
    }
    e.src = src + "?" + q
}
/**
 *
 * 重新加载验证码
 *
 */
function reload() {
    setSrcQuery(document.getElementById('validateImage'), "reload=" + (new Date()).getTime());
    return false;
}

/**
 *
 * 设置登录窗口位置
 *
 */
var setMainMargin = function () {
    var main = document.getElementById("main");
    var mainHeight = main.offsetHeight;
    var mainWidth = main.offsetWidth;
    var bodyHeight = document.documentElement.clientHeight;
    var bodyWidth = document.documentElement.clientWidth;

    if(bodyWidth>760){
        main.style.position = "absolute";
        var margin_top = (bodyHeight-mainHeight)/2;
        var margin_left = (bodyWidth-mainWidth)/2;
        if(margin_left<0){
            margin_left = 0;
        }
        if(margin_top<0){
            margin_top=0;
        }
        main.style.top = margin_top+"px";
        main.style.left = margin_left+"px";
    }

}

window.onresize = setMainMargin;
setMainMargin();

function login(e) {
    e.preventDefault();
    $.ajax({
        type:"POST"
        ,url:"/login"
        ,data:$("#loginForm").serialize()
        ,success:function(data){
            console.info(data)
        },error:function(data){
            console.error(data)
        }
    });
}