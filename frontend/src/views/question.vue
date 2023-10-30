<template>
<div style="float:left;" v-for="item in questionList">
  <el-card shadow="hover" style="margin-left: 10px;margin-top: 10px" class="box-card">
    <template #header>
      <div class="card-header">
        <span>{{item.Title}}</span>
        <el-button class="button" @click="openQuestion(item)" text>open</el-button>
      </div>
    </template>
    <div style="height: 120px">
      <br>
      {{item.Describe}}
      <br>
      <br>
      <br>
      <div v-if="isCompleted(item)" style="color: green">
        已完成
        <el-icon><Select /></el-icon>
      </div>
    </div>
  </el-card>
</div>
  <el-dialog
      v-model="dialogVisible"
      :title="question.Title"
      width="30%"
  >
    <span>{{question.Describe}}</span>
    <br>
    <br>
    <a target="_blank" :href="question.Annex">下载附件</a>
    <br>
    <br>

    <template #footer>
      <span class="dialog-footer">
<!--        <el-button @click="dialogVisible = false">Cancel</el-button>-->
<!--        <el-button type="primary" @click="dialogVisible = false">-->
<!--          Confirm-->
<!--        </el-button>-->
        <el-input v-if="!isCompleted(question)" style="width: 82%" v-model="flag" placeholder="flag格式为flag{}" />
        <el-button v-if="!isCompleted(question)" style="margin-left: 5px" @click="putFlag" :disabled="buttonDisabled" type="primary">提交</el-button>
        <div v-if="isCompleted(question)" style="color: green">
         已完成
         <el-icon><Select /></el-icon>
        </div>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import axios from "axios";
export default {
  name: "question",
  data() {
    return {
      question: "",
      questionList:[],
      dialogVisible: false,
      flag:"",
      personalRanking:[],
      userRecords:[],
      buttonDisabled:false,
    }
  },
  computed: {
    // 计算属性，用于判断题目是否已完成
    isCompleted: function () {
      return (question) => {
        return this.userRecords.some((record) => record.QuestionTitle === question.Title);
      };
    },
  },
  methods:{
    openQuestion(question) {
      this.flag = ""
      this.question=question
      this.dialogVisible=true
    },
    putFlag() {
      if (!this.buttonDisabled) {
        // 禁用按钮
        this.buttonDisabled = true;

        // 1秒后启用按钮
        setTimeout(() => {
          this.buttonDisabled = false;
        }, 1000);
        // 这里放置你的按钮点击后的逻辑
        // 例如：执行一个函数或者触发一个事件
        // 在这里添加你需要的逻辑
        let alias;
        alias=this.getCookieValue("alias")
        axios.put("/put/flag",{
          question_id: this.question.ID,
          user: alias,
          score: this.question.Score,
          flag: this.flag,
        }).then(res => {
          this.$message({
            message:"提交成功",
            type:"success"
          })
          this.getQuestion()
          this.getPersonalRanking()
          this.dialogVisible=false
        }).catch(e => {
          this.$message({
            message:e.response.data.msg,
            type:"error"
          })
        })
      }
    },
    getQuestion() {
      axios.get("/get/questionlist").then(res=>{
        this.questionList=res.data.data;
        console.log(this.questionList)
      })
    },
    getPersonalRanking() {
      axios.get("/get/personal/ranking?alias="+this.getCookieValue("alias")).then(res=>{
        this.personalRanking=res.data.data;
        this.userRecords=res.data.data;
        console.log(this.personalRanking)
      })
    },
    getCookieValue(cookieName) {
      const name = cookieName + "=";
      const decodedCookie = decodeURIComponent(document.cookie);
      const cookieArray = decodedCookie.split(';');

      for (let i = 0; i < cookieArray.length; i++) {
        let cookie = cookieArray[i].trim();
        if (cookie.indexOf(name) === 0) {
          return cookie.substring(name.length, cookie.length);
        }
      }
      return "";
    }

  },
  mounted() {
    console.log(this.questionList)
  },
  created() {
    this.getQuestion()
    this.getPersonalRanking()
  }
}
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  width: 375px;
}

</style>