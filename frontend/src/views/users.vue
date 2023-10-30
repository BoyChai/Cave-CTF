<template>
  <div style="margin-top:20px;margin-left: 100px;margin-right: 100px;">
   <el-table :data="data" style="width: 100%">
     <el-table-column prop="ID" label="ID" width="180" />
     <el-table-column prop="Name" label="Name" width="180" />
     <el-table-column prop="Alias" label="Alias" />
    </el-table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "users",
  data() {
    return {
      data:[],
      interval: "",
      // data:[{"ID":1,"Name":"test1","Score":"0"},{"ID":2,"Name":"tset2","Score":"1"},{"ID":3,"Name":"test3","Score":"1"}],
    }
  },
  methods: {
    getData(){
      axios.get("/get/users").then(res => {
        this.data = res.data.data;
      }).catch(e => {
        this.$message({
          message:"获取数据错误",
          type:"error"
        })
      })
    }
  },
  created() {
    this.getData()
    this.interval = setInterval(this.getData, 2000);

  },beforeUnmount() {
    clearInterval(this.interval);

  }
}
</script>

<style scoped>

</style>