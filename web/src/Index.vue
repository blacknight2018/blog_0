<template>
  <div id="index">
    <!-- PC端 -->
    <div id="pc">
      <div id="nav">
        <div id="nav-content">
          <div id="nav-content-left">
            <!-- <el-avatar :size="50" :src="me.avatar"></el-avatar> -->
            <img :src="me.avatar" alt />
          </div>
          <div id="nav-content-middle">
            <div>
              <el-menu :default-active="this.nav.cur" class="el-menu-demo" mode="horizontal">
                <el-menu-item
                  v-for="(item,idx) in nav.items"
                  :key="item.name"
                  :index="String(idx+1)"
                >
                  <router-link :to="item.path">{{item.name}}</router-link>
                </el-menu-item>
              </el-menu>
            </div>
          </div>
          <div id="nav-content-right">
            <div v-if="getLoginStatus()==false">
              <el-button type="primary" @click="openLogin=true">登录</el-button>
              <el-button type="danger">注册</el-button>
            </div>
            <div v-if="getLoginStatus()">
              <el-button type="warning" @click="loginOut()">退出登录</el-button>
            </div>
          </div>
        </div>
      </div>
      <div id="layout">
        <div id="layout-content">
          <router-view></router-view>
        </div>
      </div>
    </div>

    <!-- 非PC端 -->
    <div id="npc"></div>

    <!-- 抽屉或者弹出窗口  -->
    <!-- 登录框  -->
    <el-dialog title="登录" width="30%" :visible.sync="openLogin">
      <el-form ref="login.form" :model="login.form" label-width="80px">
        <el-input v-model="login.form.user" placeholder="账号"></el-input>
        <el-input v-model="login.form.password" placeholder="密码"></el-input>
      </el-form>
      <el-button type="primary" round @click="login.methods.login()">登录</el-button>
    </el-dialog>
  </div>
</template>

<script>
import router from "./router/index";
import "element-ui/lib/theme-chalk/index.css";
import ElementUI from "element-ui";

export default {
  mounted() {
    this.getUserInformation();
  },
  methods: {
    getUserInformation() {
      window.webapp.f.getUserInformation(this, undefined, res => {
        window.webapp.me = res.body.data;
        this.me = res.body.data;
      });
    },
    getLoginStatus() {
      if (window.webapp.me == undefined) return false;
      return true;
    },
    loginOut() {
      window.webapp.f.logout(this, res => {
        if (res.body.status == "ok") {
          this.$message("退出成功");
          setTimeout(() => {
            window.location.reload();
          }, 1000);
        }
      });
    }
  },
  router,
  data() {
    return {
      me: {
        avatar: "",
        uid: "",
        user: "",
        type: 0
      },
      nav: {
        cur: "1",
        items: [
          {
            name: "首页",
            path: "/Home"
          },
          {
            name: "文章",
            path: "/article"
          },
          {
            name: "归档",
            path: "/archive/"
          },
          {
            name: "后台",
            path: "/admin"
          }
        ],
        drawer: false
      },
      openLogin: false,
      login: {
        form: {
          user: "",
          password: ""
        },
        methods: {
          login: () => {
            window.webapp.f.login(this, this.login.form, res => {
              window.webapp.me = JSON.parse(res.bodyText);
              this.openLogin = false;
              this.$notify({
                title: "登录",
                message: "登录成功"
              });
              this.getUserInformation();
            });
          }
        }
      }
    };
  }
};
</script>

<style lang="less" scoped>
/* 移动端 非移动端 媒体选择 width:1200px统一移动端 */
@media screen and (max-width: 1200px) {
  #pc {
    display: none;
  }
}
@media screen and (min-width: 1200px) {
  #npc {
    display: none;
  }
}
#index {
  box-sizing: border-box;
  width: 100%;
  .el-input {
    margin-bottom: 20px;
  }
  /* 非移动端 */
  #pc {
    width: 100%;
    #nav {
      position: fixed;
      top: 0;
      left: 0;
      z-index: 1600;
      height: 60px;
      width: 100%;
      border-bottom: solid 1px #eee;
      &-content {
        width: 1200px;
        height: 100%;
        margin: 0 auto;
        display: flex;
        &-left {
          height: 100%;
          width: 152px;
          display: flex;
          align-items: center;
          background-color: white;
          img {
            width: 50px;
            height: 50px;
            border-radius: 50%;
          }
        }
        &-middle {
          flex: 1;
          li {
            padding: 0px;
          }
          a {
            text-decoration: none;
            height: 100%;
            width: 100%;
            display: inline-block;
            width: 68px;
            text-align: center;
          }
        }
        &-right {
          width: 183px;
          display: inline-block;
          display: flex;
          justify-content: center;
          align-items: center;
          background-color: white;
        }
      }
    }
    #layout {
      margin: 0px;
      padding: 0px;
      width: 100%;
      margin-top: 80px;
      &-content {
        width: 1200px;
        margin: 0 auto;
        box-sizing: border-box;
      }
    }
  }
}
</style>