<template>
<div style="float:left;" v-for="item in questionList">
  <el-card style="margin-left: 10px;margin-top: 10px" class="box-card">
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
    <a target="_blank" href="{{question.Annex}}">下载附件</a>
    <br>
    <br>

    <template #footer>
      <span class="dialog-footer">
<!--        <el-button @click="dialogVisible = false">Cancel</el-button>-->
<!--        <el-button type="primary" @click="dialogVisible = false">-->
<!--          Confirm-->
<!--        </el-button>-->
        <el-input style="width: 82%" v-model="flag" placeholder="flag格式为flag{}" />
        <el-button style="margin-left: 5px" @click="putFlag" type="primary">提交</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
export default {
  name: "question",
  data() {
    return {
      question: "",
      questionList:[{"ID":1,"Title":"test1","Describe":"我要flag我要flag我要flag我要flag我要flag我要flag我要","Flag":"flag{123}","Type":"1","Annex":"http://test.etst/11.txt","Score":"10"},{"ID":2,"Title":"test2","Describe":"快给我flag","Flag":"flag{1234}","Type":"1","Annex":"http://test.etst/11.txt","Score":"10"},{"ID":1,"Title":"test3","Describe":"我要flag","Flag":"flag{123}","Type":"1","Annex":"http://test.etst/11.txt","Score":"10"},{"ID":1,"Title":"test1","Describe":"我要flag","Flag":"flag{123}","Type":"1","Annex":"http://test.etst/11.txt","Score":"10"},{"ID":1,"Title":"test1","Describe":"我要flag","Flag":"flag{123}","Type":"1","Annex":"http://test.etst/11.txt","Score":"10"},],
      dialogVisible: false,
      flag:"",
    }
  },
  methods:{
    openQuestion(question) {
      this.question=question
      this.dialogVisible= true

    },
    putFlag() {

    }

  },
  mounted() {
    console.log(this.questionList)
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