<template>
  <div>
    <div id="condition">
      <el-dropdown split-button type="primary" @command="commandOrder">
        {{this.condition.order}}
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item command="asc">正序</el-dropdown-item>
          <el-dropdown-item command="desc">倒序</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
      <el-button icon="el-icon-search" v-on:click="this.getResultLen">筛选</el-button>
    </div>
    <el-table :data="article" stripe>
      <el-table-column label="ID" prop="id"></el-table-column>
      <el-table-column label="标题" prop="title"></el-table-column>
      <el-table-column label="作者" prop="author"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button @click="deleteItem(scope.row)" type="danger" icon="el-icon-delete"></el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      @current-change="handleCurrentChange"
      :total="this.page.total"
      :current-page="this.page.cur_page"
      :page-size="this.page.page_limit"
      layout="prev,pager,next"
    ></el-pagination>
  </div>
</template>
<script>
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import Vue from "vue";
Vue.use(ElementUI);

export default {
  data() {
    return {
      condition: {
        order: "desc"
      },
      page: {
        cur_page: 1,
        page_limit: 10,
        total: 0
      },
      article: [
        // {
        //   id: 0,
        //   title: "精品电子书！",
        //   author: "Chen"
        // }
      ]
    };
  },
  methods: {
    commandOrder(c) {
      this.condition.order = c;
    },
    getResult() {
      window.webapp.f.getArticleLen(
        this,
        false,
        this.page.page_limit,
        (this.page.cur_page - 1) * this.page.page_limit,
        res => {
          this.article = res.body.data;
          for (let i = 0; i < this.article.length; i++) {
            this.article[i].title = Base64.decode(this.article[i].title);
          }
        }
      );
    },
    getResultLen() {
      window.webapp.f.getArticleLen(
        this,
        true,
        this.page.page_limit,
        this.page.cur_page * this.page.page_limit,
        res => {
          this.page.total = res.body.data.length;
          this.handleCurrentChange(1);
        }
      );
    },
    handleCurrentChange(cur) {
      this.page.cur_page = cur;
      this.getResult();
    },
    deleteItem(row) {
      let idx = this.article.indexOf(row);
      if (-1 != idx) {
        window.webapp.f.deleteArticle(this, this.article[idx].id, res => {
          this.article.splice(idx, 1);
          this.$message("删除成功");
        });
      }
    }
  }
};
</script>

<style scope="scoped">
#condition {
  text-align: left;
}
.el-pagination {
  text-align: center;
}
</style>