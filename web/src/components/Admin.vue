<template>
  <div id="app_d">
    <el-tabs>
      <el-tab-pane label="文章管理">
        <el-tabs tab-position="left" type="border-card">
          <el-tab-pane label="发布文章">
            <ArticleWrite></ArticleWrite>
          </el-tab-pane>
          <el-tab-pane label="文章列表">
            <ArticleLists></ArticleLists>
          </el-tab-pane>
        </el-tabs>
      </el-tab-pane>
      <el-tab-pane label="用户管理"></el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import ElementUI from "element-ui";
import MavonEditor from "mavon-editor";
import Vue from "vue";
import "element-ui/lib/theme-chalk/index.css";
import "mavon-editor/dist/css/index.css";
import ArticleWrite from "./Admin/ArticleWrite";
import ArticleLists from "./Admin/ArticleLists";
Vue.use(ElementUI);
Vue.use(MavonEditor);

export default {
  components: {
    ArticleWrite,
    ArticleLists
  }
};
</script>

<style scope="scoped">
#app_d {
  background-color: transparent;
  width: 100%;
  height: 500px;
}
</style>>
