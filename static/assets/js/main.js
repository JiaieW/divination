/**
* Template Name: WeBuild
* Updated: Sep 18 2023 with Bootstrap v5.3.2
* Template URL: https://bootstrapmade.com/free-bootstrap-coming-soon-template-countdwon/
* Author: BootstrapMade.com
* License: https://bootstrapmade.com/license/
*/
(function() {
  "use strict";

  /**
   * Easy selector helper function
   */
  const select = (el, all = false) => {
    el = el.trim()
    if (all) {
      return [...document.querySelectorAll(el)]
    } else {
      return document.querySelector(el)
    }
  }

})()

function updateCurrentTime() {
  var currentTime = new Date();
  // 将当前时间转换为东八区时间
  var offset = 8; // UTC+8
  var utc = currentTime.getTime() + currentTime.getTimezoneOffset() * 60000;
  var localTime = new Date(utc + (3600000 * offset));
  
  // 格式化时间显示
  var formattedTime = localTime.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: false
});
  document.getElementById('current-time').innerHTML = formattedTime;
}

// 每秒更新时间
setInterval(updateCurrentTime, 1000);

// 初始化显示
updateCurrentTime();