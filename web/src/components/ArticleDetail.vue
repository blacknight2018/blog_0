<template>
  <div id="root">
    <h1 id="title">{{response.title}}</h1>
    <div id="author">
      <div id="info">
        <div id="info-left">
          <img id="avatar" :src="user.avatar" alt />
        </div>
        <div id="info-right">
          <span>
            <span id="name">{{user.name}}</span>
            <div id="meta">
              <span>{{response.create_time}}</span>
            </div>
          </span>
        </div>
      </div>
    </div>
    <link
      href="https://cdn.bootcss.com/github-markdown-css/2.10.0/github-markdown.min.css"
      rel="stylesheet"
    />
    <div id="content" class="markdown-body" v-html="response.content"></div>

    <div style="margin-bottom:10px">
      <el-card class="box-card" v-if="this.response.file!=undefined && this.response.file.length>0">
        <div slot="header" class="clearfix">
          <span>附件</span>
        </div>
        <div v-for="val in this.response.file" :key="val.fid" class="text item">
          <el-link type="primary" :href="getUrlByFid(val.fid)">{{val.name}}</el-link>
        </div>
      </el-card>
    </div>

    <div id="comment">
      <el-form>
        <el-form-item>
          <el-input type="textarea" rows="5" placeholder="请输入内容" v-model="comment.content"></el-input>
        </el-form-item>
      </el-form>
      <el-button type="primary" id="reply_button" @click="sendComment">回复</el-button>
    </div>

    <div id="comment-list">
      <div id="comment-list-title">
        <span>{{commentlist.num}}条评论</span>
      </div>
    </div>
    <div id="comment-content">
      <div id="comment-content-item" v-for="(val,idx) in commentlist.content" :key="val.cid">
        <div :id="val.cid">
          <h4 style="text-align:right">{{idx+1}}楼</h4>
          <div id="comment-content-item-header">
            <div id="comment-content-item-header-left">
              <img :src="val.avatar" alt />
            </div>
            <div id="comment-content-item-header-right">
              <div id="comment-content-item-header-right-name" class="comment-name">{{val.name}}</div>
              <div
                id="comment-content-item-header-right-time"
                class="comment-time"
              >{{val.last_time}}</div>
            </div>
          </div>
          <div id="comment-content-item-content">
            <span
              style="color:blue;cursor: pointer;"
              v-if="commentlist.getOrderByCid(val.replyto_cid)>0"
            >
              <a :href="'#'+val.replyto_cid">@回复{{commentlist.getOrderByCid(val.replyto_cid)}}楼</a>
            </span>
            {{val.content}}
          </div>
          <div id="comment-content-item-header-message">
            <el-button type="default" size="small" @click="openCommentDialog(val.cid)">回复</el-button>
          </div>
        </div>
      </div>
    </div>

    <!-- 弹出的回复框 -->
    <el-dialog title="回复" :visible.sync="commentDialogVisible" width="50%">
      <el-form>
        <el-form-item>
          <el-input type="textarea" rows="6" placeholder="请输入回复内容" v-model="commentContent"></el-input>
        </el-form-item>
      </el-form>
      <div>
        <p></p>
        <el-button type="primary" @click="sendComment2()">回复</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import "mavon-editor/dist/css/index.css";
let Base64 = require("js-base64").Base64;
export default {
  mounted() {
    window.webapp.f.getArticleDetail(this, this.getCurHrefArticleId(), res => {
      var obj = res.body.data;
      this.response = obj;
      this.response.content = Base64.decode(this.response.content);
      this.response.title = Base64.decode(this.response.title);

      if (this.response.file != undefined)
        this.response.file = this.response.file;
      //加载作者头像信息
      window.webapp.f.getUserInformation(this, this.response.author, res => {
        this.user.avatar = res.body.data.avatar;
        this.user.name = res.body.data.user;
      });
    });

    //获取回复数目
    window.webapp.f.getArticleDetailCommentList(
      this,
      this.getCurHrefArticleId(),
      res => {
        this.commentlist.num = res.body.data.length;
      }
    );
    //加载回复
    window.webapp.f.getComment(this, this.getCurHrefArticleId(), res => {
      // this.commentlist.content.content = Base64.decode(
      //   this.commentlist.content
      // );
      this.commentlist.content = res.body.data;
      for (let k = 0; k < this.commentlist.content.length; k++) {
        this.commentlist.content[k].content = Base64.decode(
          this.commentlist.content[k].content
        );
      }
    });
  },
  methods: {
    sendComment() {
      this.comment.article_id = this.getCurHrefArticleId();
      window.webapp.f.sendComment(this, this.comment, res => {
        // if (res.body.status == "ok") {
        this.$message("回复成功");
        setTimeout(() => {
          location.reload();
        }, 700);
        // }
      });
    },
    sendComment2() {
      var curId = this.getCurHrefArticleId();
      var form = {
        article_id: curId,
        content: this.commentContent,
        replyto_cid: this.cur_cid
      };
      window.webapp.f.sendComment(this, form, res => {
        // if (res.body.status == "ok") {
        this.$message("回复成功");
        setTimeout(() => {
          location.reload();
        }, 700);
        // }
      });
    },
    getCurHrefArticleId() {
      var href = window.location.href;
      var params = window.location.search;
      var pos = params.indexOf("=");
      var id = params.substring(pos + 1, params.length);
      return id;
    },
    openCommentDialog(cid) {
      this.commentDialogVisible = true;
      this.cur_cid = cid;
    },
    getUrlByFid(fid) {
      return window.webapp.fileurl + fid;
    }
  },
  data() {
    return {
      response: {
        title: "加载失败!!",
        author: "加载失败!!", //这里指的是作者的UID
        content: "加载失败!!",
        create_time: "1970-01-10 12:00:00",
        file: [
          {
            fid: "11",
            name: "aa"
          }
        ]
      },
      user: {
        avatar: "",
        name: ""
      },
      comment: {
        content: "",
        replyto_cid: 0,
        article_id: 0
      },
      cur_cid: 0,
      commentDialogVisible: false,
      commentContent: "",
      commentlist: {
        num: 0,
        getOrderByCid: cid => {
          for (let i = 0; i < this.commentlist.content.length; i++) {
            if (cid == this.commentlist.content[i].cid) {
              return i + 1;
            }
          }
          return 0;
        },
        content: [
          {
            cid: 11,
            replyto_cid: 0,
            avatar: "",
            name: "",
            last_time: "",
            content: ""
          }
        ]
      }
    };
  }
};
</script>

<style lang="less" scoped >
@media screen and (max-width: 1200px) {
  div#root {
    width: auto;
  }
}
@media screen and (min-width: 1200px) {
  div#root {
    width: 832px;
  }
}
#root {
  text-align: left;

  #title {
    text-align: center;
    font-size: 34px;
    font-weight: 700;
    margin: 20px 0 0;
  }
  #author {
    margin: 30px 0 40px;
  }
  #name {
    margin-right: 3px;
    font-size: 16px;
    vertical-align: middle;
    text-align: left;
  }
  #info {
    display: flex;
    &-left {
      #avatar {
        width: 48px;
        height: 48px;
        border-radius: 50%;
      }
    }
    &-right {
      flex: 1;
      margin-left: 8px;
      display: flex;
      align-items: center;
    }
  }
  #meta {
    margin-top: 5px;
    font-size: 12px;
    color: #969696;
    text-align: left;
  }
  #content {
    min-height: 300px;
    margin-bottom: 30px;
  }
  #reply_button {
    margin-top: 1.25rem;
    margin-bottom: 1.25rem;
  }
  #comment {
    border-bottom: 1px solid #eee;
  }
  #comment-list {
    border-bottom: 1px solid #eee;
    margin-top: 20px;
    padding-bottom: 20px;
    font-size: 17px;
    font-weight: 700;
  }
  #comment-content {
    &-item {
      margin-top: 50px;
      border-bottom: 1px solid #f0f0f0;
      .comment-name {
        font-size: 15px;
        color: #333;
      }
      .comment-time {
        font-size: 12px;
        color: #969696;
      }
      &-content {
        min-height: 40px;
        margin-top: 20px;
      }
      &-header {
        height: 42px;
        display: flex;
        align-items: center;
        &-left {
          img {
            width: 40px;
            height: 40px;
            border-radius: 50%;
            vertical-align: bottom;
          }
          margin-right: 10px;
        }
        &-message {
          margin-bottom: 20px;
        }
      }
      &-other {
        margin: 20px;
        padding: 10px;
        border-left: 2px solid #f0f0f0;
      }
    }
  }
}
</style>