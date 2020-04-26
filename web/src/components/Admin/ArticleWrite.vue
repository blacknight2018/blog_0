<template>
  <div>
    <el-input class="input" v-model="articleNew.title" placeholder="标题"></el-input>
    <!-- author由登录用户取代  -->
    <el-input v-if="false" class="input" v-model="articleNew.author" placeholder="作者"></el-input>
    <el-input class="input" v-model="articleNew.description" placeholder="简介描述"></el-input>
    <mavon-editor @change="mavonChange" />
    <div
      style="margin-top:20px;border-top:1px solid #eee;border-bottom:1px solid #eee;padding-top:20px;"
    >
      <el-upload
        class="upload-demo"
        ref="upload"
        :action="getActionUrl()"
        :on-preview="handlePreview"
        :on-remove="handleRemove"
        :file-list="fileList"
        :auto-upload="false"
        :on-success="upLoadSuccess"
        with-credentials
      >
        <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
        <el-button
          style="margin-left: 10px;"
          size="small"
          type="success"
          @click="submitUpload"
        >上传到服务器</el-button>
        <div slot="tip" class="el-upload__tip">不要传得太大了</div>
      </el-upload>
    </div>
    <div id="preview">
      <el-button icon="el-icon-view" type="primary" size="small" @click="checkFile">选择封面</el-button>
      <input type="file" id="fileinput" style="display: none;" @change="checkFileSure" />
      <img src alt id="viewimg" />
    </div>
    <div id="action">
      <el-button type="primary" @click="formSubmit">发布</el-button>
    </div>
  </div>
</template>
<script>
import ElementUI from "element-ui";
import MavonEditor from "mavon-editor";
import Vue from "vue";
import "element-ui/lib/theme-chalk/index.css";
import "mavon-editor/dist/css/index.css";

Vue.use(ElementUI);
Vue.use(MavonEditor);

export default {
  data() {
    return {
      articleNew: {
        content: undefined,
        title: undefined,
        description: undefined,
        view_img: undefined,
        file: []
      },
      imageUrl: "",
      fileList: [
        // {
        //   name: "food.jpeg",
        //   url:
        //     "https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100"
        // },
        // {
        //   name: "food2.jpeg",
        //   url:
        //     "https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100"
        // }
      ]
    };
  },
  methods: {
    mavonChange(value, render) {
      this.articleNew.content = render;
    },
    getActionUrl() {
      return window.webapp.uploadurl;
    },
    submitUpload() {
      this.$refs.upload.submit();
    },
    handleRemove(file, fileList) {
      // console.log(file, fileList);
    },
    handlePreview(file) {
      // console.log(file);
    },
    upLoadSuccess(response, file, fileList) {
      console.log(response, file, fileList);
      if (!(response.data.Fid > 0 && response.status == "ok")) {
        fileList.splice(0, 1);
        return;
      }
      this.articleNew.file.push(response.data.Fid);
      console.log(this.articleNew);
    },
    formSubmit() {
      var param =
        this.articleNew.content &&
        this.articleNew.title &&
        this.articleNew.description &&
        this.articleNew.view_img;
      console.log(this.articleNew);
      if (param == undefined) {
        this.$message("参数不完整");
        return;
      }
      window.webapp.f.sendArticle(
        this,
        JSON.stringify(this.articleNew),
        res => {
          this.$message("发布成功");
          setTimeout(() => {
            location.reload();
          }, 700);
        }
      );
    },
    checkFile() {
      document.querySelector("#fileinput").click();
    },
    checkFileSure(val) {
      var files = document.getElementById("fileinput").files;
      var reader = new FileReader();
      var that = this;
      reader.readAsDataURL(files[0]);
      reader.onload = function(e) {
        document.getElementById("viewimg").src = this.result;
        that.articleNew.view_img = this.result;
      };
    }
  }
};
</script>
</script>
<style scoped>
div#action {
  text-align: left;
  padding-top: 20px;
}
.input {
  margin-bottom: 20px;
}
#preview {
  border-bottom: 1px solid #eee;
  text-align: left;
  padding-top: 20px;
  padding-bottom: 10px;
}
img#viewimg {
  max-width: 250px;
  max-height: 250px;
  width: 100%;
  height: 100%;
  vertical-align: bottom;
}
</style>