<template>
  <div id="app">
    <Index></Index>
  </div>
</template>

<script>
import Index from "./Index";
import Vue from "vue";
import VueResource from "vue-resource";
import ElementUI from "element-ui";

//获取本机地址判断是否本地调试
var os = require("os");
var IPv4, hostName;


Vue.use(VueResource);
Vue.use(ElementUI);

Vue.http.options.emulateJSON = false;
Vue.http.options.xhr = { withCredentials: true };
Vue.http.interceptors.push((request, next) => {
  request.credentials = true;
  next();
});

console.log(Vue.http);
/* WebApp的一些初始化工作 */
var obj = {
  server: {
    address: "http://127.0.0.1:8080"
  }
};
hostName = os.hostname();
obj.server.address = "http://"+hostName+":8080"
window.webapp = obj;




/* WebApp的一些初始化工作 */

window.webapp.uploadurl = window.webapp.server.address + "/file";
window.webapp.fileurl = window.webapp.server.address + "/file/";
/* 一些接口功能 */
window.webapp.f = {};

/* 错误码 */
window.webapp.f.ErrorHandle = function(res) {
  switch (res.error.ErrorCode) {
    case 0:
      alert("参数错误");
      break;
    case 1:
      alert("未知错误");
      break;
    case 2:
      alert("登录失败");
      break;
    case 3:
      alert("请先登录");
      break;
    default:
      alert("发生某些错误");
  }
};
//获取用户信息
window.webapp.f.getUserInformation = function(vue, uid, success) {
  var url;
  if (uid == undefined || uid == "")
    url = window.webapp.server.address + "/user";
  else {
    url = window.webapp.server.address + "/user?uid=" + uid;
  }
  vue.$http.get(url).then(function(res) {
    if (res.body.status == "ok") {
      if (res.body.status == "ok") {
        success(res);
      } else {
        window.webapp.f.ErrorHandle(res.body);
      }
    }
  });
};
//登录
window.webapp.f.login = function(vue, form, success) {
  var request = window.webapp.server.address;
  request += "/user/login";

  vue.$http.post(request, form).then(function(res) {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

//注销
window.webapp.f.logout = function(vue, success) {
  var request = window.webapp.server.address;
  request += "/user/logout";
  vue.$http.get(request).then(function(res) {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

//获取文章详细
window.webapp.f.getArticleDetail = function(vue, articleId, success) {
  var request =
    window.webapp.server.address + "/article/" + articleId + "/detail";

  vue.$http.get(request).then(function(res) {
    if (res.body.status == "ok") {
      if (res.body.status == "ok") {
        success(res);
      } else {
        window.webapp.f.ErrorHandle(res.body);
      }
    }
  });
};

//获取封面预览文章
window.webapp.f.getSomeArticle = function(vue, limit, offset, success) {
  var request =
    window.webapp.server.address +
    "/article?offset=" +
    offset +
    "&" +
    "limit=" +
    limit;
  request += "&filed=view_img";
  vue.$http.get(request).then(res => {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

//获取文章数量
window.webapp.f.getArticleLen = function(vue, getlen, limit, offset, success) {
  var request = window.webapp.server.address + "/article?";
  if (getlen == true) {
    request += "flag=len&";
  }
  request += "limit=" + limit + "&offset=" + offset;
  vue.$http.get(request).then(res => {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

//发送文章
window.webapp.f.sendArticle = function(vue, form, success) {
  var url = window.webapp.server.address + "/article";
  vue.$http.post(url, form).then(res => {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

//获取文章回复数目
window.webapp.f.getArticleDetailCommentList = function(
  vue,
  articleId,
  success
) {
  var request =
    window.webapp.server.address + "/comment/" + articleId + "?flag=len";
  vue.$http.get(request).then(function(res) {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

//获取回复
window.webapp.f.getComment = function(vue, articleId, success) {
  var request = window.webapp.server.address;
  request += "/comment/" + articleId;

  vue.$http.get(request).then(function(res) {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

//发送回复
window.webapp.f.sendComment = function(vue, form, success) {
  var request = window.webapp.server.address;
  request += "/comment";
  vue.$http.post(request, form).then(function(res) {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

//删除文章
window.webapp.f.deleteArticle = function(vue, articleId, success) {
  var request = window.webapp.server.address + "/article/" + articleId;
  console.log(vue.$http);
  vue.$http.delete(request).then(function(res) {
    if (res.body.status == "ok") {
      success(res);
    } else {
      window.webapp.f.ErrorHandle(res.body);
    }
  });
};

/* 一些接口功能 */

export default {
  name: "App",
  components: {
    Index
  }
};
</script>

<style scoped>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  width: 100%;
  box-sizing: border-box;
}
</style>

