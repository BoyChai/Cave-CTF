<template>
  <div>
    <h3>分值统计</h3>
    <el-row>
      <el-col v-for="item in ranking" :key="item.user" :span="6">
        <el-statistic :title="item.user" :value="item.score" />
      </el-col>
    </el-row>
    <h3>分值历史</h3>
    <div style="margin-top: 25px">
      <el-table sortable="custom"  :data="rankingList" style="width: 100%;">
        <el-table-column prop="ID" label="ID" width="180" />
        <el-table-column prop="User" label="User" width="180" />
        <el-table-column prop="Score" label="Score" width="180" />
        <el-table-column prop="QuestionTitle" label="QuestionTitle" />
      </el-table>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "ranking",
  data() {
    return {
      rankingList: [],
      ranking: [],
      intervalId: "",

    }
  },
  methods: {
    getRankingList() {
      axios.get("/get/ranking").then(r =>{
        this.rankingList = []
        this.rankingList = r.data.data
        this.summary()
      }).catch(e=>{
        this.$message({
          message:"获取分数列表出现错误",
          type:"error"
        })
      })
    },
    summary() {
      this.ranking = []
      for (let i = 0; i < this.rankingList.length; i++) {
        const user = this.rankingList[i].User;
        const score = parseInt(this.rankingList[i].Score, 10); // 将字符串转换为整数

        if (!isNaN(score)) { // 检查是否成功转换为数字
          let userExists = false;
          for (let j = 0; j < this.ranking.length; j++) {
            if (this.ranking[j].user === user) {
              this.ranking[j].score += score;
              userExists = true;
              break;
            }
          }
          if (!userExists) {
            this.ranking.push({ user, score });
          }
        }
      }
      this.ranking.sort((a, b) => b.score - a.score);

    },
  },
  created() {
    this.getRankingList()
    this.rankingList.reverse()
    this.intervalId =setInterval(this.getRankingList, 2000);
  },
  beforeUnmount() {
    clearInterval(this.intervalId);
  }
}
</script>

<style scoped>
.el-col {
  text-align: center;
  margin-top: 15px;
}

</style>