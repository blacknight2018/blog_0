<template>
  <div id="articles">
    <div id="articles-left">
      <div
        class="articles-item"
        v-for="(item) in this.lists"
        :key="item.id"
        @click="openItem(item.id)"
      >
        <div id="article-left">
          <h4 id="article-left-title">{{item.title}}</h4>
          <p id="article-left-abstract">{{item.description}}</p>
          <div id="article-left-meta">
            <span id="view">查看 1395</span>
          </div>
        </div>
        <div id="article-right">
          <img :src="item.view_img" alt />
        </div>
      </div>
    </div>
    <div id="articles-right"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      lists: [
        // {
        //   id: 0,
        //   title: "2012",
        //   author: "chen",
        //   description: "hello",
        //   like: 0,
        //   view_img: ""
        // }
      ],
      offset: 0,
      limit: 5,
      alreadyload: {}
    };
  },
  methods: {
    openItem(id) {
      var url = window.location.host + "/articleDetail";
      url += "?id=";
      url += id;
      url = window.location.protocol + "//" + url;
      console.log(url);
      window.open(url, "_blank");
    }
  },
  destroyed() {
    window.onmousewheel = undefined;
  },
  mounted() {
    let getLimitArticle = () => {
      function ff(vue, i) {

        var key = "a" + vue.offset;
        if (vue.alreadyload[key] == true) {
          return;
        }
        window.webapp.f.getSomeArticle(vue, 1, vue.offset, res => {
          vue.offset++;
          //console.log(res.body);
          if (res.body.data.length == 0) {
            return;
          }
          res.body.data[0].title = Base64.decode(res.body.data[0].title)
          res.body.data[0].description = Base64.decode(res.body.data[0].description)
          vue.lists.push(res.body.data[0]);
          if (vue.limit == i) {
            return;
          }
          ff(vue, i + 1);
        });
        vue.alreadyload[key] = true;
      }
      ff(this, 1);

    };

    getLimitArticle();
    window.onmousewheel = function() {
      let scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop;
      let clientHeight =
        document.documentElement.clientHeight || document.body.clientHeight;
      let scrollHeight =
        document.documentElement.scrollHeight || document.body.scrollHeight;
      if (scrollTop + clientHeight + 10 >= scrollHeight) {
        getLimitArticle();
      }
    };
  }
};
</script>

<style lang="less" scoped>
#articles {
  min-height: 126px;
  box-sizing: border-box;
  &-left {
    width: 70%;

    height: 100%;
    box-sizing: border-box;
    .articles-item {
      box-sizing: content-box;
      display: flex;
      border-bottom: solid 1px #eee;
      cursor: pointer;
      #article-left {
        width: 680px;
        padding-top: 15px;
        padding-bottom: 15px;
        height: inherit;

        &-title {
          color: #333;
          margin: 7px 0 4px;
          display: inherit;
          font-size: 18px;
          font-weight: 700;
          line-height: 1.5;
        }
        &-abstract {
          min-height: 30px;
          margin: 0 0 8px;
          font-size: 13px;
          line-height: 24px;
          color: #555;
        }
        &-meta {
          padding-right: 0 !important;
          font-size: 12px;
          font-weight: 400;
          line-height: 20px;
          margin: 0px;
          span {
            margin-right: 10px;
            color: #666;
          }
        }
      }
      #article-right {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: center;
        img {
          width: 125px;
          height: 100px;
          max-width: 100%;
          max-height: 100%;
        }
      }
    }
  }
  &-right {
    width: 30%;
  }
}
</style>